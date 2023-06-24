use wasm_bindgen::prelude::*;
use web_sys::{WebGl2RenderingContext, WebGlProgram};
use super::geometry::Shapes;
use super::geometry::cache::CACHE_CAPACITY;
use super::gl_uniforms::GlUniforms;
use super::settings::TransformOrder;
use super::transformations::Projection;
// use super::main::log;

pub struct GlDraw;

impl GlDraw {
    pub fn scene(
        gl: &WebGl2RenderingContext,
        program: &WebGlProgram,
        uniforms: &GlUniforms,
        shapes: &Shapes,
        order: TransformOrder,
    ) -> Result<(), JsValue>{
        // NOTE FOR DEBUGGING
        // - the uniform name "my_uniform" is defined in the shader source text where the uniform is defined
        // let uniform_location = gl.get_uniform_location(program, "my_uniform").unwrap();
        // log(&format!("uniform CPU{idx}: {:?}", uniforms.my_uniforms[idx]));
        // log(&format!("uniform GPU{idx}: {:?}", js_sys::JSON::stringify(&gl.get_uniform(program, &uniform_location))));
        
        // lookup uniforms
        let translation_location = gl.get_uniform_location(program, "u_translation").unwrap();
        let order_location = gl.get_uniform_location(program, "u_order").unwrap();
        let proj_z_zero_location = gl.get_uniform_location(program, "u_projection_z_zero").unwrap();
        let rgba_location = gl.get_uniform_location(program, "u_rgba").unwrap();
        let radius_location = gl.get_uniform_location(program, "u_radius").unwrap();
        let rotation_zero_location = gl.get_uniform_location(program, "u_rotation_zero").unwrap();
        let rotation_one_location = gl.get_uniform_location(program, "u_rotation_one").unwrap();
        let rotation_two_location = gl.get_uniform_location(program, "u_rotation_two").unwrap();
        
        // set uniforms
        for idx in 0..CACHE_CAPACITY {
            gl.uniform4f(
                Some(&rgba_location), 
                uniforms.rgbas[idx][0],
                uniforms.rgbas[idx][1],
                uniforms.rgbas[idx][2],
                uniforms.rgbas[idx][3],
            );
            gl.uniform1i(
                Some(&order_location),
                match order {
                    TransformOrder::RotateThenTranslate => 1,
                    TransformOrder::TranslateThenRotate => 0,
                }
            );
            gl.uniform_matrix4fv_with_f32_array(Some(&proj_z_zero_location), false, &Projection::z_zero());
            gl.uniform_matrix4fv_with_f32_array(Some(&radius_location), false, &uniforms.radii[idx]);
            gl.uniform_matrix4fv_with_f32_array(Some(&rotation_zero_location), false, &uniforms.rotations[idx][0]);
            gl.uniform_matrix4fv_with_f32_array(Some(&rotation_one_location), false, &uniforms.rotations[idx][1]);
            gl.uniform_matrix4fv_with_f32_array(Some(&rotation_two_location), false, &uniforms.rotations[idx][2]);
            gl.uniform4f(
                Some(&translation_location), 
                uniforms.translations[idx][0],
                uniforms.translations[idx][1],
                uniforms.translations[idx][2],
                1.0
            );

            gl.draw_arrays(WebGl2RenderingContext::LINES, 0, 14);//offset as i32, count as i32);
        }
        Ok(())
    }
}


