use std::rc::Rc;
use std::cell::RefCell;
use wasm_bindgen::prelude::*;
// use crate::magic_square::main::log;

use web_sys::WebGl2RenderingContext;
use super::geometry::cache::{Cache, CACHE_CAPACITY};
use super::shader_compiler::ShaderCompiler;
use super::program_linker::ProgramLinker;


pub struct Draw;

impl Draw {
    fn draw(context: &WebGl2RenderingContext, vert_count: i32) {
        context.draw_arrays(WebGl2RenderingContext::LINES, 0, vert_count);
    }

    pub fn scene(
        geometry_cache: Rc<RefCell<Cache>>,
        frag_shader_cache: Rc<RefCell<Vec<String>>>,
        color_offset: usize,
        context: &web_sys::WebGl2RenderingContext
    ) -> Result<(), JsValue>{
        let geometry_cache = geometry_cache.borrow();
        let frag_shader_cache = frag_shader_cache.borrow();
        for idx in 0..CACHE_CAPACITY {
            Draw::shape(
                geometry_cache.gl_vertices(idx),
                &frag_shader_cache[(color_offset + idx) % CACHE_CAPACITY].clone(),
                context
            )?;
        }
        Ok(())
    }

    fn shape(
        vertices: &[f32],
        frag_shader_str: &str,
        context: &web_sys::WebGl2RenderingContext,
    ) -> Result<(), JsValue> {
        // log(frag_shader_str);
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
        Draw::draw(&context, vert_count);

        Ok(())
    }
}
