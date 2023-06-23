use std::rc::Rc;
use std::cell::RefCell;
use super::main::{Axis, Rgba};
// use super::main::log;
use super::transformations::{Mat4, Rotation, Translation, MAT4_ID, MAT4_ZERO};
use super::geometry::cache::CACHE_CAPACITY;
use super::settings::MouseTracking;
use super::animation::Animation;
use super::settings::Settings;

#[derive(Clone, Copy, Debug)]
pub struct GlUniforms {
    pub radii: [Mat4; CACHE_CAPACITY],
    pub rgbas: [Rgba; CACHE_CAPACITY],
    pub rotations: [[Mat4; 3]; CACHE_CAPACITY],
    pub translations: [Mat4; CACHE_CAPACITY],
}

impl GlUniforms {
    pub fn new() -> GlUniforms {
        GlUniforms {
            radii: [MAT4_ID; CACHE_CAPACITY],
            rgbas: [[0.0; 4]; CACHE_CAPACITY],
            rotations: [[MAT4_ID; 3]; CACHE_CAPACITY],
            translations: [MAT4_ZERO; CACHE_CAPACITY]
        }
    }

    pub fn set_uniforms(
        &mut self,
        mouse_pos_buffer: &Rc<RefCell<[f32; 2]>>,
        settings: &Settings,
    ) {
        // let max_idx = Settings::max_idx_from_draw_pattern(settings.draw_pattern);
        let mouse_pos_buffer = *mouse_pos_buffer.clone().borrow();
        let mut animation = Animation::new();
        animation.set_reel(&settings);

        // rgbas
        // TODO: might not need this here, just read from ui_buffer elsewhere perhaps
        self.rgbas = settings.colors;

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
                }.matrix(),
                MouseTracking::Off => Translation { 
                    x: settings.translation_x_base
                        + (idx_f32 * settings.translation_x_spread), 
                    y: settings.translation_y_base
                        - (idx_f32 * settings.translation_y_spread), 
                    z: settings.translation_z_base
                        + (idx_f32 * settings.translation_z_spread)
                }.matrix(),
                MouseTracking::InvX =>  Translation { 
                    x: settings.translation_x_base 
                        + (idx_f32 * settings.translation_x_spread)
                        - mouse_pos_buffer[0], 
                    y: settings.translation_y_base
                        - (idx_f32 * settings.translation_y_spread)
                        - mouse_pos_buffer[1], 
                    z: settings.translation_z_base
                        + (idx_f32 * settings.translation_z_spread)
                }.matrix(),
                MouseTracking::InvY =>  Translation { 
                    x: settings.translation_x_base 
                        + (idx_f32 * settings.translation_x_spread)
                        + mouse_pos_buffer[0], 
                    y: settings.translation_y_base 
                        - (idx_f32 * settings.translation_y_spread)
                        + mouse_pos_buffer[1], 
                    z: settings.translation_z_base
                        + (idx_f32 * settings.translation_z_spread)
                }.matrix(),
                MouseTracking::InvXY =>  Translation { 
                    x: settings.translation_x_base
                        + (idx_f32 * settings.translation_x_spread)
                        - mouse_pos_buffer[0], 
                    y: settings.translation_y_base
                        - (idx_f32 * settings.translation_y_spread)
                        + mouse_pos_buffer[1], 
                    z: settings.translation_z_base
                        + (idx_f32 * settings.translation_z_spread)
                }.matrix(),
            };
            
            // let shape: Shape = animation.reel[animation_idx][idx];

        }
    }
}
