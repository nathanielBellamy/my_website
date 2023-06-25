use std::rc::Rc;
use std::cell::RefCell;
use web_sys::{WebGl2RenderingContext, WebGlUniformLocation, WebGlProgram};
use super::main::{Axis, Rgba};
// use super::main::log;
use super::geometry::transformations::{Mat4, Rotation, Translation, MAT4_ID};
use super::geometry::cache::CACHE_CAPACITY;
use super::settings::MouseTracking;
use super::settings::Settings;

pub struct UniformLocations {
    pub order: WebGlUniformLocation,
    pub proj_z_zero: WebGlUniformLocation,
    pub radius: WebGlUniformLocation,
    pub rgba: WebGlUniformLocation,
    pub rotation_zero: WebGlUniformLocation,
    pub rotation_one: WebGlUniformLocation,
    pub rotation_two: WebGlUniformLocation,
    pub translation: WebGlUniformLocation,
}

impl UniformLocations {
    pub fn new(gl: &WebGl2RenderingContext, program: &WebGlProgram) -> UniformLocations { 
        UniformLocations {    
            order: gl.get_uniform_location(program, "u_order").unwrap(),
            proj_z_zero: gl.get_uniform_location(program, "u_projection_z_zero").unwrap(),
            translation: gl.get_uniform_location(program, "u_translation").unwrap(),
            radius: gl.get_uniform_location(program, "u_radius").unwrap(),
            rgba: gl.get_uniform_location(program, "u_rgba").unwrap(),
            rotation_zero: gl.get_uniform_location(program, "u_rotation_zero").unwrap(),
            rotation_one: gl.get_uniform_location(program, "u_rotation_one").unwrap(),
            rotation_two: gl.get_uniform_location(program, "u_rotation_two").unwrap(),
        }
    }
}

#[derive(Clone, Copy, Debug)]
pub struct GlUniforms {
    pub radii: [Mat4; CACHE_CAPACITY],
    pub rgbas: [Rgba; CACHE_CAPACITY],
    pub rotations: [[Mat4; 3]; CACHE_CAPACITY],
    pub translations: [[f32; 3]; CACHE_CAPACITY],
}

impl GlUniforms {
    pub fn new() -> GlUniforms {
        GlUniforms {
            radii: [MAT4_ID; CACHE_CAPACITY],
            rgbas: [[0.0; 4]; CACHE_CAPACITY],
            rotations: [[MAT4_ID; 3]; CACHE_CAPACITY],
            translations: [[0.0; 3]; CACHE_CAPACITY],
        }
    }

    pub fn set_uniforms(
        &mut self,
        mouse_pos_buffer: &Rc<RefCell<[f32; 2]>>,
        settings: &Settings,
        color_offset: u8,
    ) {
        // let max_idx = Settings::max_idx_from_draw_pattern(settings.draw_pattern);
        let mouse_pos_buffer = *mouse_pos_buffer.clone().borrow();

        // rgbas
        // TODO: might not need this here, just read from ui_buffer elsewhere perhaps
        for (idx, rgba) in settings.colors.iter().enumerate() {
            let new_idx: usize = (idx + color_offset as usize) % CACHE_CAPACITY;
            self.rgbas[new_idx] = *rgba;
        }

        for idx in 0..CACHE_CAPACITY { // geometry_cache.max_idx + 1 { //TODO: settings.cache_per
            // radii
            let radius = settings.radius_base + idx as f32 * settings.radius_step;
            self.radii[idx] = [
                radius, 0.0, 0.0, 0.0,
                0.0, radius, 0.0, 0.0,
                0.0, 0.0, radius, 0.0,
                0.0, 0.0, 0.0, 1.0,
            ];

            // rotations
            let idx_f32 = idx as f32;
            self.rotations[idx][0] = Rotation::new(
                Axis::X, 
                settings.x_rot_base
                    + (mouse_pos_buffer[0] 
                        + settings.translation_x_base) * settings.x_axis_x_rot_coeff
                    + (mouse_pos_buffer[1] 
                        + settings.translation_y_base) * settings.y_axis_x_rot_coeff
                    + idx_f32 * settings.x_rot_spread
            ).matrix();

            self.rotations[idx][1] = Rotation::new(
                Axis::Y,
                settings.y_rot_base
                    + (mouse_pos_buffer[0] 
                        + settings.translation_x_base) * settings.x_axis_y_rot_coeff
                    + (mouse_pos_buffer[1] 
                        + settings.translation_y_base) * settings.y_axis_y_rot_coeff
                    + idx_f32 * settings.y_rot_spread
            ).matrix();

            self.rotations[idx][2] = Rotation::new(
                Axis::Z,
                settings.z_rot_base
                    + (mouse_pos_buffer[0] 
                        + settings.translation_x_base) * settings.x_axis_z_rot_coeff
                    + (mouse_pos_buffer[1] 
                        + settings.translation_y_base) * settings.y_axis_z_rot_coeff
                    + idx_f32 * settings.z_rot_spread
            ).matrix();

            // translations
            self.translations[idx] = match settings.mouse_tracking {
                MouseTracking::On => Translation { 
                    x: settings.translation_x_base
                        + (idx_f32 * settings.translation_x_spread)
                        + mouse_pos_buffer[0], 
                    y: settings.translation_y_base 
                        - (idx_f32 * settings.translation_y_spread)
                        - mouse_pos_buffer[1], 
                    z: settings.translation_z_base 
                        + (idx_f32 * settings.translation_z_spread)
                }.arr(),
                MouseTracking::Off => Translation { 
                    x: settings.translation_x_base
                        + (idx_f32 * settings.translation_x_spread), 
                    y: settings.translation_y_base
                        - (idx_f32 * settings.translation_y_spread), 
                    z: settings.translation_z_base
                        + (idx_f32 * settings.translation_z_spread)
                }.arr(),
                MouseTracking::InvX =>  Translation { 
                    x: settings.translation_x_base 
                        + (idx_f32 * settings.translation_x_spread)
                        - mouse_pos_buffer[0], 
                    y: settings.translation_y_base
                        - (idx_f32 * settings.translation_y_spread)
                        - mouse_pos_buffer[1], 
                    z: settings.translation_z_base
                        + (idx_f32 * settings.translation_z_spread)
                }.arr(),
                MouseTracking::InvY =>  Translation { 
                    x: settings.translation_x_base 
                        + (idx_f32 * settings.translation_x_spread)
                        + mouse_pos_buffer[0], 
                    y: settings.translation_y_base 
                        - (idx_f32 * settings.translation_y_spread)
                        + mouse_pos_buffer[1], 
                    z: settings.translation_z_base
                        + (idx_f32 * settings.translation_z_spread)
                }.arr(),
                MouseTracking::InvXY =>  Translation { 
                    x: settings.translation_x_base
                        + (idx_f32 * settings.translation_x_spread)
                        - mouse_pos_buffer[0], 
                    y: settings.translation_y_base
                        - (idx_f32 * settings.translation_y_spread)
                        + mouse_pos_buffer[1], 
                    z: settings.translation_z_base
                        + (idx_f32 * settings.translation_z_spread)
                }.arr(),
            };
            
            // let shape: Shape = animation.reel[animation_idx][idx];

        }
    }
}
