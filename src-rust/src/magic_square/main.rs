use wasm_bindgen::prelude::*;
use std::rc::Rc;

use web_sys::{WebGl2RenderingContext, WebGlProgram, WebGlShader};

const VERTEX_COUNT: usize = 4;
const COORDINATE_COUNT: usize = VERTEX_COUNT * 3;
const VERTICES_EMPTY: [f32; VERTEX_COUNT * 3] = [0.0; VERTEX_COUNT * 3];

// => keep buffer in RC
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

        {   
            // init mouse move listener
            // write coordinates to buffer
            let mut buffer: [i32; 2] = [0,0];// Buffer::new();

            let canvas = canvas.clone();
            let context: web_sys::WebGl2RenderingContext = MagicSquare::context(&canvas).unwrap();
            let closure = Closure::<dyn FnMut(_)>::new(move |event: web_sys::MouseEvent| {
                context.clear_color(0.0, 0.0, 0.0, 0.0);
                context.clear(WebGl2RenderingContext::COLOR_BUFFER_BIT);
                buffer[0] = event.offset_x();
                buffer[1] = event.offset_y();
                MagicSquare::render_all_lines(&buffer, &context)
            });

            canvas.add_event_listener_with_callback("mousemove", closure.as_ref().unchecked_ref())?;
            closure.forget();
        }

        Ok(())
    }
}


pub type Rgba = [f64; 4];

impl MagicSquare {

    fn render_all_lines(buffer: &[i32; 2], context: &web_sys::WebGl2RenderingContext) {
        for idx in 1..10 {
            let vertices = MagicSquare::get_vertices(buffer, idx);
            let rgba = MagicSquare::get_rgba(buffer, idx);
            MagicSquare::render(&vertices, &rgba, context).expect("Render error");
        }
    }

    fn get_vertices(buffer: &[i32; 2], idx: usize) -> [f32; 6] {
        let mut vertices = [0.0, 0.0, 0.0, 0.0, 0.0, 0.0];

        vertices[0] = buffer[0] as f32 * 0.001 + (5.0 * idx as f32 / 50.0);
        vertices[1] = buffer[1] as f32 * 0.001 + (5.0 * idx as f32 / 50.0);
        vertices[3] = -(buffer[1] as f32 * 0.001) - (5.0 * idx as f32 / 50.0);
        vertices[4] = (buffer[0] as f32 * 0.001) - (5.0 * idx as f32 / 50.0);

        vertices
    }

    fn get_rgba(buffer: &[i32; 2], idx: usize) -> Rgba {
        let mut result: Rgba = [0.0, 0.0, 0.0, 0.0];
        
        result[0] = ((buffer[0] as f64) * 0.1).sin() / 3.0 + (0.1 * idx as f64);
        result[1] = ((buffer[1] as f64) * 0.1).sin() / 3.0 + (0.1 * idx as f64);
        result[2] = ((-buffer[0] as f64) * 0.1).sin() / 3.0 + (0.1 * idx as f64);
        result[3] = ((buffer[1] as f64) * 0.1).sin() / 3.0 + (0.1 * idx as f64);

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

    pub fn context(canvas: &web_sys::HtmlCanvasElement) -> Result<web_sys::WebGl2RenderingContext, JsValue> {
        let context = canvas
                        .get_context("webgl2")?
                        .expect("unable to get webgl2 context")
                        .dyn_into::<WebGl2RenderingContext>()?;

        Ok(context)
    }

    fn draw(context: &WebGl2RenderingContext, vert_count: i32) {
       ;

        context.draw_arrays(WebGl2RenderingContext::LINES, 0, vert_count);
    }

    fn render(vertices: &[f32; 6], color: &Rgba, context: &web_sys::WebGl2RenderingContext) ->  Result<(), JsValue>  {
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


pub struct ShaderCompiler;

impl ShaderCompiler {
    pub fn exec(
        context: &WebGl2RenderingContext,
        shader_type: u32,
        source: &str,
    ) -> Result<WebGlShader, String> {
        let shader = context
            .create_shader(shader_type)
            .ok_or_else(|| String::from("Unable to create shader object"))?;
        context.shader_source(&shader, source);
        context.compile_shader(&shader);

        if context
            .get_shader_parameter(&shader, WebGl2RenderingContext::COMPILE_STATUS)
            .as_bool()
            .unwrap_or(false)
        {
            Ok(shader)
        } else {
            Err(context
                .get_shader_info_log(&shader)
                .unwrap_or_else(|| String::from("Unknown error creating shader")))
        }
    }
    
    pub fn frag_default(
        context: &WebGl2RenderingContext,
        rgba: &Rgba
    ) -> Result<WebGlShader, String> {
        let string = format!(
                r##"#version 300 es
                precision highp float;
                out vec4 outColor;
                
                void main() {{
                    outColor = vec4({}, {}, {}, {});
                }}
                "##,rgba[0], rgba[1], rgba[2], rgba[3]);


        ShaderCompiler::exec(
            &context,
            WebGl2RenderingContext::FRAGMENT_SHADER,
            &string
        )
    }

    pub fn vert_default(
        context: &WebGl2RenderingContext,
    ) -> Result<WebGlShader, String> {
         ShaderCompiler::exec(
            &context,
            WebGl2RenderingContext::VERTEX_SHADER,
            r##"#version 300 es
     
            in vec4 position;

            void main() {
            
                gl_Position = position;
            }
            "##,
        )
    }
}

pub struct ProgramLinker;

impl ProgramLinker {
    pub fn exec(
        context: &WebGl2RenderingContext,
        vert_shader: &WebGlShader,
        frag_shader: &WebGlShader,
    ) -> Result<WebGlProgram, String> {
        let program = context
            .create_program()
            .ok_or_else(|| String::from("Unable to create shader object"))?;

        context.attach_shader(&program, vert_shader);
        context.attach_shader(&program, frag_shader);
        context.link_program(&program);

        if context
            .get_program_parameter(&program, WebGl2RenderingContext::LINK_STATUS)
            .as_bool()
            .unwrap_or(false)
        {
            Ok(program)
        } else {
            Err(context
                .get_program_info_log(&program)
                .unwrap_or_else(|| String::from("Unknown error creating program object")))
        }
    }
}




