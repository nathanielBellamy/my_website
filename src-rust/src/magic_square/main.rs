use std::rc::Rc;
use wasm_bindgen::prelude::*;
use web_sys::WebGl2RenderingContext;
use crate::magic_square::vertices::Vertices;
use crate::magic_square::shader_compiler::ShaderCompiler;
use crate::magic_square::program_linker::ProgramLinker;
use crate::magic_square::transformations::{Rotation, RotationSequence};
use crate::magic_square::geometry::Geometry;

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
        let canvas = MagicSquare::canvas().dyn_into::<web_sys::HtmlCanvasElement>()?;
        let canvas = Rc::new(canvas);

        // TODO: handle resize
        // add height and width fields to MagicSquare
        // run returns MagicSquare instance to js
        // js uses svelte watch resize to update height and width
        // pass immutable reference of h&w to closure
        let height:i32 = canvas.client_height();
        let width:i32 = canvas.client_width();

        {
            // init mouse move listener
            // write coordinates to buffer
            let mut buffer: [f32; 2] = [0.0, 0.0]; // Buffer::new();

            let canvas = canvas.clone();
            let context: web_sys::WebGl2RenderingContext = MagicSquare::context(&canvas).unwrap();
            let closure = Closure::<dyn FnMut(_)>::new(move |event: web_sys::MouseEvent| {
                context.clear_color(0.0, 0.0, 0.0, 0.0);
                context.clear(WebGl2RenderingContext::COLOR_BUFFER_BIT);
                buffer[0] = MagicSquare::clip_x(event.offset_x(), width);
                buffer[1] = MagicSquare::clip_y(event.offset_y(), height);
                MagicSquare::render_all_lines(&buffer, &context);
            });

            canvas
                .add_event_listener_with_callback("mousemove", closure.as_ref().unchecked_ref())?;
            closure.forget();
        }

        Ok(())
    }
}

pub type Rgba = [f64; 4];

impl MagicSquare {
    pub fn clip_x(offset_x: i32, width: i32) -> f32 {
        // x coordinate of mouse position in clip space
        (2.0 * (offset_x as f32) / width as f32) - 1.0
    }

    pub fn clip_y(offset_y: i32, height: i32) -> f32 {
        // y coordinate of mouse position in clip space
        1.0 - ((2.0 * offset_y as f32) / height as f32)
    }

    fn render_all_lines(buffer: &[f32; 2], context: &web_sys::WebGl2RenderingContext) {
        // TODO: multithread
        // stuff vertices
        let mut vertices: Vertices = Vertices::new();

        for idx in 1..20 {
            let rot_seq = RotationSequence::new(
                Rotation::new(Axis::X, buffer[0] as f32 * 3.14 + idx as f32 * 0.05),
                Rotation::new(Axis::Y, buffer[1] as f32 * 3.14 + idx as f32 * 0.05),
                Rotation::new(Axis::Z, 0.0),// (buffer[0] as f32 * buffer[1] as f32) * (3.14 / 2.0) + idx as f32 * 0.05) 
            );

            Geometry::hexagon(
                buffer, 0.025 * idx as f32, 
                rot_seq,
                &mut vertices
            );

            let rgba = MagicSquare::get_rgba(buffer, idx);
            MagicSquare::render(&vertices, &rgba, context).expect("Render error");
        }
    }

    fn get_rgba(buffer: &[f32; 2], idx: usize) -> Rgba {
        let mut result: Rgba = [0.0, 0.0, 0.0, 0.0];

        result[0] = 1.0 - buffer[0] as f64;
        result[1] = 1.0 - buffer[1] as f64;
        result[2] = 1.0 - (buffer[0] * buffer[1]) as f64;
        result[3] = 0.1 * idx as f64;
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
        vertices: &Vertices,
        color: &Rgba,
        context: &web_sys::WebGl2RenderingContext,
    ) -> Result<(), JsValue> {
        let vert_shader = ShaderCompiler::vert_default(&context).unwrap();
        let frag_shader = ShaderCompiler::frag_default(&context, &color).unwrap();

        let program = ProgramLinker::exec(&context, &vert_shader, &frag_shader)?;
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
            let positions_array_buf_view = js_sys::Float32Array::view(&vertices.arr);

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

        let vert_count = (vertices.arr.len() / 4) as i32;
        MagicSquare::draw(&context, vert_count);

        Ok(())
    }
}


