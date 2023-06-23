use wasm_bindgen::prelude::*;
use web_sys::{WebGl2RenderingContext, WebGlProgram};

// use crate::magic_square::main::log;
use super::animation::Shapes;
use super::geometry::Geometry;
use super::geometry::cache::CACHE_CAPACITY;
use super::gl_uniforms::GlUniforms;
use super::main::log;
// use super::main::log;
use super::settings::TransformOrder;


pub struct GlDraw;

impl GlDraw {
    pub fn scene(
        gl: &WebGl2RenderingContext,
        program: &WebGlProgram,
        uniforms: &GlUniforms,
        shapes: &Shapes,
        _order: TransformOrder,
    ) -> Result<(), JsValue>{
        // lookup uniforms
        // let translation_location = gl.get_uniform_location(program, "u_translation").unwrap();
        // // let order_location: WebGlUniformLocation = gl.get_uniform_location(program, "u_order").unwrap();
        let rgba_location = gl.get_uniform_location(program, "u_rgba").unwrap();
        // let radius_location = gl.get_uniform_location(program, "u_radius").unwrap();
        let rotation_0_location = gl.get_uniform_location(program, "u_rotation_0").unwrap();
        // let rotation_1_location = gl.get_uniform_location(program, "u_rotation_1").unwrap();
        // let rotation_2_location = gl.get_uniform_location(program, "u_rotation_2").unwrap();

        // set uniforms
        for idx in 0..4 { // CACHE_CAPACITY {
            gl.uniform4f(
                Some(&rgba_location), 
                uniforms.rgbas[idx][0],
                uniforms.rgbas[idx][1],
                uniforms.rgbas[idx][2],
                uniforms.rgbas[idx][3],
            );
            // // gl.uniform1f(
            // //     Some(&order_location),
            // //     match order {
            // //         TransformOrder::RotateThenTranslate => 1.0,
            // //         TransformOrder::TranslateThenRotate => 0.0,
            // //     }
            // // );
            // gl.uniform4fv_with_f32_array(Some(&radius_location), &uniforms.radii[idx]);
            gl.uniform_matrix4fv_with_f32_array(
                Some(&rotation_0_location),
                false,
                &[
                    0.0, 1.0, 0.0, 0.0,
                    -1.0, 0.0, 0.0, 0.0,
                    0.0, 0.0, 1.0, 0.0,
                    0.0, 0.0, 0.0, 1.0,
                ]
                //&uniforms.rotations[idx][0]
            );
            // gl.uniform4fv_with_f32_array(Some(&rotation_1_location), &uniforms.rotations[idx][1]);
            // gl.uniform4fv_with_f32_array(Some(&rotation_2_location), &uniforms.rotations[idx][2]);
            // gl.uniform4fv_with_f32_array(Some(&translation_location), &uniforms.translations[idx]);

            // log("hey wow wow wow");
            // log(&format!("TRANS {idx}: {:?}", js_sys::JSON::stringify(&gl.get_uniform(program, &translation_location))));
            // log(&format!("TRANS uniforms {idx}: {:?}", uniforms.translations[idx]));
            // log(&format!("RGBA {idx}: {:?}", js_sys::JSON::stringify(&gl.get_uniform(program, &rgba_location))));
            // log(&format!("RGBA uniforms {idx}: {:?}", uniforms.rgbas[idx]));

            log(&format!("ROT 0 {idx}: {:?}", js_sys::JSON::stringify(&gl.get_uniform(program, &rotation_0_location))));

            // Draw the geometry.
            let _offset = idx * 300;
            let _count = Geometry::to_vertex_count(shapes[idx]);
            gl.use_program(Some(program));
            let idx_left: usize = if idx == 0 { 0 } else if idx == 1 { 2 } else if idx == 2 { 4 } else { 6 };
            gl.draw_arrays(WebGl2RenderingContext::LINES, idx_left as i32, 2);//offset as i32, count as i32);
        }
        Ok(())
    }
}


