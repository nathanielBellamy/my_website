use crate::magic_square::settings::Settings;
use crate::magic_square::ui_manifest::{
    INPUT_IDS,
    INPUT_COLOR_ORIGIN, 
    INPUT_COLOR_NW, INPUT_COLOR_NE, INPUT_COLOR_SE, INPUT_COLOR_SW
};
use crate::magic_square::main::MagicSquare;

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
        let val = val;
        let r: f32 = 0.0;
        let g: f32 = 0.0;
        let b: f32 = 0.0;
        let a: f32 = 0.0;
        match input_id.as_str() {
            // INPUT_A => self.function.a = val.parse::<f64>().unwrap(),
            INPUT_COLOR_ORIGIN => {
                self.settings.color_origin_r = r;
                self.settings.color_origin_g = g;
                self.settings.color_origin_b = b;
                self.settings.color_origin_a = a;
            },
            INPUT_COLOR_NW  => {
                self.settings.color_nw_r = r;
                self.settings.color_nw_g = g;
                self.settings.color_nw_b = b;
                self.settings.color_nw_a = a;
            },
            INPUT_COLOR_NE => {
                self.settings.color_ne_r = r;
                self.settings.color_ne_g = g;
                self.settings.color_ne_b = b;
                self.settings.color_ne_a = a;
            },
            INPUT_COLOR_SE => {
                self.settings.color_se_r = r;
                self.settings.color_se_g = g;
                self.settings.color_se_b = b;
                self.settings.color_se_a = a;
            },
            INPUT_COLOR_SW => {
                self.settings.color_sw_r = r;
                self.settings.color_sw_g = g;
                self.settings.color_sw_b = b;
                self.settings.color_sw_a = a;
            },
            _ => {}
        }
    }

    pub fn set_ui_initial_values(&self) {
        for id in INPUT_IDS.iter() {
            self.set_ui_initial_value(id.to_string());
        }
    }

    pub fn set_ui_initial_value(&self, input_id: String) {
        let element: web_sys::Element = MagicSquare::document().get_element_by_id(&input_id)
            .expect("to get element {input_id}");
        let val: String = match input_id.as_str() {
            INPUT_COLOR_ORIGIN => self.settings.origin_rgba_string(),
            INPUT_COLOR_NW => self.settings.nw_rgba_string(),
            INPUT_COLOR_NE => self.settings.ne_rgba_string(),
            INPUT_COLOR_SE => self.settings.se_rgba_string(),
            INPUT_COLOR_SW => self.settings.sw_rgba_string(),
            _ => "-1".to_string()
        };
        element.set_attribute("value", &val).expect("to set attribute value on {input_id}"); 
    }
}
