use crate::magic_square::main::Rgba;
use crate::magic_square::lfo::{LfoDestination, LfoShape};
use serde::{Deserialize, Serialize};
use wasm_bindgen::JsValue;

use super::geometry::Shape;
use super::geometry::cache::CACHE_CAPACITY;
// use crate::magic_square::main::log;

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub enum DrawPatternType {
    #[default]
    Fix,
    In,
    Out,
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
pub enum TransformOrder{
    #[default]
    RotateThenTranslate,
    TranslateThenRotate,
}

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub enum ColorDirection{
    In,
    #[default]
    Fix,
    Out,
}

pub type Colors = [Rgba; CACHE_CAPACITY];


#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub struct ColorGradient {
    idx_a: i32,
    idx_b: i32
}

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub struct Settings {
    // COLOR
    pub colors: Colors,
    pub color_direction: ColorDirection,
    pub color_speed: u8,

    // DRAW PATTERN
    pub draw_pattern_type: DrawPatternType,
    pub draw_pattern_count: i32,
    pub draw_pattern_speed: i32,
    pub draw_pattern_offset: i32,

    // GEOMETRY
    pub radius_base: f32,
    pub radius_step: f32,
    pub shapes: [Shape; CACHE_CAPACITY],
    pub transform_order: TransformOrder,

    // lfo_1
    pub lfo_1_active: bool,
    pub lfo_1_amp: f32,
    pub lfo_1_dest: LfoDestination,
    pub lfo_1_freq: f32,
    pub lfo_1_phase: f32,
    pub lfo_1_shape: LfoShape,

    // lfo_2
    pub lfo_2_active: bool,
    pub lfo_2_amp: f32,
    pub lfo_2_dest: LfoDestination,
    pub lfo_2_freq: f32,
    pub lfo_2_phase: f32,
    pub lfo_2_shape: LfoShape,

    // lfo_3
    pub lfo_3_active: bool,
    pub lfo_3_amp: f32,
    pub lfo_3_dest: LfoDestination,
    pub lfo_3_freq: f32,
    pub lfo_3_phase: f32,
    pub lfo_3_shape: LfoShape,

    // lfo_4
    pub lfo_4_active: bool,
    pub lfo_4_amp: f32,
    pub lfo_4_dest: LfoDestination,
    pub lfo_4_freq: f32,
    pub lfo_4_phase: f32,
    pub lfo_4_shape: LfoShape,

    // PRESET
    pub preset: usize,

    // TODO:
    // mouse settings
    // MouseFollow - Always, Click + Drag, DoubleClick On/Off

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
    pub translation_x_base: f32,
    pub translation_x_spread: f32,
    pub translation_y_base: f32,
    pub translation_y_spread: f32,
    pub translation_z_base: f32,
    pub translation_z_spread: f32,
    pub mouse_tracking: MouseTracking,
}


impl Settings {
    pub fn new() -> Settings {
        Settings {
            // COLOR
            color_direction: ColorDirection::Fix,
            color_speed: 17,
            colors: [
                [1.0, 1.0, 1.0, 1.0],
                [0.0, 0.0, 0.0, 1.0],
                [1.0, 1.0, 1.0, 1.0],
                [0.0, 0.0, 0.0, 1.0],
                [1.0, 1.0, 1.0, 1.0],
                [0.0, 0.0, 0.0, 1.0],
                [1.0, 1.0, 1.0, 1.0],
                [0.0, 0.0, 0.0, 1.0],
                [1.0, 1.0, 1.0, 1.0],
                [0.0, 0.0, 0.0, 1.0],
                [1.0, 1.0, 1.0, 1.0],
                [0.0, 0.0, 0.0, 1.0],
                [1.0, 1.0, 1.0, 1.0],
                [0.0, 0.0, 0.0, 1.0],
                [1.0, 1.0, 1.0, 1.0],
                [0.0, 0.0, 0.0, 1.0],
            ],

            // DRAW PATTERN
            draw_pattern_type: DrawPatternType::Out,
            draw_pattern_count: 7,
            draw_pattern_offset: 8,
            draw_pattern_speed: 17,

            // GEOMETRY
            radius_base: 0.1,
            radius_step: 0.1,
            transform_order: TransformOrder::RotateThenTranslate,
            shapes: [Shape::Hexagon; CACHE_CAPACITY],

            // lfo_1
            lfo_1_active: true,
            lfo_1_amp: 0.05,
            lfo_1_dest: LfoDestination::TranslationXSpread,
            lfo_1_freq: 35.0,
            lfo_1_phase: 0.0,
            lfo_1_shape: LfoShape::Sine,

            // lfo_2
            lfo_2_active: true,
            lfo_2_amp: 0.1,
            lfo_2_dest: LfoDestination::RollBase,
            lfo_2_freq: 15.0,
            lfo_2_phase: 0.0,
            lfo_2_shape: LfoShape::Sine,

            // lfo_3
            lfo_3_active: true,
            lfo_3_amp: 0.1,
            lfo_3_dest: LfoDestination::RadiusStep,
            lfo_3_freq: 15.0,
            lfo_3_phase: 0.0,
            lfo_3_shape: LfoShape::Sine,

            // lfo_4
            lfo_4_active: true,
            lfo_4_amp: 0.3,
            lfo_4_dest: LfoDestination::YawSpread,
            lfo_4_freq: 35.0,
            lfo_4_phase: 0.0,
            lfo_4_shape: LfoShape::Sine,

            // PRESET
            preset: 0,
            
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
            translation_x_base: 0.0,
            translation_x_spread: 0.0,
            translation_y_base: 0.0,
            translation_y_spread: 0.0,
            translation_z_base: 0.0,
            translation_z_spread: 0.0,
            mouse_tracking: MouseTracking::Off,
        }
    }

    pub fn try_into_color_direction(cd: String) -> Result<ColorDirection, ()> {
        match cd.as_str() {
            "In" => Ok(ColorDirection::In),
            "Fix" => Ok(ColorDirection::Fix),
            "Out" => Ok(ColorDirection::Out),
            _ => Err(())
        }
    }
 
    pub fn try_into_io_gradient(val: String) -> Result<IOGradient, JsValue> {
        let val = js_sys::JSON::parse(&val).unwrap();
        let res: IOGradient = serde_wasm_bindgen::from_value(val)?;
        Ok(res)
    }

    pub fn try_into_lfo_destination(dest: String) -> Result<LfoDestination, ()> {
        match dest.as_str() {
            // rotation
            "PitchBase" => Ok(LfoDestination::PitchBase),
            "PitchSpread" => Ok(LfoDestination::PitchSpread),
            "PitchX" => Ok(LfoDestination::PitchX),
            "PitchY" => Ok(LfoDestination::PitchY),
            "RollBase" => Ok(LfoDestination::RollBase),
            "RollSpread" => Ok(LfoDestination::RollSpread),
            "RollX" => Ok(LfoDestination::RollX),
            "RollY" => Ok(LfoDestination::RollY),
            "YawBase" => Ok(LfoDestination::YawBase),
            "YawSpread" => Ok(LfoDestination::YawSpread),
            "YawX" => Ok(LfoDestination::YawX),
            "YawY" => Ok(LfoDestination::YawY),

            // radius
            "RadiusBase" => Ok(LfoDestination::RadiusBase),
            "RadiusStep" => Ok(LfoDestination::RadiusStep),

            // translation
            "TranslationXBase" => Ok(LfoDestination::TranslationXBase),
            "TranslationXSpread" => Ok(LfoDestination::TranslationXSpread),
            "TranslationYBase" => Ok(LfoDestination::TranslationYBase),
            "TranslationYSpread" => Ok(LfoDestination::TranslationYSpread),
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

    pub fn try_into_draw_pattern_type(pt: String) -> Result<DrawPatternType, ()> {
        // log(&pt);
        match pt.as_str() {
            "Fix" => Ok(DrawPatternType::Fix),
            "Out" => Ok(DrawPatternType::Out),
            "In" => Ok(DrawPatternType::In),
            _ => Err(())
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

    pub fn try_into_transform_order(order: String) -> Result<TransformOrder, ()> {
        match order.as_str() {
            "RotateThenTranslate" => Ok(TransformOrder::RotateThenTranslate),
            "TranslateThenRotate" => Ok(TransformOrder::TranslateThenRotate),
            _ => Err(())
        }
    }

    pub fn try_into_preset_idx(val: String) -> Result<usize, ()> {
        let u: usize = match val.parse::<usize>() {
            Ok(val) => val,
            Err(_) => 0
        };
        match u < 64 {
            true => Ok(u),
            false => Ok(0)
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

    pub fn try_into_io_shape(val: String) -> Result<IOShape, JsValue> {
        let val = js_sys::JSON::parse(&val).unwrap();
        let res: IOShape = serde_wasm_bindgen::from_value(val)?;
        Ok(res)
    }

    pub fn try_into_io_color(val: String) -> Result<IOColor, JsValue> {
        let val = js_sys::JSON::parse(&val).unwrap();
        let mut res: IOColor = serde_wasm_bindgen::from_value(val)?;
        for i in 0..3 {
            res.rgba[i] = res.rgba[i] / 255.0
        }
        Ok(res)
    }

    pub fn try_into_io_preset(val: String) -> Result<IOPreset, JsValue> {
        let val = js_sys::JSON::parse(&val).unwrap();
        let res: IOPreset = serde_wasm_bindgen::from_value(val)?;
        Ok(res)     
    }


    pub fn try_into_shape(val: String) -> Result<Shape, ()> {
        match val.as_str() {
            "Hexagon" => Ok(Shape::Hexagon),
            "Icosahedron" => Ok(Shape::Icosahedron),
            _ => Err(())
        }
    }
}

#[derive(Clone, Deserialize, Serialize, Debug)]
pub struct IOColor {
    pub rgba: Rgba, 
    pub idx: usize
}


#[derive(Clone, Deserialize, Serialize, Debug)]
pub struct IOGradient {
    pub idx_a: usize,
    pub idx_b: usize,
}

#[derive(Clone, Deserialize, Serialize, Debug)]
pub struct IOPreset {
    pub preset: usize,
    pub action: IOPresetAction
}

#[derive(Clone, Deserialize, Serialize, Debug)]
pub enum IOPresetAction {
    Load,
    Save,
    Set,
    Show,
}

#[derive(Clone, Deserialize, Serialize, Debug)]
pub struct IOShape {
    pub shape: Shape, 
    pub index: usize
}
