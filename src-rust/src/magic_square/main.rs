use ndarray::prelude::*;
use ndarray::Array;
use ndarray::Dim;
use std::rc::Rc;
use wasm_bindgen::prelude::*;
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

        // TODO: handle resize
        let height:i32 = canvas.client_height();
        let width:i32 = canvas.client_width();

        {
            // init mouse move listener
            // write coordinates to buffer
            let mut buffer: [i32; 2] = [0, 0]; // Buffer::new();

            let canvas = canvas.clone();
            let context: web_sys::WebGl2RenderingContext = MagicSquare::context(&canvas).unwrap();
            let closure = Closure::<dyn FnMut(_)>::new(move |event: web_sys::MouseEvent| {
                context.clear_color(0.0, 0.0, 0.0, 0.0);
                context.clear(WebGl2RenderingContext::COLOR_BUFFER_BIT);
                buffer[0] = event.offset_x();
                buffer[1] = event.offset_y();
                MagicSquare::render_all_lines(&buffer, &context, height, width);
            });

            canvas
                .add_event_listener_with_callback("mousemove", closure.as_ref().unchecked_ref())?;
            closure.forget();
        }

        Ok(())
    }
}

pub type Rgba = [f64; 4];
pub type Vertex = [f32; 4];
pub type VertexArr = [f32; 400];

pub struct Vertices {
    arr: VertexArr,
    idx: usize
}

impl Vertices {
    pub fn new() -> Vertices {
        Vertices { 
            arr: [0.0; 400], 
            idx: 0 
        }
    }

    pub fn set_next(&mut self, vertex: Vertex) {
        if self.idx > self.arr.len() - 1 { return; }
        for i in 0..3 {
            self.arr[self.idx + i] = vertex[i]
        }
        self.idx += 4;
    }
}

impl MagicSquare {
    fn render_all_lines(buffer: &[i32; 2], context: &web_sys::WebGl2RenderingContext, height: i32, width: i32) {
        let mut all_vertices = Vertices::new();

        // for idx in 1..10 {
            let v = MagicSquare::get_vertices(buffer, 1, 'x', height, width);
            all_vertices.set_next([v[0], v[1], v[2], v[3]]);
            all_vertices.set_next([v[4], v[5], v[6], v[7]]);
        // }
    
        // for idx in 1..10 {
        //     let v = MagicSquare::get_vertices(buffer, idx, 'y', height, width);
        //     all_vertices.set_next([v[0], v[1], v[2], v[3]]);
        //     all_vertices.set_next([v[4], v[5], v[6], v[7]]);
        // }

        // for idx in 1..10 {
        //     let v = MagicSquare::get_vertices(buffer, idx, 'z', height, width);
        //     all_vertices.set_next([v[0], v[1], v[2], v[3]]);
        //     all_vertices.set_next([v[4], v[5], v[6], v[7]]);
        // }


        let rgba = MagicSquare::get_rgba(buffer, 1);
        MagicSquare::render(&all_vertices, &rgba, context).expect("Render error");
    }

   fn rotx_matrix(theta: f32) -> Array<f32, Ix2> {
        array![
            [1.0, 0.0, 0.0, 0.0],
            [0.0, theta.cos(), theta.sin(), 0.0],
            [0.0, -theta.sin(), theta.cos(), 0.0],
            [0.0, 0.0, 0.0, 1.0],
        ]
    }

    fn roty_matrix(theta: f32) -> Array<f32, Ix2> {
        array![
            [theta.cos(), 0.0, -theta.sin(), 0.0],
            [0.0, 1.0, 0.0, 0.0],
            [theta.sin(), 0.0, theta.cos(), 0.0],
            [0.0, 0.0, 0.0, 0.0]
        ]
    }

    fn rotz_matrix(theta: f32) -> Array<f32, Ix2> {
        array![
            [theta.cos(), theta.sin(), 0.0, 0.0],
            [-theta.sin(), theta.cos(), 0.0, 0.0],
            [0.0, 0.0, 1.0, 0.0],
            [0.0, 0.0, 0.0, 1.0]
        ]
    }

    fn get_vertices(buffer: &[i32; 2], idx: usize, axis: char, height: i32, width: i32) -> [f32; 8] {
        let mut result: [f32; 8] = [0.0; 8];
        
        let clip_x: f32 = (2.0 * (buffer[0] as f32) / width as f32) - 1.0;
        let clip_y: f32 = 1.0 - ((2.0 * buffer[1] as f32) / height as f32);

        let line_base: Array<f32, _> = array![
            [clip_x, clip_y],
            [0.0, 0.0],
            [0.0, 0.0],
            [0.0, 0.0],
        ];
        
        let theta: f32 = buffer[0] as f32 / (100.0 * idx as f32);
        let rot_matrix = match axis {
            'y' => MagicSquare::roty_matrix(theta),
            'z' => MagicSquare::rotz_matrix(theta),
            _ => MagicSquare::rotx_matrix(theta),
        };
        let rotated_line: Array<f32, _> = rot_matrix.dot(&line_base);
        // let rotated_line = line_base;

        // flatten
        let mut counter: usize = 0;
        for coord in rotated_line.iter() {
            result[counter] = *coord;
            counter += 1;
        }

        result
    }

    fn get_rgba(buffer: &[i32; 2], idx: usize) -> Rgba {
        let mut result: Rgba = [0.0, 0.0, 0.0, 0.0];

        result[0] = ((buffer[0] as f64) * 0.1).sin() / 15.0 + (0.3 * idx as f64);
        result[1] = ((buffer[1] as f64) * 0.1).sin() / 15.0 + (0.3 * idx as f64);
        result[2] = ((-buffer[0] as f64) * 0.1).sin() / 15.0 + (0.3 * idx as f64);
        result[3] = 1.0;
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
        rgba: &Rgba,
    ) -> Result<WebGlShader, String> {
        let string = format!(
            r##"#version 300 es
                precision highp float;
                out vec4 outColor;
                
                void main() {{
                    outColor = vec4({}, {}, {}, {});
                }}
                "##,
            rgba[0], rgba[1], rgba[2], rgba[3]
        );

        ShaderCompiler::exec(&context, WebGl2RenderingContext::FRAGMENT_SHADER, &string)
    }

    pub fn vert_default(context: &WebGl2RenderingContext) -> Result<WebGlShader, String> {
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



