use std::rc::Rc;
use std::cell::RefCell;
use wasm_bindgen::prelude::*;
use web_sys::{WebGl2RenderingContext, WebGlProgram};
use crate::magic_square::geometry::cache::{Cache as GeometryCache, CACHE_CAPACITY};
use crate::magic_square::ui_buffer::UiBuffer;
use crate::magic_square::lfo::Lfo;
use super::gl_draw::GlDraw;
use super::gl_program::GlProgram;
use super::gl_uniforms::{GlUniforms, UniformLocations};
use super::settings::ColorDirection;
use super::geometry::geom::Geom;

#[wasm_bindgen]
extern "C" {
    #[wasm_bindgen(js_namespace = console)]
    pub fn log(s: &str);

    #[wasm_bindgen(js_name = "performance")]
    pub static PERFORMANCE: web_sys::Performance;
}

#[derive(Clone, Copy, Debug)]
pub enum Axis {
    X,
    Y,
    Z
}

#[wasm_bindgen]
pub struct MagicSquare;



#[wasm_bindgen]
impl MagicSquare {
    // Entry point into Rust WASM from JS
    // https://rustwasm.github.io/wasm-bindgen/examples/webgl.html
    pub async fn run(settings: JsValue, presets: JsValue) -> JsValue {
        let ui_buffer = UiBuffer::from(settings, presets);

        let canvas = MagicSquare::canvas().dyn_into::<web_sys::HtmlCanvasElement>().unwrap();
        let canvas = Rc::new(canvas);

        let geometry_cache: GeometryCache = GeometryCache::new(&ui_buffer.settings.shapes);
        let geometry_cache = Rc::new(RefCell::new(geometry_cache));

        // wrap ui_buffer to be accessed by different closures
        let ui_buffer = Rc::new(RefCell::new(ui_buffer));

        let magic_square = MagicSquare::magic_square(); // awesome naming, great job!
        let magic_square = Rc::new(magic_square);

        // triggers cleanup requestAnimationFrame closure
        let destroy_flag: bool = false;
        let destroy_flag: Rc<RefCell<bool>> = Rc::new(RefCell::new(destroy_flag));

        let height:i32 = canvas.client_height();
        let width:i32 = canvas.client_width();
        // incriment idx_delay each render
        // when idx_delay reaches a desired delay value
        // incriment idx_offset
        let mut color_idx_offset_delay: [u8; 2] = [0, 0];

        let form = MagicSquare::form();
        let form = Rc::new(form);

        let mouse_pos_buffer: [f32; 2] = [0.0, 0.0];
        let mouse_pos_buffer: Rc<RefCell<[f32; 2]>> = Rc::new(RefCell::new(mouse_pos_buffer));

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
                            &mut *geometry_cache.clone().borrow_mut(),
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

        // log("index out of bounds hunt 1");

        {
            // set up animation loop
            let geometry_cache = geometry_cache.clone();
            let ui_buffer = ui_buffer.clone();
            let mut animation_idx: usize = 0;

            // let performance = MagicSquare::performance();
            let mut x: f32 = -3.14159;
            
            let destroy_flag = destroy_flag.clone();

            // closures used to allocate and clean up resources
            let f: Rc<RefCell<Option<wasm_bindgen::prelude::Closure<_>> >> = Rc::new(RefCell::new(None));
            let g = f.clone();

            let mut frame_counter: usize = 0;

            // set up WebGL
            let gl: web_sys::WebGl2RenderingContext = MagicSquare::context(&canvas).unwrap();
            gl.clear_color(0.0, 0.0, 0.0, 0.0);
            gl.clear(WebGl2RenderingContext::COLOR_BUFFER_BIT);

            let program: WebGlProgram =  GlProgram::new(&gl).expect(&format!("ISSUE INIT GL_PROGRAM"));
            gl.use_program(Some(&program));
            let uniform_locations = UniformLocations::new(&gl, &program);
            let mut uniforms = GlUniforms::new();

            let gl_buffer = gl.create_buffer().ok_or("Failed to create buffer").unwrap();
            gl.bind_buffer(WebGl2RenderingContext::ARRAY_BUFFER, Some(&gl_buffer));

            // log("index out of bounds hunt 2");
            // set gl to read vertex datt from geometry_cache.vertices

            // Note that `Float32Array::view` is somewhat dangerous (hence the
            // `unsafe`!). This is creating a raw view into our module's
            // `WebAssembly.Memory` buffer, but if we allocate more pages for ourself
            // (aka do a memory allocation in Rust) it'll cause the buffer to change,
            // causing the `Float32Array` to be invalid.
            //
            // As a result, after `Float32Array::view` we have to be very careful not to
            // do any memory allocations before it's dropped.
            let vertices = Geom::f32_array();
            unsafe {
                let positions_array_buf_view = js_sys::Float32Array::view(&vertices);

                gl.buffer_data_with_array_buffer_view(
                    WebGl2RenderingContext::ARRAY_BUFFER,
                    &positions_array_buf_view,
                    WebGl2RenderingContext::STATIC_DRAW,
                );
            }

            let vao = gl.create_vertex_array().ok_or("Could not create vertex array object").unwrap();
            gl.bind_vertex_array(Some(&vao));

            let position_attribute_location = gl.get_attrib_location(&program, "position");
            gl.vertex_attrib_pointer_with_i32(
                position_attribute_location as u32,
                3,
                WebGl2RenderingContext::FLOAT,
                false,
                0,
                0,
            );
            gl.enable_vertex_attrib_array(position_attribute_location as u32);
            gl.bind_vertex_array(Some(&vao));
            
            *g.borrow_mut() = Some(Closure::new(move || {
                let mut settings = ui_buffer.clone().borrow().settings;

                if *destroy_flag.clone().borrow() {
                    // cleanup resource
                    let _ = f.borrow_mut().take();
                    return;
                } else {
                    // TODO: consider applying LFO per cache slot
                    // could allow for some cool snake-like movement
                    let lfo_1 = Lfo::new(
                        settings.lfo_1_active,
                        settings.lfo_1_amp,
                        settings.lfo_1_dest,
                        settings.lfo_1_freq,
                        settings.lfo_1_phase,
                        settings.lfo_1_shape,
                    );

                    let lfo_2 = Lfo::new(
                        settings.lfo_2_active,
                        settings.lfo_2_amp,
                        settings.lfo_2_dest,
                        settings.lfo_2_freq,
                        settings.lfo_2_phase,
                        settings.lfo_2_shape,
                    );

                    let lfo_3 = Lfo::new(
                        settings.lfo_3_active,
                        settings.lfo_3_amp,
                        settings.lfo_3_dest,
                        settings.lfo_3_freq,
                        settings.lfo_3_phase,
                        settings.lfo_3_shape,
                    );

                    let lfo_4 = Lfo::new(
                        settings.lfo_4_active,
                        settings.lfo_4_amp,
                        settings.lfo_4_dest,
                        settings.lfo_4_freq,
                        settings.lfo_4_phase,
                        settings.lfo_4_shape,
                    );

                    // let start: f64 = performance.now();
                    // let val: f32 = lfo_1.eval(x);
                    // log(&format!("{x}, {val}"));
                    x = x + 0.001;
                    if x == 3.142 {
                        x = -3.142;
                    }
                    
                    // harvest current ui_buffer for computation
                    lfo_1.modify(x, &mut settings);
                    lfo_2.modify(x, &mut settings);
                    lfo_3.modify(x, &mut settings);
                    lfo_4.modify(x, &mut settings);


                    let delay_reset: u8 = std::cmp::max(22 - settings.color_speed, 1);

                    if color_idx_offset_delay[1] > delay_reset {
                        // this can happen when user changes settings.color_speed 
                        // new speed will eventually kick in anyway
                        // but this makes it immediate
                        color_idx_offset_delay[1] = 0
                    }
                    let color_idx_offset: u8 = color_idx_offset_delay[0];
                    let color_idx_delay: u8 = color_idx_offset_delay[1];
                    
                    // log("index out of bounds hunt 3");
                    // 0 < color_speed < 21
                    if color_idx_delay == delay_reset {
                        color_idx_offset_delay[0] = match settings.color_direction {
                            ColorDirection::In => (color_idx_offset + 1) % CACHE_CAPACITY as u8,
                            ColorDirection::Fix => color_idx_offset,
                            ColorDirection::Out => (color_idx_offset - 1) % CACHE_CAPACITY as u8,
                        };
                        color_idx_offset_delay[1] = 0;
                    }
                    
                    color_idx_offset_delay[1] = color_idx_offset_delay[1] + 1;
                    
                    // compute
                    uniforms.set_uniforms(&mouse_pos_buffer, &settings);
                    // log(&format!("{:?}", uniforms));
                    
                    // draw
                    let shapes = geometry_cache.clone().borrow().shapes;
                    // log(&format!("{:?}", geometry_cache.clone().borrow().vertices));
                    if let Err(_) = GlDraw::scene(
                        &gl,
                        &uniforms,
                        &uniform_locations,
                        &shapes,
                        settings.transform_order
                    ) {
                        log("DRAW ERROR");
                    }
                    
                    let frame_counter_limit: i32 = if settings.draw_pattern_speed > 19 { 1 } else { settings.draw_pattern_speed };

                    if frame_counter % frame_counter_limit as usize == 0 {
                        animation_idx = (animation_idx  + 1) % CACHE_CAPACITY;
                    }
                    frame_counter = (frame_counter + 1) % (frame_counter_limit as usize);
                    // log("index out of bounds hunt 4");
                    MagicSquare::request_animation_frame(f.borrow().as_ref().unwrap());
                }
            }));

            MagicSquare::request_animation_frame(g.borrow().as_ref().unwrap());
        }
        
        let to_js = ui_buffer.clone().borrow().clone().settings;
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
}
