use std::rc::Rc;
use std::cell::RefCell;
// use std::sync::{Arc, Mutex};
use wasm_bindgen::prelude::*;
use web_sys::WebGl2RenderingContext;
// use crate::magic_square::vertices::Vertices;
use crate::magic_square::shader_compiler::ShaderCompiler;
use crate::magic_square::program_linker::ProgramLinker;
use crate::magic_square::transformations::{Rotation, RotationSequence, Translation};
use crate::magic_square::geometry::{Geometry, Shape};
use crate::magic_square::geometry::cache::{Cache as GeometryCache, CACHE_CAPACITY};
use crate::magic_square::ui_buffer::UiBuffer;

use super::settings::{MouseTracking, Settings};
// use crate::magic_square::traits::VertexStore;
// use super::geometry::icosohedron::Icosohedron;
// use crate::magic_square::worker::Worker;

#[wasm_bindgen]
extern "C" {
    #[wasm_bindgen(js_namespace = console)]
    pub fn log(s: &str);

    #[wasm_bindgen(js_name = "performance")]
    pub static PERFORMANCE: web_sys::Performance;
}

#[derive(Clone, Copy)]
pub enum Axis {
    X,
    Y,
    Z
}

// => keep buffer in a Refcell in an RC
// => one event listener has mutable reference to write
// => another closure has the animation loop with an immutable reference

#[wasm_bindgen]
pub struct MagicSquare;

#[wasm_bindgen]
impl MagicSquare {
    // Entry point into Rust WASM from JS
    // https://rustwasm.github.io/wasm-bindgen/examples/webgl.html
    pub fn run() -> Result<i32, JsValue> {
        // testing multithreading
        //
        // let (to_worker, from_main) = std::sync::mpsc::channel();
        // let (to_main, from_worker) = std::sync::mpsc::channel();
        // Worker::spawn(move || { to_main.send(from_main.recv().unwrap() + 1.0); });
        // to_worker.send(1.0);
        // assert_eq!(from_worker.recv().unwrap(), 2.0);
        

        let magic_square = MagicSquare::magic_square(); // awesome naming, great job!
        let magic_square = Rc::new(magic_square);
        let canvas = MagicSquare::canvas().dyn_into::<web_sys::HtmlCanvasElement>()?;
        let canvas = Rc::new(canvas);

        // TODO: handle resize
        // add height and width fields to MagicSquare
        // run returns MagicSquare instance to js
        // js uses svelte watch resize to update height and width
        // pass immutable reference of h&w to closure
        let height:i32 = canvas.client_height();
        let width:i32 = canvas.client_width();
        // incriment idx_delay each render
        // when idx_delay reaches a desired delay value
        // incriment idx_offset
        let mut color_idx_offset_delay: [usize; 2] = [0, 0];
        //Arc::new(Mutex::new(
        let geometry_cache = GeometryCache::new(
                26, 
                [[0.0; 1_200]; CACHE_CAPACITY], 
                [[0.0, 0.0, 0.0, 0.0]; CACHE_CAPACITY],
                [Shape::None; CACHE_CAPACITY]
            );

        let mut geometry_cache = Rc::new(RefCell::new(geometry_cache));
        // )); 
        
        let form = MagicSquare::form();
        let form = Rc::new(form);
        let ui_buffer = UiBuffer::new();
        let ui_buffer = Rc::new(RefCell::new(ui_buffer));
        let mouse_pos_buffer: [f32; 2] = [0.0, 0.0];
        let mouse_pos_buffer: Rc<RefCell<[f32; 2]>> = Rc::new(RefCell::new(mouse_pos_buffer));

        let context: web_sys::WebGl2RenderingContext = MagicSquare::context(&canvas).unwrap();
        context.clear_color(1.0, 1.0, 0.0, 0.0);
        context.clear(WebGl2RenderingContext::COLOR_BUFFER_BIT);

        {
            // init UI control settings listener
            let magic_square = magic_square.clone();
            let form = form.clone();
            let ui_buffer = ui_buffer.clone();

            let closure_handle_input =
                Closure::<dyn FnMut(_)>::new(move |event: web_sys::Event| {
                    let input = event
                        .target()
                        .unwrap()
                        .dyn_into::<web_sys::HtmlInputElement>()
                        .unwrap();
                    let id = input.id();
                    let val = input.value();
                    // log(&id);
                    // log(&val);
                    ui_buffer.clone().borrow_mut().update(id, val);
                    //trigger re-render
                    magic_square.dispatch_event(&web_sys::Event::new("render").unwrap()).unwrap();
                });

            form.add_event_listener_with_callback(
                "input",
                closure_handle_input.as_ref().unchecked_ref(),
            )
            .unwrap();
            closure_handle_input.forget(); // allow JS to garbage collect the listener
        }

        {
            // init mouse move listener
            // write coordinates to mouse_pos_buffer
            let magic_square = magic_square.clone();
            let mouse_pos_buffer = mouse_pos_buffer.clone();
            let canvas = canvas.clone();
            let context: web_sys::WebGl2RenderingContext = MagicSquare::context(&canvas).unwrap();
            let closure = Closure::<dyn FnMut(_)>::new(move |event: web_sys::MouseEvent| {
                context.clear_color(0.0, 0.0, 0.0, 0.0);
                context.clear(WebGl2RenderingContext::COLOR_BUFFER_BIT);
                mouse_pos_buffer.clone().borrow_mut()[0] = MagicSquare::clip_x(event.offset_x(), width);
                mouse_pos_buffer.clone().borrow_mut()[1] = MagicSquare::clip_x(event.offset_y(), height);
                // mouse_pos_buffer[0] = MagicSquare::clip_x(event.offset_x(), width);
                // mouse_pos_buffer[1] = MagicSquare::clip_y(event.offset_y(), height);
                magic_square.dispatch_event(&web_sys::Event::new("render").unwrap()).unwrap();
            });

            canvas.add_event_listener_with_callback(
                "mousemove", 
                closure.as_ref().unchecked_ref()
            )?;
            closure.forget();
        }

        {
            let magic_square = magic_square.clone();
            // let mouse_pos_buffer = &mouse_pos_buffer;
            let geometry_cache = geometry_cache.clone();
            let mouse_pos_buffer = mouse_pos_buffer.clone();
            let ui_buffer = ui_buffer.clone();

            // set up render listener
            let closure = Closure::<dyn FnMut(_)>::new(move |_event: web_sys::Event| {
                MagicSquare::render_all_lines(
                    &mouse_pos_buffer,
                    &ui_buffer.clone().borrow(), 
                    &mut color_idx_offset_delay, 
                    &geometry_cache,
                );
            });

            magic_square.add_event_listener_with_callback(
                "render",
                closure.as_ref().unchecked_ref()
            )?;
            closure.forget();
        }

        
        let animation_id: i32;
        {
            // set up animation loop
            let geometry_cache = geometry_cache.clone();
            let ui_buffer = ui_buffer.clone();
            let max_idx = Settings::max_idx_from_draw_pattern(ui_buffer.borrow().settings.draw_pattern);

            // closures used to allocate and clean up resources
            let f: Rc<RefCell<Option<wasm_bindgen::prelude::Closure<_>> >> = Rc::new(RefCell::new(None));
            let g = f.clone();

            *g.borrow_mut() = Some(Closure::new(move || {
                // TODO SOON: DEBUG cancelAnimationFrame
                // seemingly passing valid requestId back to js
                // but passing the id to cancleAnimationFrame in onDestroy callback
                // does not seem to be cancelling the animation frame
                // thus we can wind up with multiple (performance killing) instances
                // by navigating back and forth
                for idx in 0..max_idx {
                    MagicSquare::render(
                        geometry_cache.borrow().gl_vertices(idx), 
                        &geometry_cache.borrow().rgbas[idx], 
                        &context
                    ).expect("Render error");
                }
                MagicSquare::request_animation_frame(f.borrow().as_ref().unwrap());
            }));

            animation_id = MagicSquare::request_animation_frame(g.borrow().as_ref().unwrap());
        }

        Ok(animation_id)
    }
}

pub type Rgba = [f32; 4];

impl MagicSquare {
    pub fn clip_x(offset_x: i32, width: i32) -> f32 {
        // x coordinate of mouse position in clip space
        (2.0 * (offset_x as f32) / width as f32) - 1.0
    }

    pub fn clip_y(offset_y: i32, height: i32) -> f32 {
        // y coordinate of mouse position in clip space
        1.0 - ((2.0 * offset_y as f32) / height as f32)
    }

    pub fn request_animation_frame(f: &Closure<dyn FnMut()>) -> i32 {
        MagicSquare::window()
            .request_animation_frame(f.as_ref().unchecked_ref())
            .expect("should register `requestAnimationFrame` OK")
    }

    fn render_all_lines(
        mouse_pos_buffer: &Rc<RefCell<[f32; 2]>>,
        ui_buffer: &UiBuffer,
        // geometry_cache: Arc<Mutex<GeometryCache>>,
        color_idx_offset_delay: &mut [usize; 2],
        geometry_cache: &Rc<RefCell<GeometryCache>>,
    ) {
        let max_idx = Settings::max_idx_from_draw_pattern(ui_buffer.settings.draw_pattern);
        let mouse_pos_buffer = *mouse_pos_buffer.clone().borrow();
        for idx in 0..max_idx { // geometry_cache.max_idx + 1 { //TODO: settings.cache_per
            let rot_seq = RotationSequence::new(
                Rotation::new(
                    Axis::X, 
                    mouse_pos_buffer[0] * ui_buffer.settings.x_axis_x_rot_coeff
                        + mouse_pos_buffer[1] * ui_buffer.settings.y_axis_x_rot_coeff
                        + idx as f32 * ui_buffer.settings.x_rot_spread
                ),
                Rotation::new(
                    Axis::Y,
                    mouse_pos_buffer[0] * ui_buffer.settings.x_axis_y_rot_coeff
                        + mouse_pos_buffer[1] * ui_buffer.settings.y_axis_y_rot_coeff
                        + idx as f32 * ui_buffer.settings.y_rot_spread
                ),
                Rotation::new(
                    Axis::Z,
                    mouse_pos_buffer[0] + ui_buffer.settings.x_axis_z_rot_coeff
                        + mouse_pos_buffer[1] + ui_buffer.settings.y_axis_z_rot_coeff
                        + idx as f32 * ui_buffer.settings.z_rot_spread
                ),
            );

            let translation = match ui_buffer.settings.mouse_tracking {
                MouseTracking::On => Translation { x: mouse_pos_buffer[0], y: mouse_pos_buffer[1], z: 0.0 },
                MouseTracking::Off => Translation { x: 0.0, y: 0.0, z: 0.0 },
                MouseTracking::InvX =>  Translation { x: - mouse_pos_buffer[0], y: mouse_pos_buffer[1], z: 0.0 },
                MouseTracking::InvY =>  Translation { x: mouse_pos_buffer[0], y: - mouse_pos_buffer[1], z: 0.0 },
                MouseTracking::InvXY =>  Translation { x: - mouse_pos_buffer[0], y: - mouse_pos_buffer[1], z: 0.0 },
            };

                // let hexagon = Geometry::hexagon(
                //     ui_buffer.settings.radius_step * idx as f32 + ui_buffer.settings.radius_min, 
                //     rot_seq,
                //     translation
                // );

            // let _ = Worker::spawn(move || {
                let icosohedron = Geometry::icosohedron(
                    ui_buffer.settings.radius_step * idx as f32 + ui_buffer.settings.radius_min, 
                    rot_seq,
                    translation
                );
                
                let color_idx_offset: usize = color_idx_offset_delay[0];
                let color_idx_delay: usize = color_idx_offset_delay[1];
                let rgba = MagicSquare::get_rgba(
                    mouse_pos_buffer.clone(),
                    ui_buffer, 
                    (idx + color_idx_offset) % 8
                );
                color_idx_offset_delay[1] = color_idx_delay + 1;
                if color_idx_delay == 50 {
                    color_idx_offset_delay[0] = color_idx_offset + 1_usize % 8;
                    color_idx_offset_delay[1] = 0;
                }
                geometry_cache.borrow_mut().set_next(icosohedron.arr, rgba, Shape::Icosohedron, max_idx);
            // });
        }


    }

    fn get_rgba(mouse_pos_buffer: [f32; 2], ui_buffer: &UiBuffer, idx: usize) -> Rgba {
        // let mut result: Rgba = [0.0, 0.0, 0.0, 0.0];
        // result[0] = ui_buffer.settings.color_1[0] / 255.0;// 1.0 - mouse_pos_buffer[0];
        // result[1] = ui_buffer.settings.color_1[1] / 255.0;// 1.0 - mouse_pos_buffer[1];
        // result[2] = ui_buffer.settings.color_1[2] / 255.0; // 1.0 - (idx as f32 / CACHE_CAPACITY as f32);
        // result[3] = ui_buffer.settings.color_1[3] / 255.0; // 0.1 * idx as f32;
        // result
        //
        // log(&format!("{idx}"));
        match idx {
            0 => ui_buffer.settings.color_1,
            1 => ui_buffer.settings.color_2,
            2 => ui_buffer.settings.color_3,
            3 => ui_buffer.settings.color_4,
            4 => ui_buffer.settings.color_5,
            5 => ui_buffer.settings.color_6,
            6 => ui_buffer.settings.color_7,
            7 => ui_buffer.settings.color_8,
            8 => ui_buffer.settings.color_1,
            9 => ui_buffer.settings.color_2,
            10 => ui_buffer.settings.color_3,
            11 => ui_buffer.settings.color_4,
            12 => ui_buffer.settings.color_5,
            13 => ui_buffer.settings.color_6,
            14 => ui_buffer.settings.color_7,
            15 => ui_buffer.settings.color_8,
            _ => ui_buffer.settings.color_1
        }
    }

    fn window() -> web_sys::Window {
        web_sys::window().expect("no global `window` exists")
    }

    pub fn document() -> web_sys::Document {
        MagicSquare::window()
            .document()
            .expect("should have a document on window")
    }

    pub fn magic_square() -> web_sys::Element {
        MagicSquare::document()
            .get_element_by_id("magic_square")
            .expect("unable to find magic_square element")
    }

    fn canvas() -> web_sys::Element {
        MagicSquare::document()
            .get_element_by_id("magic_square_canvas")
            .expect("unable to find canvas element")
    }

    fn form() -> web_sys::Element {
        MagicSquare::document()
            .get_element_by_id("magic_square_control")
            .expect("unable to find control element")
    }

    pub fn context(
        canvas: &web_sys::HtmlCanvasElement,
    ) -> Result<web_sys::WebGl2RenderingContext, JsValue> {
        let context = canvas
            .get_context("webgl2")?
            .expect("unable to get webgl2 context")
            .dyn_into::<WebGl2RenderingContext>()?;

        Ok(context)
    }

    fn draw(context: &WebGl2RenderingContext, vert_count: i32) {
        context.draw_arrays(WebGl2RenderingContext::LINES, 0, vert_count);
    }

    fn render(
        vertices: &[f32],
        color: &Rgba,
        context: &web_sys::WebGl2RenderingContext,
    ) -> Result<(), JsValue> {
        let vert_shader = ShaderCompiler::vert_default(context).unwrap();
        let frag_shader = ShaderCompiler::frag_default(context, color).unwrap();

        let program = ProgramLinker::exec(context, &vert_shader, &frag_shader)?;
        context.use_program(Some(&program));

        let position_attribute_location = context.get_attrib_location(&program, "position");
        let buffer = context.create_buffer().ok_or("Failed to create buffer")?;
        context.bind_buffer(WebGl2RenderingContext::ARRAY_BUFFER, Some(&buffer));

        // Note that `Float32Array::view` is somewhat dangerous (hence the
        // `unsafe`!). This is creating a raw view into our module's
        // `WebAssembly.Memory` buffer, but if we allocate more pages for ourself
        // (aka do a memory allocation in Rust) it'll cause the buffer to change,
        // causing the `Float32Array` to be invalid.
        //
        // As a result, after `Float32Array::view` we have to be very careful not to
        // do any memory allocations before it's dropped.
        unsafe {
            let positions_array_buf_view = js_sys::Float32Array::view(vertices);

            context.buffer_data_with_array_buffer_view(
                WebGl2RenderingContext::ARRAY_BUFFER,
                &positions_array_buf_view,
                WebGl2RenderingContext::STATIC_DRAW,
            );
        }

        let vao = context
            .create_vertex_array()
            .ok_or("Could not create vertex array object")?;

        context.bind_vertex_array(Some(&vao));

        context.vertex_attrib_pointer_with_i32(
            position_attribute_location as u32,
            3,
            WebGl2RenderingContext::FLOAT,
            false,
            0,
            0,
        );
        context.enable_vertex_attrib_array(position_attribute_location as u32);

        context.bind_vertex_array(Some(&vao));

        let vert_count = (vertices.len() / 3) as i32;
        MagicSquare::draw(&context, vert_count);

        Ok(())
    }
}


