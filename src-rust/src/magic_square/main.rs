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
use crate::magic_square::lfo::Lfo;

use super::settings::{ColorDirection, MouseTracking, Settings};
use super::transformations::Transformation;
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
    pub async fn run(prev_settings: JsValue) -> JsValue {
        // testing multithreading
        //
        // let (to_worker, from_main) = std::sync::mpsc::channel();
        // let (to_main, from_worker) = std::sync::mpsc::channel();
        // Worker::spawn(move || { to_main.send(from_main.recv().unwrap() + 1.0); });
        // to_worker.send(1.0);
        // assert_eq!(from_worker.recv().unwrap(), 2.0);
        
        let ui_buffer: UiBuffer = match serde_wasm_bindgen::from_value(prev_settings){
            Ok(res) => {
                log("SUCCESSFUL SETTINGS PARSE");
                // log(&format!("{:?}", res));
                UiBuffer::from_prev_settings(res)
            },
            Err(_e) => {
                // log(&format!("{:?}", e));
                UiBuffer::new()
            }
        };
        let ui_buffer = Rc::new(RefCell::new(ui_buffer));

        let magic_square = MagicSquare::magic_square(); // awesome naming, great job!
        let magic_square = Rc::new(magic_square);
        let canvas = MagicSquare::canvas().dyn_into::<web_sys::HtmlCanvasElement>().unwrap();
        let canvas = Rc::new(canvas);

        // triggers cleanup requestAnimationFrame closure
        let destroy_flag: bool = false;
        let destroy_flag: Rc<RefCell<bool>> = Rc::new(RefCell::new(destroy_flag));

        let height:i32 = canvas.client_height();
        let width:i32 = canvas.client_width();
        // incriment idx_delay each render
        // when idx_delay reaches a desired delay value
        // incriment idx_offset
        let mut color_idx_offset_delay: [u8; 2] = [0, 0];
        //Arc::new(Mutex::new(
        let geometry_cache = GeometryCache::new(
                26, 
                [[0.0; 300]; CACHE_CAPACITY], 
                [Shape::None; CACHE_CAPACITY]
            );

        let geometry_cache = Rc::new(RefCell::new(geometry_cache));
        // )); 
        
        let frag_shader_cache: Vec<String> = vec![
            ShaderCompiler::into_frag_shader_string(&ui_buffer.clone().borrow().settings.color_1),
            ShaderCompiler::into_frag_shader_string(&ui_buffer.clone().borrow().settings.color_2),
            ShaderCompiler::into_frag_shader_string(&ui_buffer.clone().borrow().settings.color_3),
            ShaderCompiler::into_frag_shader_string(&ui_buffer.clone().borrow().settings.color_4),
            ShaderCompiler::into_frag_shader_string(&ui_buffer.clone().borrow().settings.color_5),
            ShaderCompiler::into_frag_shader_string(&ui_buffer.clone().borrow().settings.color_6),
            ShaderCompiler::into_frag_shader_string(&ui_buffer.clone().borrow().settings.color_7),
            ShaderCompiler::into_frag_shader_string(&ui_buffer.clone().borrow().settings.color_8)
        ];
        let frag_shader_cache: Rc<RefCell<Vec<String>>> = Rc::new(RefCell::new(frag_shader_cache));

        let form = MagicSquare::form();
        let form = Rc::new(form);

        let mouse_pos_buffer: [f32; 2] = [0.0, 0.0];
        let mouse_pos_buffer: Rc<RefCell<[f32; 2]>> = Rc::new(RefCell::new(mouse_pos_buffer));

        let context: web_sys::WebGl2RenderingContext = MagicSquare::context(&canvas).unwrap();
        context.clear_color(0.0, 0.0, 0.0, 0.0);
        context.clear(WebGl2RenderingContext::COLOR_BUFFER_BIT);

        {
            // init destroy listener on app_main
            // onDestroy hook in Main.svelte dispatches destroymswasm event
            // this closure flips destroy_flag
            // requestAnimationFrame checks value, cleans up resources

            let app_main = MagicSquare::app_main();
            let destroy_flag = destroy_flag.clone();
            let closure = Closure::<dyn FnMut(_)>::new(move |_event: web_sys::Event| {
                *destroy_flag.clone().borrow_mut() = true;
            });

            app_main.add_event_listener_with_callback(
                "destroymswasm",
                closure.as_ref().unchecked_ref()
            ).unwrap();

            closure.forget();
        }

        {
            // init UI control settings listener
            let form = form.clone();
            let ui_buffer = ui_buffer.clone();
            let frag_shader_cache = frag_shader_cache.clone();
            let geometry_cache = geometry_cache.clone();

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
                    ui_buffer
                        .clone()
                        .borrow_mut()
                        .update(
                            id, 
                            val, 
                            &mut *frag_shader_cache.clone().borrow_mut(),
                            &mut *geometry_cache.clone().borrow_mut()
                    );
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
                magic_square.dispatch_event(&web_sys::Event::new("render").unwrap()).unwrap();
            });

            canvas.add_event_listener_with_callback(
                "mousemove", 
                closure.as_ref().unchecked_ref()
            ).unwrap();
            closure.forget();
        }

        {
            // set up animation loop
            let geometry_cache = geometry_cache.clone();
            let ui_buffer = ui_buffer.clone();
            let frag_shader_cache = frag_shader_cache.clone();

            // let performance = MagicSquare::performance();
            let mut x: f32 = -3.14159;
            
            let max_idx = Settings::max_idx_from_draw_pattern(ui_buffer.borrow().settings.draw_pattern);
            let destroy_flag = destroy_flag.clone();

            // closures used to allocate and clean up resources
            let f: Rc<RefCell<Option<wasm_bindgen::prelude::Closure<_>> >> = Rc::new(RefCell::new(None));
            let g = f.clone();

            *g.borrow_mut() = Some(Closure::new(move || {
                if *destroy_flag.clone().borrow() {
                    // cleanup resource
                    let _ = f.borrow_mut().take();
                    return;
                } else {
                    // TODO: consider applying LFO per cache slot
                    // could allow for some cool snake-like movement
                    let lfo_1 = Lfo::new(
                        ui_buffer.borrow().settings.lfo_1_active,
                        ui_buffer.borrow().settings.lfo_1_amp,
                        ui_buffer.borrow().settings.lfo_1_dest,
                        ui_buffer.borrow().settings.lfo_1_freq,
                        ui_buffer.borrow().settings.lfo_1_phase,
                        ui_buffer.borrow().settings.lfo_1_shape,
                    );

                    let lfo_2 = Lfo::new(
                        ui_buffer.borrow().settings.lfo_2_active,
                        ui_buffer.borrow().settings.lfo_2_amp,
                        ui_buffer.borrow().settings.lfo_2_dest,
                        ui_buffer.borrow().settings.lfo_2_freq,
                        ui_buffer.borrow().settings.lfo_2_phase,
                        ui_buffer.borrow().settings.lfo_2_shape,
                    );

                    let lfo_3 = Lfo::new(
                        ui_buffer.borrow().settings.lfo_3_active,
                        ui_buffer.borrow().settings.lfo_3_amp,
                        ui_buffer.borrow().settings.lfo_3_dest,
                        ui_buffer.borrow().settings.lfo_3_freq,
                        ui_buffer.borrow().settings.lfo_3_phase,
                        ui_buffer.borrow().settings.lfo_3_shape,
                    );

                    let lfo_4 = Lfo::new(
                        ui_buffer.borrow().settings.lfo_4_active,
                        ui_buffer.borrow().settings.lfo_4_amp,
                        ui_buffer.borrow().settings.lfo_4_dest,
                        ui_buffer.borrow().settings.lfo_4_freq,
                        ui_buffer.borrow().settings.lfo_4_phase,
                        ui_buffer.borrow().settings.lfo_4_shape,
                    );

                    // let start: f64 = performance.now();
                    // let val: f32 = lfo_1.eval(x);
                    // log(&format!("{x}, {val}"));
                    x = x + 0.001;
                    if x == 3.142 {
                        x = -3.142;
                    }
                    
                    // harvest current ui_buffer for computation
                    let mut ui_buffer = (*ui_buffer.clone().borrow()).clone();
                    ui_buffer = ui_buffer.copy();
                    lfo_1.modify(x, &mut ui_buffer);
                    lfo_2.modify(x, &mut ui_buffer);
                    lfo_3.modify(x, &mut ui_buffer);
                    lfo_4.modify(x, &mut ui_buffer);

                    let delay_reset: u8 = std::cmp::max(22 - ui_buffer.settings.color_speed, 1);

                    if color_idx_offset_delay[1] > delay_reset {
                        // this can happen when user changes ui_buffer.settings.color_speed 
                        // new speed will eventually kick in anyway
                        // but this makes it immediate
                        color_idx_offset_delay[1] = 0
                    }
                    let color_idx_offset: u8 = color_idx_offset_delay[0];
                    let color_idx_delay: u8 = color_idx_offset_delay[1];
                    
                    // 0 < color_speed < 21
                    if color_idx_delay == delay_reset {
                        color_idx_offset_delay[0] = match ui_buffer.settings.color_direction {
                            ColorDirection::In => (color_idx_offset + 1) % 8,
                            ColorDirection::Fix => color_idx_offset,
                            ColorDirection::Out => (color_idx_offset - 1) % 8,
                        };
                        color_idx_offset_delay[1] = 0;
                    }
                    
                    color_idx_offset_delay[1] = color_idx_offset_delay[1] + 1;

                    // compute
                    MagicSquare::render_all_lines(
                        &mouse_pos_buffer,
                        &ui_buffer, 
                        &geometry_cache,
                    );
                    
                    // display
                    for idx in 0..max_idx+1 {
                        match MagicSquare::render(
                            geometry_cache.borrow().gl_vertices(idx), 
                            &*frag_shader_cache.clone().borrow()[(idx + color_idx_offset_delay[0] as usize) % 8], 
                            &context
                        ) {
                            Ok(_) =>  {}, // log("SUCCESSFUL RENDER"),
                            Err(_e) => {}// log(&format!("{:?}", e))  
                        };
                    }

                    MagicSquare::request_animation_frame(f.borrow().as_ref().unwrap());
                }
            }));

            MagicSquare::request_animation_frame(g.borrow().as_ref().unwrap());
        }
        
        let to_js = ui_buffer.clone().borrow().clone();
        serde_wasm_bindgen::to_value(&to_js).unwrap()
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
        geometry_cache: &Rc<RefCell<GeometryCache>>,
    ) {
        let max_idx = Settings::max_idx_from_draw_pattern(ui_buffer.settings.draw_pattern);
        let mouse_pos_buffer = *mouse_pos_buffer.clone().borrow();

        for idx in 0..max_idx { // geometry_cache.max_idx + 1 { //TODO: settings.cache_per
            let idx_f32 = idx as f32;
            let rot_seq = RotationSequence::new(
                Rotation::new(
                    Axis::X, 
                    ui_buffer.settings.x_rot_base
                        + (mouse_pos_buffer[0] 
                            + ui_buffer.settings.translation_x_base) * ui_buffer.settings.x_axis_x_rot_coeff
                        + (mouse_pos_buffer[1] 
                            + ui_buffer.settings.translation_y_base) * ui_buffer.settings.y_axis_x_rot_coeff
                        + idx_f32 * ui_buffer.settings.x_rot_spread
                ),
                Rotation::new(
                    Axis::Y,
                    ui_buffer.settings.y_rot_base
                        + (mouse_pos_buffer[0] 
                            + ui_buffer.settings.translation_x_base) * ui_buffer.settings.x_axis_y_rot_coeff
                        + (mouse_pos_buffer[1] 
                            + ui_buffer.settings.translation_y_base) * ui_buffer.settings.y_axis_y_rot_coeff
                        + idx_f32 * ui_buffer.settings.y_rot_spread
                ),
                Rotation::new(
                    Axis::Z,
                    ui_buffer.settings.z_rot_base
                        + (mouse_pos_buffer[0] 
                            + ui_buffer.settings.translation_x_base) * ui_buffer.settings.x_axis_z_rot_coeff
                        + (mouse_pos_buffer[1] 
                            + ui_buffer.settings.translation_y_base) * ui_buffer.settings.y_axis_z_rot_coeff
                        + idx_f32 * ui_buffer.settings.z_rot_spread
                ),
            );

            let translation = match ui_buffer.settings.mouse_tracking {
                MouseTracking::On => Translation { 
                    x: ui_buffer.settings.translation_x_base
                        + (idx_f32 * ui_buffer.settings.translation_x_spread)
                        + mouse_pos_buffer[0], 
                    y: ui_buffer.settings.translation_y_base 
                        - (idx_f32 * ui_buffer.settings.translation_y_spread)
                        - mouse_pos_buffer[1], 
                    z: ui_buffer.settings.translation_z_base 
                        + (idx_f32 * ui_buffer.settings.translation_z_spread)
                },
                MouseTracking::Off => Translation { 
                    x: ui_buffer.settings.translation_x_base
                        + (idx_f32 * ui_buffer.settings.translation_x_spread), 
                    y: ui_buffer.settings.translation_y_base
                        - (idx_f32 * ui_buffer.settings.translation_y_spread), 
                    z: ui_buffer.settings.translation_z_base
                        + (idx_f32 * ui_buffer.settings.translation_z_spread)
                },
                MouseTracking::InvX =>  Translation { 
                    x: ui_buffer.settings.translation_x_base 
                        + (idx_f32 * ui_buffer.settings.translation_x_spread)
                        - mouse_pos_buffer[0], 
                    y: ui_buffer.settings.translation_y_base
                        - (idx_f32 * ui_buffer.settings.translation_y_spread)
                        - mouse_pos_buffer[1], 
                    z: ui_buffer.settings.translation_z_base
                        + (idx_f32 * ui_buffer.settings.translation_z_spread)
                },
                MouseTracking::InvY =>  Translation { 
                    x: ui_buffer.settings.translation_x_base 
                        + (idx_f32 * ui_buffer.settings.translation_x_spread)
                        + mouse_pos_buffer[0], 
                    y: ui_buffer.settings.translation_y_base 
                        - (idx_f32 * ui_buffer.settings.translation_y_spread)
                        + mouse_pos_buffer[1], 
                    z: ui_buffer.settings.translation_z_base
                        + (idx_f32 * ui_buffer.settings.translation_z_spread)
                },
                MouseTracking::InvXY =>  Translation { 
                    x: ui_buffer.settings.translation_x_base
                        + (idx_f32 * ui_buffer.settings.translation_x_spread)
                        - mouse_pos_buffer[0], 
                    y: ui_buffer.settings.translation_y_base
                        - (idx_f32 * ui_buffer.settings.translation_y_spread)
                        + mouse_pos_buffer[1], 
                    z: ui_buffer.settings.translation_z_base
                        + (idx_f32 * ui_buffer.settings.translation_z_spread)
                },
            };

                // let hexagon = Geometry::hexagon(
                //     ui_buffer.settings.radius_step * idx as f32 + ui_buffer.settings.radius_base, 
                //     rot_seq,
                //     translation
                // );

            // let _ = Worker::spawn(move || {
            let icosohedron = Geometry::icosohedron(
                ui_buffer.settings.radius_step * idx_f32 + ui_buffer.settings.radius_base, 
                Transformation { 
                    order: ui_buffer.settings.transform_order,
                    rot_seq,
                    translation
                }
            );
            
            geometry_cache.borrow_mut().set_next(icosohedron.arr, Shape::Icosohedron, max_idx);
            // });
        }
    }

    fn window() -> web_sys::Window {
        web_sys::window().expect("no global `window` exists")
    }

    // fn performance() -> web_sys::Performance {
    //     MagicSquare::window()
    //         .performance()
    //         .expect("performance should be available")
    // }

    pub fn document() -> web_sys::Document {
        MagicSquare::window()
            .document()
            .expect("should have a document on window")
    }

    pub fn app_main() -> web_sys::Element {
        MagicSquare::document()
            .get_element_by_id("app_main")
            .expect("unable to find app_main element")
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
            .get_element_by_id("magic_square_control_rack")
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
        frag_shader_str: &str,
        context: &web_sys::WebGl2RenderingContext,
    ) -> Result<(), JsValue> {
        // TODO
        let vert_shader = ShaderCompiler::vert_default(context)?;
        let frag_shader = ShaderCompiler::frag_default(context, frag_shader_str)?;

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


