use crate::magic_square::settings::Settings;
use crate::magic_square::ui_manifest::{
    INPUT_IDS,
    INPUT_COLOR_1, INPUT_COLOR_2, INPUT_COLOR_3, INPUT_COLOR_4, INPUT_COLOR_5, INPUT_COLOR_6, INPUT_COLOR_7, INPUT_COLOR_8,
    INPUT_DRAW_PATTERN
};
use crate::magic_square::main::MagicSquare;
// use crate::magic_square::main::log;

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
                    
                    let r: f32 = rgba[0].parse::<f32>().unwrap();
                    let g: f32 = rgba[1].parse::<f32>().unwrap();
                    let b: f32 = rgba[2].parse::<f32>().unwrap();
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
            INPUT_COLOR_1 => Settings::rgba_string(self.settings.color_1),
            INPUT_COLOR_2 => Settings::rgba_string(self.settings.color_2),
            INPUT_COLOR_3 => Settings::rgba_string(self.settings.color_3),
            INPUT_COLOR_4 => Settings::rgba_string(self.settings.color_4),
            INPUT_COLOR_5 => Settings::rgba_string(self.settings.color_5),
            INPUT_COLOR_6 => Settings::rgba_string(self.settings.color_6),
            INPUT_COLOR_7 => Settings::rgba_string(self.settings.color_7),
            INPUT_COLOR_8 => Settings::rgba_string(self.settings.color_8),
            _ => "-1".to_string()
        };
        element.set_attribute("value", &val).expect("to set attribute value on {input_id}"); 
    }
}
