use std::rc::Rc;
use std::sync::{Arc, Mutex};
use wasm_bindgen::prelude::*;
use web_sys::WebGl2RenderingContext;
// use crate::magic_square::vertices::Vertices;
use crate::magic_square::shader_compiler::ShaderCompiler;
use crate::magic_square::program_linker::ProgramLinker;
use crate::magic_square::transformations::{Rotation, RotationSequence, Translation};
use crate::magic_square::geometry::{Geometry, Shape};
use crate::magic_square::geometry::cache::{Cache as GeometryCache, CACHE_CAPACITY};
// use crate::magic_square::traits::VertexStore;
// use super::geometry::icosohedron::Icosohedron;
use crate::magic_square::worker::Worker;

#[wasm_bindgen]
extern "C" {
    #[wasm_bindgen(js_namespace = console)]
    fn log(s: &str);

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
    pub fn run() -> Result<(), JsValue> {

        log("fooooo");

        // testing multithreading
        //
        // let (to_worker, from_main) = std::sync::mpsc::channel();
        // let (to_main, from_worker) = std::sync::mpsc::channel();
        // Worker::spawn(move || { to_main.send(from_main.recv().unwrap() + 1.0); });
        // to_worker.send(1.0);
        // assert_eq!(from_worker.recv().unwrap(), 2.0);

        let canvas = MagicSquare::canvas().dyn_into::<web_sys::HtmlCanvasElement>()?;
        let canvas = Rc::new(canvas);

        // TODO: handle resize
        // add height and width fields to MagicSquare
        // run returns MagicSquare instance to js
        // js uses svelte watch resize to update height and width
        // pass immutable reference of h&w to closure
        let height:i32 = canvas.client_height();
        let width:i32 = canvas.client_width();
        let geometry_cache = Arc::new(Mutex::new(
            GeometryCache::new(
                26, 
                [[0.0; 1_200]; CACHE_CAPACITY], 
                [[0.0, 0.0, 0.0, 0.0]; CACHE_CAPACITY],
                [Shape::None; CACHE_CAPACITY]
            )
        ));
        
        log("here!");

        let context: web_sys::WebGl2RenderingContext = MagicSquare::context(&canvas).unwrap();
        context.clear_color(1.0, 1.0, 0.0, 0.0);
        context.clear(WebGl2RenderingContext::COLOR_BUFFER_BIT);
        MagicSquare::render_all_lines([0.0, 0.0], geometry_cache.clone(), &context);
        
        log("there!");
        {
            // init mouse move listener
            // write coordinates to buffer
            let mut buffer: [f32; 2] = [0.0, 0.0]; // Buffer::new();
            let geometry_cache = geometry_cache.clone();

            let canvas = canvas.clone();
            let context: web_sys::WebGl2RenderingContext = MagicSquare::context(&canvas).unwrap();
            let closure = Closure::<dyn FnMut(_)>::new(move |event: web_sys::MouseEvent| {
                context.clear_color(0.0, 0.0, 0.0, 0.0);
                context.clear(WebGl2RenderingContext::COLOR_BUFFER_BIT);
                buffer[0] = MagicSquare::clip_x(event.offset_x(), width);
                buffer[1] = MagicSquare::clip_y(event.offset_y(), height);
                MagicSquare::render_all_lines(buffer, geometry_cache.clone(), &context);
            });

            canvas
                .add_event_listener_with_callback("mousemove", closure.as_ref().unchecked_ref())?;
            closure.forget();
        }

        Ok(())
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

    fn render_all_lines(
        buffer: [f32; 2], 
        geometry_cache: Arc<Mutex<GeometryCache>>, 
        context: &web_sys::WebGl2RenderingContext
    ) {
        for idx in 0..2 { // geometry_cache.max_idx + 1 { //TODO: settings.cache_per
            let geometry_cache = geometry_cache.clone();
            // let  idx = geometry_cache.idx + 1;// avoid 0 idx
            let rot_seq = RotationSequence::new(
                Rotation::new(Axis::X, buffer[0] * 3.14), // + (idx as f32) * 0.05),
                Rotation::new(Axis::Y, buffer[1] * 3.14), // + (idx as f32) * 0.05),
                Rotation::new(Axis::Z, 0.0),
            );

            let translation = Translation { x: 0.0, y: 0.0, z: 0.0 }; // { x: buffer[0], y: buffer[1], z: 0.0 };

            // let hexagon = Geometry::hexagon(
            //     0.025 * idx as f32, 
            //     rot_seq,
            //     translation
            // );

            let icosohedron = Geometry::icosohedron(
                0.25 * idx as f32 + 0.01, 
                rot_seq,
                translation
            );

            let rgba = MagicSquare::get_rgba(buffer, idx);
            
            let _ = Worker::spawn(move || {
                geometry_cache.lock().unwrap().set_next(icosohedron.arr, rgba, Shape::Icosohedron);
            });
        }

        let geometry_cache = geometry_cache.lock().unwrap();
        for idx in 0..2 {//geometry_cache.max_idx { // settings.
            MagicSquare::render(
                geometry_cache.gl_vertices(idx), 
                &geometry_cache.rgbas[idx], 
                &context
            ).expect("Render error");
        }
    }

    fn get_rgba(buffer: [f32; 2], idx: usize) -> Rgba {
        let mut result: Rgba = [0.0, 0.0, 0.0, 0.0];

        result[0] = 1.0 - buffer[0];
        result[1] = 1.0 - buffer[1];
        result[2] = 1.0 - (idx as f32 / CACHE_CAPACITY as f32);
        result[3] = 0.1 * idx as f32;
        result
    }

    fn window() -> web_sys::Window {
        web_sys::window().expect("no global `window` exists")
    }

    fn document() -> web_sys::Document {
        MagicSquare::window()
            .document()
            .expect("should have a document on window")
    }

    fn canvas() -> web_sys::Element {
        MagicSquare::document()
            .get_element_by_id("magic_square_canvas")
            .expect("unable to find canvas element")
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


