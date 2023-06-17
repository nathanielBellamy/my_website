use crate::magic_square::main::Rgba;
use crate::magic_square::lfo::{LfoDestination, LfoShape};
use serde::{Deserialize, Serialize};
use wasm_bindgen::JsValue;

use super::geometry::Shape;
use super::geometry::cache::CACHE_CAPACITY;
use crate::magic_square::main::log;

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

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug, Eq, PartialEq)]
pub enum ColorMode{
    #[default]
    Eight,
    Gradient,
}

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub struct Settings {
    // COLOR
    pub color_direction: ColorDirection,
    pub color_mode: ColorMode,
    pub color_speed: u8,
    pub color_1: Rgba,
    pub color_2: Rgba,
    pub color_3: Rgba,
    pub color_4: Rgba,
    pub color_5: Rgba,
    pub color_6: Rgba,
    pub color_7: Rgba,
    pub color_8: Rgba,

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

    // // cache
    // cache_max_idx: usize, // 0..50
    // cache_per: usize,
}


impl Settings {
    pub fn new() -> Settings {
        Settings {
            // COLOR
            color_direction: ColorDirection::Fix,
            color_mode: ColorMode::Eight,
            color_speed: 17,
            color_1: [1.0, 0.0, 1.0, 1.0],
            color_2: [0.0, 1.0, 1.0, 1.0],
            color_3: [1.0, 0.0, 0.5, 1.0],
            color_4: [1.0, 0.1, 1.0, 1.0],
            color_5: [0.0, 0.9, 0.64, 1.0],
            color_6: [0.0, 1.0, 1.0, 1.0],
            color_7: [0.80, 0.44, 0.925, 1.0],
            color_8: [0.0, 0.1, 1.0, 1.0],

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

    pub fn try_into_color_mode(cm: String) -> Result<ColorMode, ()> {
        match cm.as_str() {
            "Eight" => Ok(ColorMode::Eight),
            "Gradient" => Ok(ColorMode::Gradient),
            _ => Err(())
        }
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
        log("HERHEHERHE WOW ZOW");
        log(&pt);
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

    pub fn rgba_string(arr: Rgba) -> String {
        format!(
            "{},{},{},{}",
            arr[0],
            arr[1],
            arr[2],
            arr[3]
        )
    }

    pub fn try_into_indexed_shape(val: String) -> Result<IndexedShape, ()> {
        let val = js_sys::JSON::parse(&val).unwrap();
        let res: IndexedShapeRaw = serde_wasm_bindgen::from_value(val).unwrap();
        let shape = Settings::try_into_shape(res.shape).unwrap();
        Ok(IndexedShape { shape, index: res.index })
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
pub struct IndexedShapeRaw {
    pub shape: String, 
    pub index: usize
}

#[derive(Clone, Deserialize, Serialize, Debug)]
pub struct IndexedShape {
    pub shape: Shape, 
    pub index: usize
}




