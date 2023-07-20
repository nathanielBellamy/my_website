use js_sys::{DataView, ArrayBuffer};
use wasm_bindgen::JsValue;
use crate::log;
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
use web_sys::{MessageEvent, WebGl2RenderingContext, WebGlProgram};

use crate::{magic_square::{settings::Settings, main::MagicSquare, main::X_MAX}, websocket::{Websocket, WebsocketConnError}};

use super::deser::Deser;
use super::ui_buffer::JsValueBit;

// TODO:
//  run an instance of magic square that has access to the websocket
//  UiBuffer sends serialized JSON (see deser.rs) to backend through socket on update
//  UiBuffer is updated with new received Settings when message recevied

const ARRAY_BUFFER_CAPACITY: u32 = 5000;
const URL: &str = "ws://localhost:8080/public-square-wasm-ws";

#[derive(Debug)]
pub struct PublicSquare {
    ab: ArrayBuffer,
    pub client_id: u64,
    pub dv: DataView,
    pub websocket: Websocket
}

impl PublicSquare {
    pub fn new(client_id: u64) -> Result<PublicSquare, WebsocketConnError> {
        let ab = ArrayBuffer::new(ARRAY_BUFFER_CAPACITY);
        let websocket = Websocket::new(URL.to_owned())?;
        Ok(PublicSquare {
            dv: DataView::new(&ab, 0, ARRAY_BUFFER_CAPACITY as usize),
            ab,
            client_id,
            websocket,
        })
    }

    pub fn send(&mut self, settings: Settings) -> Result<(), JsValue> {
        let u8_slice: &[u8] = unsafe { Deser::any_as_u8_slice(&settings) };
        for (idx, val) in u8_slice.iter().enumerate() {
            self.dv.set_uint8(idx, *val)
        }
        self.websocket.conn.send_with_array_buffer(&self.ab)
    }

    pub fn close(&self) -> Result<(), JsValue> {
        self.websocket.conn.close()
    }
}

#[wasm_bindgen]
struct PubSq;

#[wasm_bindgen]
impl PubSq {
    #[allow(unused)]
    // called from js
    pub async fn run(touch_screen: JsValue) -> JsValue {
        let touch_screen: bool = serde_wasm_bindgen::from_value(touch_screen).unwrap();

        // TODO: retrieve settings from websocket

        let settings = Settings::default();
        let pub_sq: PublicSquare;
        match PublicSquare::new(1){
            Ok(ps) => pub_sq = ps,
            Err(e) => {
                return serde_wasm_bindgen::to_value(&e).unwrap();
            }
        }
        log(&format!("{:?}", pub_sq));
        let pub_sq = Rc::new(RefCell::new(pub_sq));
        
        log("WOW ZOW 1");
        let canvas = MagicSquare::canvas()
            .dyn_into::<web_sys::HtmlCanvasElement>()
            .unwrap();
        let canvas = Rc::new(canvas);

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

        let height: i32 = canvas.client_height();
        let width: i32 = canvas.client_width();
        // incriment idx_delay each render
        // when idx_delay reaches a desired delay value
        // incriment idx_offset
        let mut color_idx_offset_delay: [u8; 2] = [0, 0];

        let form = MagicSquare::form();
        let form = Rc::new(form);

        let mouse_pos_buffer: [f32; 2] = [0.0, 0.0];
        let mouse_pos_buffer: Rc<RefCell<[f32; 2]>> = Rc::new(RefCell::new(mouse_pos_buffer));

        log("wasm var init done");
        {
            // init destroy listener on app_main
            // onDestroy hook in Main.svelte dispatches destroymswasm event
            // this closure flips destroy_flag
            // requestAnimationFrame checks value, cleans up resources
            let app_main = MagicSquare::app_main();
            let destroy_flag = destroy_flag.clone();
            
            // close wasm websocket
            let pub_sq = pub_sq.clone();
            let closure = Closure::<dyn FnMut(_)>::new(move |_event: web_sys::Event| {
                *destroy_flag.clone().borrow_mut() = true;
                pub_sq.clone().borrow().close();
            });

            app_main
                .add_event_listener_with_callback("destroymswasm", closure.as_ref().unchecked_ref())
                .unwrap();

            closure.forget();
        }

        {
            let ui_buffer = ui_buffer.clone();
            let pub_sq = pub_sq.clone();

            let onmessage_callback = Closure::<dyn FnMut(_)>::new(move |e: MessageEvent| {
                // pub_sq sends bin seriaiized settings
                // here it receives and deserializes
                log(&format!("wasm websocket onmessage_callback: {:?}", e));
                let raw_bin = JsValueBit(&e.data() as *const JsValue);
                if let Ok(res) = bytemuck::try_cast::<JsValueBit, Settings>(raw_bin) {
                    log(&format!("{:?}", res));
                    ui_buffer.clone().borrow_mut().settings = res;
                }
            });

            pub_sq.clone().borrow_mut()
                .websocket
                .conn
                .set_onmessage(Some(onmessage_callback.as_ref().unchecked_ref()));
            onmessage_callback.forget();
            log("WOW ZOW NOW!");
        }

        {
            // init UI control settings listener
            let form = form.clone();
            let ui_buffer = ui_buffer.clone();
            let geometry_cache = geometry_cache.clone();
            let pub_sq = pub_sq.clone();

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
                    ui_buffer.clone().borrow_mut().update(
                        id,
                        val,
                        &mut geometry_cache.clone().borrow_mut(),
                        pub_sq.clone(),
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
            if !touch_screen {
                let closure = Closure::<dyn FnMut(_)>::new(move |event: web_sys::MouseEvent| {
                    mouse_pos_buffer.clone().borrow_mut()[0] =
                        MagicSquare::clip_x(event.offset_x(), width);
                    mouse_pos_buffer.clone().borrow_mut()[1] =
                        MagicSquare::clip_x(event.offset_y(), height);
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
                    mouse_pos_buffer.clone().borrow_mut()[0] = MagicSquare::clip_x(
                        event.target_touches().item(0).unwrap().client_x()
                            - inner_canvas.clone().offset_left(),
                        width,
                    );

                    mouse_pos_buffer.clone().borrow_mut()[1] = MagicSquare::clip_x(
                        event.target_touches().item(0).unwrap().client_y()
                            - inner_canvas.clone().offset_top(),
                        height,
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
                        log("DRAW ERROR");
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
