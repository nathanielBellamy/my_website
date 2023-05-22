use web_sys::HtmlInputElement;
use wasm_bindgen::JsCast;
use crate::magic_square::settings::Settings;
use crate::magic_square::ui_manifest::{
    INPUT_IDS,
    INPUT_COLOR_1, INPUT_COLOR_2, INPUT_COLOR_3, INPUT_COLOR_4, INPUT_COLOR_5, INPUT_COLOR_6, INPUT_COLOR_7, INPUT_COLOR_8,
    INPUT_DRAW_PATTERN, INPUT_MOUSE_TRACKING,
    INPUT_RADIUS_MIN, INPUT_RADIUS_STEP,
    INPUT_X_ROT_COEFF, INPUT_Y_ROT_COEFF, INPUT_Z_ROT_COEFF,
    INPUT_X_ROT_SPREAD, INPUT_Y_ROT_SPREAD, INPUT_Z_ROT_SPREAD,
    INPUT_X_AXIS_X_ROT_COEFF, INPUT_X_AXIS_Y_ROT_COEFF, INPUT_X_AXIS_Z_ROT_COEFF,
    INPUT_Y_AXIS_X_ROT_COEFF, INPUT_Y_AXIS_Y_ROT_COEFF, INPUT_Y_AXIS_Z_ROT_COEFF
};

use crate::magic_square::main::MagicSquare;
use crate::magic_square::main::log;

#[derive(Clone, Copy)]
pub struct UiBuffer {
    pub settings: Settings,
}

impl UiBuffer {
    pub fn new() -> UiBuffer {
        UiBuffer {
            settings: Settings::new(),
        }
    }

    pub fn update(&mut self, input_id: String, val: String) {
        log(&input_id);
        log(&val);
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
                    
                    let r: f32 = rgba[0].parse::<f32>().unwrap() / 255.0;
                    let g: f32 = rgba[1].parse::<f32>().unwrap() / 255.0;
                    let b: f32 = rgba[2].parse::<f32>().unwrap() / 255.0;
                    let a: f32 = rgba[3].parse::<f32>().unwrap();
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
                },
            INPUT_DRAW_PATTERN => {
                self.settings.draw_pattern = Settings::draw_pattern_from_string(val)
            },
            INPUT_MOUSE_TRACKING => {
                self.settings.mouse_tracking = Settings::mouse_tracking_from_string(val)
            },
            INPUT_RADIUS_MIN => {
                self.settings.radius_min = val.parse::<f32>().unwrap() 
            },
            INPUT_RADIUS_STEP => {
                self.settings.radius_step = val.parse::<f32>().unwrap()
            },
            INPUT_X_ROT_COEFF => {
                self.settings.x_rot_coeff = val.parse::<f32>().unwrap()
            },
            INPUT_Y_ROT_COEFF => {
                self.settings.y_rot_coeff = val.parse::<f32>().unwrap()
            },
            INPUT_Z_ROT_COEFF => {
                self.settings.z_rot_coeff = val.parse::<f32>().unwrap()
            },
            INPUT_X_ROT_SPREAD => {
                self.settings.x_rot_spread = val.parse::<f32>().unwrap()
            },
            INPUT_Y_ROT_SPREAD => {
                self.settings.y_rot_spread = val.parse::<f32>().unwrap()
            },
            INPUT_Z_ROT_SPREAD => {
                self.settings.z_rot_spread = val.parse::<f32>().unwrap()
            },
            INPUT_X_AXIS_X_ROT_COEFF => {
                self.settings.x_axis_x_rot_coeff = val.parse::<f32>().unwrap()
            }, 
            INPUT_X_AXIS_Y_ROT_COEFF => {
                self.settings.x_axis_y_rot_coeff = val.parse::<f32>().unwrap()
            }, 
            INPUT_X_AXIS_Z_ROT_COEFF => {
                self.settings.x_axis_z_rot_coeff = val.parse::<f32>().unwrap()
            },
            INPUT_Y_AXIS_X_ROT_COEFF => {
                self.settings.y_axis_x_rot_coeff = val.parse::<f32>().unwrap()
            }, 
            INPUT_Y_AXIS_Y_ROT_COEFF => {
                self.settings.y_axis_y_rot_coeff = val.parse::<f32>().unwrap()
            }, 
            INPUT_Y_AXIS_Z_ROT_COEFF => {
                self.settings.y_axis_z_rot_coeff = val.parse::<f32>().unwrap()
            }
            _ => {}

        }
    }

    pub fn set_ui_initial_values(&self) {
        for id in INPUT_IDS.iter() {
            self.set_ui_initial_value(id.to_string());
        }
    }

    pub fn set_ui_initial_value(&self, input_id: String) {
        let input = MagicSquare::document().get_element_by_id(&input_id)
            .expect("to get element {input_id}")
            .dyn_into::<HtmlInputElement>()
            .expect("unable to get {input_id} from the dom");
        let val: String = match input_id.as_str() {
            INPUT_COLOR_1 => Settings::rgba_string(self.settings.color_1),
            INPUT_COLOR_2 => Settings::rgba_string(self.settings.color_2),
            INPUT_COLOR_3 => Settings::rgba_string(self.settings.color_3),
            INPUT_COLOR_4 => Settings::rgba_string(self.settings.color_4),
            INPUT_COLOR_5 => Settings::rgba_string(self.settings.color_5),
            INPUT_COLOR_6 => Settings::rgba_string(self.settings.color_6),
            INPUT_COLOR_7 => Settings::rgba_string(self.settings.color_7),
            INPUT_COLOR_8 => Settings::rgba_string(self.settings.color_8),
            INPUT_DRAW_PATTERN => Settings::string_from_draw_pattern(self.settings.draw_pattern),
            INPUT_X_ROT_COEFF => format!("{}", self.settings.x_rot_coeff),
            INPUT_Y_ROT_COEFF => format!("{}", self.settings.y_rot_coeff),
            INPUT_Z_ROT_COEFF => format!("{}", self.settings.z_rot_coeff),
            INPUT_X_ROT_SPREAD => format!("{}", self.settings.x_rot_spread),
            INPUT_Y_ROT_SPREAD => format!("{}", self.settings.y_rot_spread),
            INPUT_Z_ROT_SPREAD => format!("{}", self.settings.z_rot_spread),


            _ => "-1".to_string()
        };
        input.set_value(&val); 
    }
}
