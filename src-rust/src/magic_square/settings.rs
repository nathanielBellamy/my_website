use crate::magic_square::main::Rgba;
use crate::magic_square::lfo::{LfoDestination, LfoShape};
use serde::{Deserialize, Serialize};
// use crate::magic_square::main::log;

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
    // COLOR
    pub color_1: Rgba,
    pub color_2: Rgba,
    pub color_3: Rgba,
    pub color_4: Rgba,
    pub color_5: Rgba,
    pub color_6: Rgba,
    pub color_7: Rgba,
    pub color_8: Rgba,

    // LFO
    pub lfo_1_active: bool,
    pub lfo_1_amp: f32,
    pub lfo_1_dest: LfoDestination,
    pub lfo_1_freq: f32,
    pub lfo_1_phase: f32,
    pub lfo_1_shape: LfoShape,

    // TODO:
    // mouse settings
    // MouseFollow - Always, Click + Drag, DoubleClick On/Off

    // PATTERN
    pub draw_pattern: DrawPattern,

    // RADIUS
    pub radius_min: f32,
    pub radius_step: f32,

    
    // ROTATION
    pub x_rot_base: f32,
    pub y_rot_base: f32,
    pub z_rot_base: f32,

    pub x_rot_spread: f32,
    pub y_rot_spread: f32,
    pub z_rot_spread: f32,

        // rotation sensitivity to mouse movement
    pub x_axis_x_rot_coeff: f32,
    pub x_axis_y_rot_coeff: f32,
    pub x_axis_z_rot_coeff: f32,

    pub y_axis_x_rot_coeff: f32,
    pub y_axis_y_rot_coeff: f32,
    pub y_axis_z_rot_coeff: f32,

    // TRANSLATION
    pub translation_x: f32,
    pub translation_y: f32,
    pub translation_z: f32,
    pub mouse_tracking: MouseTracking,

    // // cache
    // cache_max_idx: usize, // 0..50
    // cache_per: usize,
}

impl Settings {
    pub fn new() -> Settings {
        Settings {
            // COLOR
            color_1: [1.0, 0.0, 1.0, 1.0],
            color_2: [0.0, 1.0, 1.0, 1.0],
            color_3: [1.0, 0.0, 0.5, 1.0],
            color_4: [1.0, 0.1, 1.0, 1.0],
            color_5: [0.0, 0.9, 0.64, 1.0],
            color_6: [0.0, 1.0, 1.0, 1.0],
            color_7: [0.80, 0.44, 0.925, 1.0],
            color_8: [0.0, 0.1, 1.0, 1.0],

            // LFO
            lfo_1_active: true,
            lfo_1_amp: 0.3,
            lfo_1_dest: LfoDestination::TranslationY,
            lfo_1_freq: 35.0,
            lfo_1_phase: 0.0,
            lfo_1_shape: LfoShape::Sine,

            // PATTERN
            draw_pattern: DrawPattern::Out8,
            
            // RADIUS
            radius_min: 0.1,
            radius_step: 0.1,

            // ROTATION
            x_rot_base: 0.0,
            y_rot_base: 0.0,
            z_rot_base: 0.0,

            x_rot_spread: 0.0,
            y_rot_spread: 0.0, 
            z_rot_spread: 0.0,

            x_axis_x_rot_coeff: 0.0,
            x_axis_y_rot_coeff: -1.0,
            x_axis_z_rot_coeff: 0.0,

            y_axis_x_rot_coeff: 1.0,
            y_axis_y_rot_coeff: 0.0,
            y_axis_z_rot_coeff: 0.0,
            
            // TRANSLATION
            translation_x: 0.0,
            translation_y: 0.0,
            translation_z: 0.0,
            mouse_tracking: MouseTracking::Off,
        }
    }

    pub fn try_into_lfo_destination(dest: String) -> Result<LfoDestination, ()> {
        match dest.as_str() {
            "TranslationX" => Ok(LfoDestination::TranslationX),
            "TranslationY" => Ok(LfoDestination::TranslationY),
            "None" => Ok(LfoDestination::None),
            _ => Err(()),
        }
    }

    pub fn try_into_lfo_shape(shape: String) -> Result<LfoShape, ()> {
        match shape.as_str() {
            "Linear" => Ok(LfoShape::Linear),
            "Sine" => Ok(LfoShape::Sine),
            _ => Err(())
        }
    }

    pub fn try_into_draw_pattern(pattern: String) -> Result<DrawPattern, ()> {
        match pattern.as_str() {
            "Fix1" => Ok(DrawPattern::Fix1),
            "Fix2" => Ok(DrawPattern::Fix2),
            "Fix3" => Ok(DrawPattern::Fix3),
            "Fix4" => Ok(DrawPattern::Fix4),
            "Fix5" => Ok(DrawPattern::Fix5),
            "Fix6" => Ok(DrawPattern::Fix6),
            "Fix7" => Ok(DrawPattern::Fix7),
            "Fix8" => Ok(DrawPattern::Fix8),
            "Out1" => Ok(DrawPattern::Out1),
            "Out2" => Ok(DrawPattern::Out2),
            "Out3" => Ok(DrawPattern::Out3),
            "Out4" => Ok(DrawPattern::Out4),
            "Out5" => Ok(DrawPattern::Out5),
            "Out6" => Ok(DrawPattern::Out6),
            "Out7" => Ok(DrawPattern::Out7),
            "Out8" => Ok(DrawPattern::Out8),
            "In1" => Ok(DrawPattern::In1),
            "In2" => Ok(DrawPattern::In2),
            "In3" => Ok(DrawPattern::In3),
            "In4" => Ok(DrawPattern::In4),
            "In5" => Ok(DrawPattern::In5),
            "In6" => Ok(DrawPattern::In6),
            "In7" => Ok(DrawPattern::In7),
            "In8" => Ok(DrawPattern::In8),
            _ => Err(())
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
            DrawPattern::In8 => "In8".to_string()
        }
    }

    pub fn try_into_mouse_tracking(mt: String) -> Result<MouseTracking, ()> {
        match mt.as_str() {
            "On" => Ok(MouseTracking::On),
            "Off" => Ok(MouseTracking::Off),
            "Inv X" => Ok(MouseTracking::InvX),
            "Inv Y" => Ok(MouseTracking::InvY),
            "Inv XY" => Ok(MouseTracking::InvXY),
            _ => Err(())
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

