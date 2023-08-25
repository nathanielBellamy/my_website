// use bytemuck::{bytes_of, from_bytes}
use wasm_bindgen::JsValue;
// use crate::log;
use crate::magic_square::animation::Animation;
use crate::magic_square::geometry::geom::Geom;
use crate::magic_square::gl_draw::GlDraw;
use crate::magic_square::gl_program::GlProgram;
use crate::magic_square::gl_uniforms::{GlUniforms, UniformLocations};
use crate::magic_square::settings::ColorDirection;
use crate::magic_square::geometry::cache::{Cache as GeometryCache, CACHE_CAPACITY};
use crate::magic_square::lfo::Lfo;
use crate::public_square::ui_buffer::UiBuffer;
use std::cell::RefCell;
use std::rc::Rc;
use wasm_bindgen::prelude::*;
use web_sys::{MessageEvent, WebGl2RenderingContext, WebGlProgram, WebSocket};
use crate::magic_square::{settings::Settings, main::MagicSquare, main::X_MAX};

#[derive(Debug)]
pub struct PublicSquare;

#[wasm_bindgen]
struct PubSq;

#[wasm_bindgen]
impl PubSq {
    #[allow(unused)] // called from js
    pub async fn run(base_url: JsValue, set_all_settings: &js_sys::Function, touch_screen: JsValue) -> JsValue {
        let touch_screen: bool = serde_wasm_bindgen::from_value(touch_screen).unwrap();
        
        let base_url: String = serde_wasm_bindgen::from_value(base_url).unwrap();
        let mut protocol: String;
        if base_url == "localhost:8080".to_string() {
            protocol = "ws".to_string();
        } else {
            protocol = "wss".to_string();
        }
        let url: String = format!("{protocol}://{base_url}/public-square-wasm-ws");
        // setup websocket
        let ws: WebSocket;
        match WebSocket::new(url.as_str()) {
            Ok(socket) => ws = socket,
            Err(_) => {
                // log("Unable to connect to WASM Websocket");
                return JsValue::from_str("WASM websocket error");
            }
        }
        ws.set_binary_type(web_sys::BinaryType::Blob);
        // let on_open_cb = Closure::<dyn FnMut()>::new(move || {
        //     log("socket opened");
        // });

        // ws.set_onopen(Some(on_open_cb.as_ref().unchecked_ref()));
        // on_open_cb.forget();

        // TODO: retrieve settings from websocket
        // ws.conn.send_with_str("__init__ps__").unwrap();
        let settings = Settings::default();

        let canvas = MagicSquare::canvas()
            .dyn_into::<web_sys::HtmlCanvasElement>()
            .unwrap();
        let canvas = Rc::new(canvas);

        let canvas_container = MagicSquare::canvas_container();
        let canvas_container = Rc::new(canvas_container);

        let geometry_cache: GeometryCache = GeometryCache::new(&settings.shapes);
        let geometry_cache = Rc::new(RefCell::new(geometry_cache));

        // wrap settings to be accessed by different closures
        let ui_buffer = UiBuffer { settings };
        let ui_buffer = Rc::new(RefCell::new(ui_buffer));

        let magic_square = MagicSquare::magic_square(); // awesome naming, great job!
        let magic_square = Rc::new(magic_square);

        // triggers cleanup requestAnimationFrame closure
        let destroy_flag: bool = false;
        let destroy_flag: Rc<RefCell<bool>> = Rc::new(RefCell::new(destroy_flag));

        let side_length: u32 = 0;
        let side_length = Rc::new(RefCell::new(side_length));
        // incriment idx_delay each render
        // when idx_delay reaches a desired delay value
        // incriment idx_offset
        let mut color_idx_offset_delay: [u8; 2] = [0, 0];

        let form = MagicSquare::form();
        let form = Rc::new(form);

        let mouse_pos_buffer: [f32; 2] = [0.0, 0.0];
        let mouse_pos_buffer: Rc<RefCell<[f32; 2]>> = Rc::new(RefCell::new(mouse_pos_buffer));

        // log("wasm var init done");
        {
            // init destroy listener on app_main
            // onDestroy hook in Main.svelte dispatches destroymswasm event
            // this closure flips destroy_flag
            // requestAnimationFrame checks value, cleans up resources
            let app_main = MagicSquare::app_main();
            let destroy_flag = destroy_flag.clone();
            let ws_c = ws.clone();
            
            // close wasm websocket
            let closure = Closure::<dyn FnMut(_)>::new(move |_event: web_sys::Event| {
                *destroy_flag.clone().borrow_mut() = true;
                ws_c.close_with_code(1001);
            });

            app_main
                .add_event_listener_with_callback("destroymswasm", closure.as_ref().unchecked_ref())
                .unwrap();

            closure.forget();
        }

        // set up cavnas container ResizeObserver
        {
            let canvas = canvas.clone();
            let canvas_container = canvas_container.clone();
            let side_length = side_length.clone();
            let handle_resize = Closure::<dyn FnMut(_)>::new(move |_: Vec<web_sys::ResizeObserverEntry>| {
                let canvas = canvas.clone();
                let side_length = side_length.clone();
                
                let container = canvas.parent_element().unwrap();
                let width: i32 = container.client_width();
                let height = container.client_height();
                let new_sl: u32 = ((i32::min(width, height) as f32) / 1.3) as u32;
                
                // update canvas
                canvas.set_height(new_sl);
                canvas.set_width(new_sl);
                *side_length.clone().borrow_mut() = new_sl;

                // update context
                let gl = MagicSquare::context(&canvas).unwrap();
                let new_sl_i32 = new_sl as i32;
                gl.viewport(0, 0, new_sl_i32, new_sl_i32);
            });

            if let Ok(resize_observer) = web_sys::ResizeObserver::new(handle_resize.as_ref().unchecked_ref()) {
                resize_observer.observe(&canvas_container.clone());
            }
            handle_resize.forget();
        }

        // set up WebSocket onMessage
        // take ownership of the passed-in js closure
        let set_all_settings = (*set_all_settings).clone();
        let set_all_settings: Rc<RefCell<js_sys::Function>> = Rc::new(RefCell::new(set_all_settings));

        {
            let ui_buffer = ui_buffer.clone();
            let ws_c = ws.clone();
            let set_all_settings = set_all_settings.clone();

            let onmessage_callback = Closure::<dyn FnMut(_)>::new(move |e: MessageEvent| {
                let ui_buffer = ui_buffer.clone();
                let set_all_settings = set_all_settings.clone();
                // here it receives and deserializes
                // log("wasm websocket onmessage_callback");
                unsafe { // it is the bytemuck-ing here tffhat is unsafe
                    let blob = e.data().dyn_into::<web_sys::Blob>().expect("Error Parsing Blob from Server");
                    
                    let fr = web_sys::FileReader::new().unwrap();
                    let fr_c = fr.clone();
                    // create onLoadEnd callback
                    let onloadend_cb = Closure::<dyn FnMut(_)>::new(move |_e: web_sys::ProgressEvent| {
                        let vec: Vec<u8> = js_sys::Uint8Array::new(&fr_c.result().unwrap()).to_vec();
                        let new_settings: &Settings = bytemuck::from_bytes(&vec[..]);
                        // set new_ettings in Wasm
                        ui_buffer.clone().borrow_mut().settings = *new_settings;
                        //
                        let this: JsValue = JsValue::null();
                        let settings_js = serde_wasm_bindgen::to_value(new_settings).expect("serde new_settings error");
                        let _ = set_all_settings.clone().borrow().call1(&this.clone(), &(settings_js.clone()));
                        // log(&format!("New Settings: {:?}", new_settings));
                    });
                    fr.set_onloadend(Some(onloadend_cb.as_ref().unchecked_ref()));
                    fr.read_as_array_buffer(&blob).expect("blob not readable");
                    onloadend_cb.forget();
                }
            });

            ws_c.set_onmessage(Some(onmessage_callback.as_ref().unchecked_ref()));
            onmessage_callback.forget();
            // log("WOW ZOW NOW!");
        }

        {
            // init UI control settings listener
            let form = form.clone();
            let ui_buffer = ui_buffer.clone();
            let geometry_cache = geometry_cache.clone();
            let ws_c = ws.clone();

            let closure_handle_input =
                Closure::<dyn FnMut(_)>::new(move |event: web_sys::Event| {
                    let input = event
                        .target()
                        .unwrap()
                        .dyn_into::<web_sys::HtmlInputElement>()
                        .unwrap();
                    let id = input.id();
                    let val = input.value();
                    ui_buffer.clone().borrow_mut().update(
                        id,
                        val,
                        &mut geometry_cache.clone().borrow_mut()
                    );
                    // bytemuck is the unsafe part here
                    // however, Settings is repr(c) 
                    // hence, we can rely on bytemuck for de/serialization through transmutation
                    unsafe {
                        let settings = ui_buffer.clone().borrow().settings.clone();
                        let settings_blob = bytemuck::bytes_of(&settings);
                        // log(&format!("settings blob: {:?}", settings_blob));
                        ws_c.send_with_u8_array(settings_blob);
                    }
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
            let side_length = side_length.clone();

            if !touch_screen {
                let closure = Closure::<dyn FnMut(_)>::new(move |event: web_sys::MouseEvent| {
                    let side_length = *side_length.clone().borrow() as i32;
                    mouse_pos_buffer.clone().borrow_mut()[0] =
                        MagicSquare::clip_x(event.offset_x(), side_length);
                    mouse_pos_buffer.clone().borrow_mut()[1] =
                        MagicSquare::clip_x(event.offset_y(), side_length);
                    magic_square
                        .dispatch_event(&web_sys::Event::new("render").unwrap())
                        .unwrap();
                });

                canvas
                    .add_event_listener_with_callback("mousemove", closure.as_ref().unchecked_ref())
                    .unwrap();

                closure.forget();
            } else {
                let inner_canvas = canvas.clone();
                let closure = Closure::<dyn FnMut(_)>::new(move |event: web_sys::TouchEvent| {
                    let side_length = *side_length.clone().borrow() as i32;
                    mouse_pos_buffer.clone().borrow_mut()[0] = MagicSquare::clip_x(
                        event.target_touches().item(0).unwrap().client_x()
                            - inner_canvas.clone().offset_left(),
                        side_length
                    );

                    mouse_pos_buffer.clone().borrow_mut()[1] = MagicSquare::clip_x(
                        event.target_touches().item(0).unwrap().client_y()
                            - inner_canvas.clone().offset_top(),
                        side_length
                    );
                });
                canvas
                    .clone()
                    .add_event_listener_with_callback("touchmove", closure.as_ref().unchecked_ref())
                    .unwrap();

                closure.forget();
            }
        }

        {
            // set up animation loop
            let canvas = canvas.clone();
            let ui_buffer = ui_buffer.clone();
            let mut animation = Animation::new();
            animation.set_from(&ui_buffer.clone().borrow().settings);

            let mut x: f32 = -X_MAX;

            let destroy_flag = destroy_flag.clone();

            // closures used to allocate and clean up resources
            let f: Rc<RefCell<Option<wasm_bindgen::prelude::Closure<_>>>> =
                Rc::new(RefCell::new(None));
            let g = f.clone();

            let mut frame_counter: usize = 0;

            // set up WebGL
            let gl: web_sys::WebGl2RenderingContext = MagicSquare::context(&canvas).unwrap();
            gl.clear_color(0.0, 0.0, 0.0, 0.0);
            gl.clear(WebGl2RenderingContext::COLOR_BUFFER_BIT);

            let program: WebGlProgram =
                GlProgram::new(&gl).expect(&format!("ISSUE INIT GL_PROGRAM"));
            gl.use_program(Some(&program));
            let uniform_locations = UniformLocations::new(&gl, &program);
            let mut uniforms = GlUniforms::new();

            let gl_buffer = gl.create_buffer().ok_or("Failed to create buffer").unwrap();
            gl.bind_buffer(WebGl2RenderingContext::ARRAY_BUFFER, Some(&gl_buffer));

            // set gl to read vertex data from geometry_cache.vertices

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

            let vao = gl
                .create_vertex_array()
                .ok_or("Could not create vertex array object")
                .unwrap();
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

            // animation callback
            *g.borrow_mut() = Some(Closure::new(move || {
                let mut settings = ui_buffer.clone().borrow().settings;
                animation.set_from(&settings);

                if *destroy_flag.clone().borrow() {
                    // cleanup resource
                    let _ = f.borrow_mut().take();
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

                    x += 0.001;
                    if x == X_MAX {
                        x = -X_MAX;
                    }

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

                    // 0 < color_speed < 21
                    if color_idx_delay == delay_reset {
                        color_idx_offset_delay[0] = match settings.color_direction {
                            ColorDirection::In => (color_idx_offset - 1) % CACHE_CAPACITY as u8,
                            ColorDirection::Fix => color_idx_offset,
                            ColorDirection::Out => (color_idx_offset + 1) % CACHE_CAPACITY as u8,
                        };
                        color_idx_offset_delay[1] = 0;
                    }

                    color_idx_offset_delay[1] += 1;

                    // compute
                    uniforms.set_uniforms(&mouse_pos_buffer, &settings, color_idx_offset);
                    // log(&format!("{:?}", uniforms));

                    // draw
                    // log(&format!("{:?}", geometry_cache.clone().borrow().vertices));
                    if GlDraw::scene(
                        &gl,
                        &uniforms,
                        &uniform_locations,
                        &animation.curr_shapes(),
                        &settings.transform_order,
                        &x,
                    )
                    .is_err()
                    {
                        // log("DRAW ERROR");
                    }

                    let frame_counter_limit: i32 = if settings.draw_pattern_speed > 19 {
                        1
                    } else {
                        21 - settings.draw_pattern_speed
                    };

                    if frame_counter > frame_counter_limit as usize {
                        frame_counter = 0;
                    }

                    if frame_counter % frame_counter_limit as usize == 0 {
                        animation.inc();
                    }
                    frame_counter = (frame_counter + 1) % (frame_counter_limit as usize);
                    MagicSquare::request_animation_frame(f.borrow().as_ref().unwrap());
                }
            }));

            MagicSquare::request_animation_frame(g.borrow().as_ref().unwrap());
        }
    
        let to_js = ui_buffer.clone().borrow().settings;
        serde_wasm_bindgen::to_value(&to_js).unwrap()
    }
}

pub fn setup_gl(canvas: &web_sys::HtmlCanvasElement) -> web_sys::WebGl2RenderingContext {
    // set up WebGL
    let gl: web_sys::WebGl2RenderingContext = MagicSquare::context(&canvas).unwrap();
    gl.clear_color(0.0, 0.0, 0.0, 0.0);
    gl.clear(WebGl2RenderingContext::COLOR_BUFFER_BIT);

    let program: WebGlProgram = GlProgram::new(&gl).expect(&format!("ISSUE INIT GL_PROGRAM"));
    gl.use_program(Some(&program));

    let gl_buffer = gl.create_buffer().ok_or("Failed to create buffer").unwrap();
    gl.bind_buffer(WebGl2RenderingContext::ARRAY_BUFFER, Some(&gl_buffer));

    // set gl to read vertex data from geometry_cache.vertices

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

    let vao = gl
        .create_vertex_array()
        .ok_or("Could not create vertex array object")
        .unwrap();
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

    gl
}
