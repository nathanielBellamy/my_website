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
        let translation_location = gl.get_uniform_location(program, "u_translation").unwrap();
        // // let order_location: WebGlUniformLocation = gl.get_uniform_location(program, "u_order").unwrap();
        let rgba_location = gl.get_uniform_location(program, "u_rgba").unwrap();
        // let radius_location = gl.get_uniform_location(program, "u_radius").unwrap();
        // let rotation_0_location = gl.get_uniform_location(program, "u_rotation_0").unwrap();
        // let rotation_1_location = gl.get_uniform_location(program, "u_rotation_1").unwrap();
        // let rotation_2_location = gl.get_uniform_location(program, "u_rotation_2").unwrap();
        
        gl.use_program(Some(program));
        // set uniforms
        for idx in 0..CACHE_CAPACITY {
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
            // let arr: [f32; 16]= [
            //         0.0, 1.0, 0.0, 0.0,
            //         -1.0, 0.0, 0.0, 0.0,
            //         0.0, 0.0, 1.0, 0.0,
            //         0.0, 0.0, 0.0, 1.0,
            //     ];
            // log(&format!("{:?}", arr));
            // gl.uniform_matrix4fv_with_f32_array(
            //     Some(&rotation_0_location),
            //     false,
            //     &arr                //&uniforms.rotations[idx][0]
            // );
            // gl.uniform4fv_with_f32_array(Some(&rotation_1_location), &uniforms.rotations[idx][1]);
            // gl.uniform4fv_with_f32_array(Some(&rotation_2_location), &uniforms.rotations[idx][2]);
            gl.uniform4f(
                Some(&translation_location), 
                uniforms.translations[idx][0],
                uniforms.translations[idx][1],
                uniforms.translations[idx][2],
                1.0
            );

            // log("hey wow wow wow");
            log(&format!("TRANS {idx}: {:?}", js_sys::JSON::stringify(&gl.get_uniform(program, &translation_location))));
            // log(&format!("TRANS uniforms {idx}: {:?}", uniforms.translations[idx]));
            // log(&format!("RGBA {idx}: {:?}", js_sys::JSON::stringify(&gl.get_uniform(program, &rgba_location))));
            // log(&format!("RGBA uniforms {idx}: {:?}", uniforms.rgbas[idx]));

            // log(&format!("ROT 0 {idx}: {:?}", js_sys::JSON::stringify(&gl.get_uniform(program, &rotation_0_location))));

            // Draw the geometry.
            let _offset = idx * 300;
            let _count = Geometry::to_vertex_count(shapes[idx]);
            let idx_left: usize = match idx {
                0 => 0,
                1 => 2,
                2 => 4,
                3 => 6,
                4 => 8,
                5 => 0,
                6 => 2,
                7 => 4,
                8 => 6,
                9 => 8,
                10 => 0,
                11 => 2,
                12 => 4,
                13 => 6,
                14 => 8,
                15 => 3,
                _ => 0
            };

            gl.draw_arrays(WebGl2RenderingContext::LINES, idx_left as i32, 2);//offset as i32, count as i32);
        }
        Ok(())
    }
}


