use serde::{Deserialize, Serialize};
use crate::magic_square::main::log;
use crate::magic_square::settings::Settings;
use crate::magic_square::ui_manifest::{
    INPUT_COLORS,
    INPUT_COLOR_DIRECTION, INPUT_COLOR_GRADIENT, INPUT_COLOR_SPEED,
    // INPUT_COLOR_1, INPUT_COLOR_2, INPUT_COLOR_3, INPUT_COLOR_4, INPUT_COLOR_5, INPUT_COLOR_6, INPUT_COLOR_7, INPUT_COLOR_8,
    INPUT_DRAW_PATTERN_TYPE, INPUT_DRAW_PATTERN_COUNT, INPUT_DRAW_PATTERN_OFFSET, INPUT_DRAW_PATTERN_SPEED,
    INPUT_MOUSE_TRACKING,
    INPUT_SHAPES,
    INPUT_RADIUS_BASE, INPUT_RADIUS_STEP,
    INPUT_TRANSFORM_ORDER,
    INPUT_X_ROT_BASE, INPUT_Y_ROT_BASE, INPUT_Z_ROT_BASE,
    INPUT_X_ROT_SPREAD, INPUT_Y_ROT_SPREAD, INPUT_Z_ROT_SPREAD,
    INPUT_X_AXIS_X_ROT_COEFF, INPUT_X_AXIS_Y_ROT_COEFF, INPUT_X_AXIS_Z_ROT_COEFF,
    INPUT_Y_AXIS_X_ROT_COEFF, INPUT_Y_AXIS_Y_ROT_COEFF, INPUT_Y_AXIS_Z_ROT_COEFF,
    INPUT_TRANSLATION_X_BASE, INPUT_TRANSLATION_Y_BASE, INPUT_TRANSLATION_Z_BASE,
    INPUT_TRANSLATION_X_SPREAD, INPUT_TRANSLATION_Y_SPREAD, INPUT_TRANSLATION_Z_SPREAD,
    INPUT_LFO_1_ACTIVE, INPUT_LFO_1_AMP, INPUT_LFO_1_DEST, INPUT_LFO_1_FREQ, INPUT_LFO_1_PHASE, INPUT_LFO_1_SHAPE,
    INPUT_LFO_2_ACTIVE, INPUT_LFO_2_AMP, INPUT_LFO_2_DEST, INPUT_LFO_2_FREQ, INPUT_LFO_2_PHASE, INPUT_LFO_2_SHAPE,
    INPUT_LFO_3_ACTIVE, INPUT_LFO_3_AMP, INPUT_LFO_3_DEST, INPUT_LFO_3_FREQ, INPUT_LFO_3_PHASE, INPUT_LFO_3_SHAPE,
    INPUT_LFO_4_ACTIVE, INPUT_LFO_4_AMP, INPUT_LFO_4_DEST, INPUT_LFO_4_FREQ, INPUT_LFO_4_PHASE, INPUT_LFO_4_SHAPE,
};
use super::geometry::cache::{Cache, CACHE_CAPACITY};
use super::main::Rgba;
use super::shader_compiler::ShaderCompiler;
use super::settings::{Colors, IndexedGradient};

pub const EMPTY_COLORS: Colors = [[0.0;4]; 16];

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub struct UiBuffer {
    pub settings: Settings,
}

impl UiBuffer {
    pub fn new() -> UiBuffer {
        UiBuffer {
            settings: Settings::new(),
        }
    }

    pub fn from_prev_settings(prev_settings: Settings) -> UiBuffer {
        let mut colors: Colors = EMPTY_COLORS;
        prev_settings.colors.iter().enumerate().map(|(idx, x)| {
            colors[idx] = UiBuffer::convert_rgba(*x)
        });
        UiBuffer { 
            settings: Settings {
                colors,
                ..prev_settings 
            }
        }
    }

    pub fn convert_rgba(rgba: [f32; 4]) -> [f32; 4] {
        let mut res: [f32; 4] = [0.0, 0.0, 0.0, 0.0];
        res[0] = rgba[0] / 255.0;
        res[1] = rgba[1] / 255.0;
        res[2] = rgba[2] / 255.0;
        res[3] = rgba[3];
        res
    }

    pub fn copy(&self) -> UiBuffer {
        UiBuffer { settings: Settings {..self.settings} }
    }

    pub fn color_gradient_at_step(&self, step: u8, idx_a: usize, idx_b:usize) -> Rgba {
        let mut result = [0.0, 0.0, 0.0, 0.0];
        let width: usize = idx_b - idx_a;
        if width > 0 && step < width as u8 {
            let t: f32 = step as f32 / width as f32;
            
            for idx in 0..4 as usize {
                result[idx] = (1.0 - t) * self.settings.colors[idx_a][idx]
                                + t * self.settings.colors[idx_b][idx];
            }
        }
        result
    }

    pub fn set_color_gradient(&mut self, ig: IndexedGradient){
        let mut idx: usize = 1;
        while ig.idx_a + idx < ig.idx_b {
            self.settings.colors[idx] = self.color_gradient_at_step(idx as u8, ig.idx_a,  ig.idx_b);
            idx += 1
        }
    }

    pub fn update_frag_shader_cache(&self, frag_shader_cache: &mut Vec<String>) {
        for (idx, shader) in frag_shader_cache.iter_mut().enumerate() {
            *shader = ShaderCompiler::into_frag_shader_string(&self.settings.colors[idx]);
        }
    }

    pub fn update(
            &mut self, 
            input_id: String, 
            val: String, 
            frag_shader_cache: &mut Vec<String>,
            geometry_cache: &mut Cache,
            // animation: &mut Animation,
    ) {
        // log(&input_id);
        // log(&val);
        match input_id.as_str() {
            INPUT_COLOR_DIRECTION => {
                if let Ok(val) = Settings::try_into_color_direction(val) {
                    self.settings.color_direction = val
                }
            },
            INPUT_COLOR_GRADIENT => {
                if let Ok(indexed_gradient) = Settings::try_into_indexed_gradient(val) {
                    self.set_color_gradient(indexed_gradient);
                    self.update_frag_shader_cache(frag_shader_cache);
                }
            },
            INPUT_COLOR_SPEED => {
                if let Ok(val) = val.parse::<u8>() {
                    self.settings.color_speed = val
                }
            },
            INPUT_COLORS => {
                if let Ok(indexed_color) = Settings::try_into_indexed_color(val) {
                    self.settings.colors[indexed_color.index] = indexed_color.rgba;
                    self.update_frag_shader_cache(frag_shader_cache);
                }
            },
            INPUT_SHAPES => {
                if let Ok(indexed_shape) = Settings::try_into_indexed_shape(val) {
                    if indexed_shape.index == 16 { // magic number indicating update all
                        for shape in self.settings.shapes.iter_mut() {
                            *shape = indexed_shape.shape;
                        }
                    } else {
                        self.settings.shapes[indexed_shape.index] = indexed_shape.shape;
                    }
                }
            },
            INPUT_DRAW_PATTERN_TYPE => {
                if let Ok(draw_pattern_type) = Settings::try_into_draw_pattern_type(val) {
                    log(&format!("{:?}", draw_pattern_type));
                    geometry_cache.clear();
                    self.settings.draw_pattern_type = draw_pattern_type;
                }
            },
            INPUT_DRAW_PATTERN_COUNT => {
                if let Ok(count) = val.parse::<i32>() {
                    if count > 0 && count < 17 {
                        geometry_cache.clear();
                        self.settings.draw_pattern_count = count;
                    }
                }
            },
            INPUT_DRAW_PATTERN_OFFSET => {
                if let Ok(offset) = val.parse::<i32>() {
                    if offset > -1 && offset < CACHE_CAPACITY as i32 {
                        geometry_cache.clear();
                        self.settings.draw_pattern_offset = offset;
                    }
                }
            },
            INPUT_DRAW_PATTERN_SPEED => {
                if let Ok(speed) = val.parse::<i32>() {
                    if speed > -1 && speed < 21 {
                        geometry_cache.clear();
                        self.settings.draw_pattern_speed = speed;
                    }
                }
            },
            INPUT_MOUSE_TRACKING => {
                if let Ok(mouse_traking) = Settings::try_into_mouse_tracking(val) {
                    self.settings.mouse_tracking = mouse_traking            
                }
            },
            INPUT_RADIUS_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.radius_base = val;
                }
            },
            INPUT_RADIUS_STEP => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.radius_step = val;
                }
            },
            INPUT_TRANSFORM_ORDER => {
                if let Ok(val) = Settings::try_into_transform_order(val) {
                    self.settings.transform_order = val;
                }
            },
            INPUT_X_ROT_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.x_rot_base = val;
                }
            },
            INPUT_Y_ROT_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.y_rot_base = val;
                }
            },
            INPUT_Z_ROT_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.z_rot_base = val;
                }
            },
            INPUT_X_ROT_SPREAD => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.x_rot_spread = val;
                }
            },
            INPUT_Y_ROT_SPREAD => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.y_rot_spread = val;
                }
            },
            INPUT_Z_ROT_SPREAD => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.z_rot_spread = val;
                }
            },
            INPUT_X_AXIS_X_ROT_COEFF => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.x_axis_x_rot_coeff = val;
                }
            }, 
            INPUT_X_AXIS_Y_ROT_COEFF => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.x_axis_y_rot_coeff = val;
                }
            }, 
            INPUT_X_AXIS_Z_ROT_COEFF => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.x_axis_z_rot_coeff = val;
                }
            },
            INPUT_Y_AXIS_X_ROT_COEFF => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.y_axis_x_rot_coeff = val;
                }
            }, 
            INPUT_Y_AXIS_Y_ROT_COEFF => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.y_axis_y_rot_coeff = val;
                }
            }, 
            INPUT_Y_AXIS_Z_ROT_COEFF => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.y_axis_z_rot_coeff = val;
                }
            },
            INPUT_TRANSLATION_X_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_x_base = val;
                }
            },
            INPUT_TRANSLATION_X_SPREAD => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_x_spread = val;
                }
            },
            INPUT_TRANSLATION_Y_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_y_base = val;
                }
            },
            INPUT_TRANSLATION_Y_SPREAD => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_y_spread = val;
                }
            },
            INPUT_TRANSLATION_Z_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_z_base = val;
                }
            },
            INPUT_TRANSLATION_Z_SPREAD => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_z_spread = val;
                }
            },
            INPUT_LFO_1_ACTIVE => {
                if let Ok(val) = val.parse::<bool>() {
                    self.settings.lfo_1_active = val;
                }
            },
            INPUT_LFO_1_AMP => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_1_amp = val;
                }
            },
            INPUT_LFO_1_DEST => {
                if let Ok(dest) = Settings::try_into_lfo_destination(val) {
                    self.settings.lfo_1_dest = dest;
                }
            },
            INPUT_LFO_1_FREQ => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_1_freq = val;
                }
            },
            INPUT_LFO_1_PHASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_1_phase = val;
                }
            },
            INPUT_LFO_1_SHAPE => {
                if let Ok(shape) = Settings::try_into_lfo_shape(val) {
                    self.settings.lfo_1_shape = shape;
                }
            },
            INPUT_LFO_2_ACTIVE => {
                if let Ok(val) = val.parse::<bool>() {
                    self.settings.lfo_2_active = val;
                }
            },
            INPUT_LFO_2_AMP => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_2_amp = val;
                }
            },
            INPUT_LFO_2_DEST => {
                if let Ok(dest) = Settings::try_into_lfo_destination(val) {
                    self.settings.lfo_2_dest = dest;
                }
            },
            INPUT_LFO_2_FREQ => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_2_freq = val;
                }
            },
            INPUT_LFO_2_PHASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_2_phase = val;
                }
            },
            INPUT_LFO_2_SHAPE => {
                if let Ok(shape) = Settings::try_into_lfo_shape(val) {
                    self.settings.lfo_2_shape = shape;
                }
            },
            INPUT_LFO_3_ACTIVE => {
                if let Ok(val) = val.parse::<bool>() {
                    self.settings.lfo_3_active = val;
                }
            },
            INPUT_LFO_3_AMP => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_3_amp = val;
                }
            },
            INPUT_LFO_3_DEST => {
                if let Ok(dest) = Settings::try_into_lfo_destination(val) {
                    self.settings.lfo_3_dest = dest;
                }
            },
            INPUT_LFO_3_FREQ => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_3_freq = val;
                }
            },
            INPUT_LFO_3_PHASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_3_phase = val;
                }
            },
            INPUT_LFO_3_SHAPE => {
                if let Ok(shape) = Settings::try_into_lfo_shape(val) {
                    self.settings.lfo_3_shape = shape;
                }
            },
            INPUT_LFO_4_ACTIVE => {
                if let Ok(val) = val.parse::<bool>() {
                    self.settings.lfo_4_active = val;
                }
            },
            INPUT_LFO_4_AMP => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_4_amp = val;
                }
            },
            INPUT_LFO_4_DEST => {
                if let Ok(dest) = Settings::try_into_lfo_destination(val) {
                    self.settings.lfo_4_dest = dest;
                }
            },
            INPUT_LFO_4_FREQ => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_4_freq = val;
                }
            },
            INPUT_LFO_4_PHASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_4_phase = val;
                }
            },
            INPUT_LFO_4_SHAPE => {
                if let Ok(shape) = Settings::try_into_lfo_shape(val) {
                    self.settings.lfo_4_shape = shape;
                }
            },
            _ => {}

        }
    }
}

