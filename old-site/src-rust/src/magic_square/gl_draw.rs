use super::geometry::cache::CACHE_CAPACITY;
use super::geometry::geom::Geom;
use super::geometry::transformations::Projection;
use super::geometry::Shapes;
use super::gl_uniforms::{GlUniforms, UniformLocations};
use super::settings::TransformOrder;
use wasm_bindgen::prelude::*;
use web_sys::WebGl2RenderingContext;
// use super::main::log;

pub struct GlDraw;

impl GlDraw {
    pub fn scene(
        gl: &WebGl2RenderingContext,
        uniforms: &GlUniforms,
        u_locs: &UniformLocations,
        shapes: &Shapes,
        order: &TransformOrder,
        _x: &f32,
    ) -> Result<(), JsValue> {
        // NOTE FOR DEBUGGING
        // - the uniform name "my_uniform" is defined in the shader source text where the uniform is defined
        // let uniform_location = gl.get_uniform_location(program, "my_uniform").unwrap();
        // log(&format!("uniform CPU{idx}: {:?}", uniforms.my_uniforms[idx]));
        // log(&format!("uniform GPU{idx}: {:?}", js_sys::JSON::stringify(&gl.get_uniform(program, &uniform_location))));

        // set uniforms
        for idx in 0..CACHE_CAPACITY {
            gl.uniform4f(
                Some(&u_locs.rgba),
                uniforms.rgbas[idx][0],
                uniforms.rgbas[idx][1],
                uniforms.rgbas[idx][2],
                uniforms.rgbas[idx][3],
            );
            gl.uniform1i(
                Some(&u_locs.order),
                match order {
                    TransformOrder::RotateThenTranslate => 1,
                    TransformOrder::TranslateThenRotate => 0,
                },
            );
            gl.uniform_matrix4fv_with_f32_array(
                Some(&u_locs.proj_z_zero),
                false,
                &Projection::z_zero(),
            );
            gl.uniform_matrix4fv_with_f32_array(Some(&u_locs.radius), false, &uniforms.radii[idx]);
            gl.uniform_matrix4fv_with_f32_array(
                Some(&u_locs.rotation_zero),
                false,
                &uniforms.rotations[idx][0],
            );
            gl.uniform_matrix4fv_with_f32_array(
                Some(&u_locs.rotation_one),
                false,
                &uniforms.rotations[idx][1],
            );
            gl.uniform_matrix4fv_with_f32_array(
                Some(&u_locs.rotation_two),
                false,
                &uniforms.rotations[idx][2],
            );
            gl.uniform4f(
                Some(&u_locs.translation),
                uniforms.translations[idx][0],
                uniforms.translations[idx][1],
                uniforms.translations[idx][2],
                1.0,
            );
            let offset_vc: (i32, i32) = Geom::into_offset_vc(shapes[idx]);
            // TODO: Mosh DrawPattern Setting
            let offset = offset_vc.0; // 6 * (50.0 * (x/4.0).sin()).abs().floor() as i32;
            let vert_count = offset_vc.1; // i32 = 100;
            gl.draw_arrays(WebGl2RenderingContext::LINES, offset, vert_count); //offset as i32, count as i32);
        }
        Ok(())
    }
}
