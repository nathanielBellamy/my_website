use serde::{Deserialize, Serialize};
// use crate::magic_square::main::log;
use crate::magic_square::settings::Settings;
use crate::magic_square::ui_manifest::{
    INPUT_COLOR_1, INPUT_COLOR_2, INPUT_COLOR_3, INPUT_COLOR_4, INPUT_COLOR_5, INPUT_COLOR_6, INPUT_COLOR_7, INPUT_COLOR_8,
    INPUT_DRAW_PATTERN, INPUT_MOUSE_TRACKING,
    INPUT_RADIUS_MIN, INPUT_RADIUS_STEP,
    INPUT_X_ROT_BASE, INPUT_Y_ROT_BASE, INPUT_Z_ROT_BASE,
    INPUT_X_ROT_SPREAD, INPUT_Y_ROT_SPREAD, INPUT_Z_ROT_SPREAD,
    INPUT_X_AXIS_X_ROT_COEFF, INPUT_X_AXIS_Y_ROT_COEFF, INPUT_X_AXIS_Z_ROT_COEFF,
    INPUT_Y_AXIS_X_ROT_COEFF, INPUT_Y_AXIS_Y_ROT_COEFF, INPUT_Y_AXIS_Z_ROT_COEFF,
    INPUT_TRANSLATION_X, INPUT_TRANSLATION_Y, INPUT_TRANSLATION_Z,
    INPUT_LFO_1_ACTIVE, INPUT_LFO_1_AMP, INPUT_LFO_1_DEST, INPUT_LFO_1_FREQ, INPUT_LFO_1_PHASE, INPUT_LFO_1_SHAPE
};


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

    pub fn copy(&self) -> UiBuffer {
        UiBuffer { settings: Settings {..self.settings} }
    }

    pub fn update(&mut self, input_id: String, val: String) {
        // log(&input_id);
        // log(&val);
        match input_id.as_str() {
            INPUT_COLOR_1
                | INPUT_COLOR_2
                | INPUT_COLOR_3
                | INPUT_COLOR_4
                | INPUT_COLOR_5
                | INPUT_COLOR_6
                | INPUT_COLOR_7
                | INPUT_COLOR_8 => {
                    let rgba: Vec<&str> = val.split(",").collect();
                    // validate rgba string
                    if let (Ok(r), Ok(g), Ok(b), Ok(a)) = (
                        rgba[0].parse::<f32>(),
                        rgba[1].parse::<f32>(),
                        rgba[2].parse::<f32>(),
                        rgba[3].parse::<f32>(),
                    ) {
                        let r: f32 = r / 255.0; // CSS uses u8, WebGl uses f32:0.0-1.0
                        let g: f32 = g / 255.0;
                        let b: f32 = b / 255.0;
                        match input_id.as_str() {
                            // INPUT_A => self.function.a = val.parse::<f64>().unwrap(),
                            INPUT_COLOR_1 => self.settings.color_1 = [r,g,b,a],
                            INPUT_COLOR_2 => self.settings.color_2 = [r,g,b,a],
                            INPUT_COLOR_3 => self.settings.color_3 = [r,g,b,a],
                            INPUT_COLOR_4 => self.settings.color_4 = [r,g,b,a],
                            INPUT_COLOR_5 => self.settings.color_5 = [r,g,b,a],
                            INPUT_COLOR_6 => self.settings.color_6 = [r,g,b,a],
                            INPUT_COLOR_7 => self.settings.color_7 = [r,g,b,a],
                            INPUT_COLOR_8 => self.settings.color_8 = [r,g,b,a],
                            _ => {}
                        }
                    }
            },
            INPUT_DRAW_PATTERN => {
                if let Ok(draw_pattern) = Settings::try_into_draw_pattern(val) {
                    self.settings.draw_pattern = draw_pattern            
                }
            },
            INPUT_MOUSE_TRACKING => {
                if let Ok(mouse_traking) = Settings::try_into_mouse_tracking(val) {
                    self.settings.mouse_tracking = mouse_traking            
                }
            },
            INPUT_RADIUS_MIN => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.radius_min = val;
                }
            },
            INPUT_RADIUS_STEP => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.radius_step = val;
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
            INPUT_TRANSLATION_X => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_x = val;
                }
            },
            INPUT_TRANSLATION_Y => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_y = val;
                }
            },
            INPUT_TRANSLATION_Z => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_z = val;
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
            _ => {}

        }
    }
}

