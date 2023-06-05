use crate::magic_square::main::Rgba;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub enum DrawPattern {
    Fix1,
    Fix2,
    Fix3,
    Fix4,
    Fix5,
    Fix6,
    Fix7,
    Fix8,
    Out1,
    Out2,
    Out3,
    Out4,
    Out5,
    Out6,
    Out7,
    #[default]
    Out8,
    In1,
    In2,
    In3,
    In4,
    In5,
    In6,
    In7,
    In8,
    // TODO:
    //  Conv,
    //  Div,
    //  Random
}

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub enum MouseTracking {
    On,
    #[default]
    Off,
    InvX,
    InvY,
    InvXY
}

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub struct Settings {
    pub draw_pattern: DrawPattern,
    pub mouse_tracking: MouseTracking,

    pub radius_min: f32,
    pub radius_step: f32,

    pub color_1: Rgba,
    pub color_2: Rgba,
    pub color_3: Rgba,
    pub color_4: Rgba,
    pub color_5: Rgba,
    pub color_6: Rgba,
    pub color_7: Rgba,
    pub color_8: Rgba,

    // TODO:
    // mouse settings
    // MouseFollow - Always, Click + Drag, DoubleClick On/Off

    pub x_rot_spread: f32,
    pub y_rot_spread: f32,
    pub z_rot_spread: f32,

    // rotation sensitivity
    pub x_axis_x_rot_coeff: f32,
    pub x_axis_y_rot_coeff: f32,
    pub x_axis_z_rot_coeff: f32,

    pub y_axis_x_rot_coeff: f32,
    pub y_axis_y_rot_coeff: f32,
    pub y_axis_z_rot_coeff: f32,

    pub translation_x: f32,
    pub translation_y: f32,
    pub translation_z: f32,


    // // cache
    // cache_max_idx: usize, // 0..50
    // cache_per: usize,
}

impl Settings {
    pub fn new() -> Settings {
        Settings {
            draw_pattern: DrawPattern::Out8,
            mouse_tracking: MouseTracking::Off,

            x_rot_spread: 0.0,
            y_rot_spread: 0.0, 
            z_rot_spread: 0.0,

            x_axis_x_rot_coeff: 0.0,
            x_axis_y_rot_coeff: -1.0,
            x_axis_z_rot_coeff: 0.0,

            y_axis_x_rot_coeff: 1.0,
            y_axis_y_rot_coeff: 0.0,
            y_axis_z_rot_coeff: 0.0,
            
            radius_min: 0.1,
            radius_step: 0.1,

            color_1: [1.0, 0.0, 1.0, 1.0],
            color_2: [0.0, 1.0, 1.0, 1.0],
            color_3: [1.0, 0.0, 0.5, 1.0],
            color_4: [1.0, 0.1, 1.0, 1.0],
            color_5: [0.0, 0.9, 0.64, 1.0],
            color_6: [0.0, 1.0, 1.0, 1.0],
            color_7: [0.80, 0.44, 0.925, 1.0],
            color_8: [0.0, 0.1, 1.0, 1.0],

            translation_x: 0.0,
            translation_y: 0.0,
            translation_z: 0.0,
        }
    }

    pub fn draw_pattern_from_string(pattern: String) -> DrawPattern {
        match pattern.as_str() {
            "Fix1" => DrawPattern::Fix1,
            "Fix2" => DrawPattern::Fix2,
            "Fix3" => DrawPattern::Fix3,
            "Fix4" => DrawPattern::Fix4,
            "Fix5" => DrawPattern::Fix5,
            "Fix6" => DrawPattern::Fix6,
            "Fix7" => DrawPattern::Fix7,
            "Fix8" => DrawPattern::Fix8,
            "Out1" => DrawPattern::Out1,
            "Out2" => DrawPattern::Out2,
            "Out3" => DrawPattern::Out3,
            "Out4" => DrawPattern::Out4,
            "Out5" => DrawPattern::Out5,
            "Out6" => DrawPattern::Out6,
            "Out7" => DrawPattern::Out7,
            "Out8" => DrawPattern::Out8,
            "In1" => DrawPattern::In1,
            "In2" => DrawPattern::In2,
            "In3" => DrawPattern::In3,
            "In4" => DrawPattern::In4,
            "In5" => DrawPattern::In5,
            "In6" => DrawPattern::In6,
            "In7" => DrawPattern::In7,
            "In8" => DrawPattern::In8,
            _ => DrawPattern::Fix8
        }
    }

    pub fn string_from_draw_pattern(pattern: DrawPattern) -> String {
        match pattern {
            DrawPattern::Fix1 => "Fix1".to_string(), 
            DrawPattern::Fix2 => "Fix2".to_string(), 
            DrawPattern::Fix3 => "Fix3".to_string(), 
            DrawPattern::Fix4 => "Fix4".to_string(), 
            DrawPattern::Fix5 => "Fix5".to_string(), 
            DrawPattern::Fix6 => "Fix6".to_string(), 
            DrawPattern::Fix7 => "Fix7".to_string(), 
            DrawPattern::Fix8 => "Fix8".to_string(), 
            DrawPattern::Out1 => "Out1".to_string(), 
            DrawPattern::Out2 => "Out2".to_string(),
            DrawPattern::Out3 => "Out3".to_string(),
            DrawPattern::Out4 => "Out4".to_string(),
            DrawPattern::Out5 => "Out5".to_string(),
            DrawPattern::Out6 => "Out6".to_string(),
            DrawPattern::Out7 => "Out7".to_string(),
            DrawPattern::Out8 => "Out8".to_string(),
            DrawPattern::In1 => "In1".to_string(),
            DrawPattern::In2 => "In2".to_string(),
            DrawPattern::In3 => "In3".to_string(),
            DrawPattern::In4 => "In4".to_string(),
            DrawPattern::In5 => "In5".to_string(),
            DrawPattern::In6 => "In6".to_string(),
            DrawPattern::In7 => "In7".to_string(),
            DrawPattern::In8 => "In8".to_string(),
            _ => "Fix8".to_string()
        }
    }

    pub fn mouse_tracking_from_string(mt: String) -> MouseTracking {
        match mt.as_str() {
            "On" => MouseTracking::On,
            "Off" => MouseTracking::Off,
            "Inv X" => MouseTracking::InvX,
            "Inv Y" => MouseTracking::InvY,
            "Inv XY" => MouseTracking::InvXY,
            _ => MouseTracking::Off
        }
    }

    pub fn max_idx_from_draw_pattern(pattern: DrawPattern) -> usize {
        match pattern {
                DrawPattern::Fix1 => 1,
                DrawPattern::Fix2 => 2,
                DrawPattern::Fix3 => 3,
                DrawPattern::Fix4 => 4,
                DrawPattern::Fix5 => 5,
                DrawPattern::Fix6 => 6,
                DrawPattern::Fix7 => 7,
                DrawPattern::Fix8 => 8,
                DrawPattern::Out1 => 1,
                DrawPattern::Out2 => 2,
                DrawPattern::Out3 => 3,
                DrawPattern::Out4 => 4,
                DrawPattern::Out5 => 5,
                DrawPattern::Out6 => 6,
                DrawPattern::Out7 => 7,
                DrawPattern::Out8 => 8,
                DrawPattern::In1 => 1,
                DrawPattern::In2 => 2,
                DrawPattern::In3 => 3,
                DrawPattern::In4 => 4,
                DrawPattern::In5 => 5,
                DrawPattern::In6 => 6,
                DrawPattern::In7 => 7,
                DrawPattern::In8 => 8,
        }
    }

    pub fn rgba_string(arr: Rgba) -> String {
        format!(
            "{},{},{},{}",
            arr[0],
            arr[1],
            arr[2],
            arr[3]
        )
    }
}

