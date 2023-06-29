use super::geometry::cache::{Cache, CACHE_CAPACITY};
use super::main::Rgba;
use super::settings::{Colors, IOGradient, IOPresetAction};
use crate::magic_square::main::log;
use crate::magic_square::settings::{Settings, Validate};
use crate::magic_square::ui_manifest::{
    INPUT_COLORS, INPUT_COLOR_DIRECTION, INPUT_COLOR_SPEED, INPUT_DRAW_PATTERN_COUNT,
    INPUT_DRAW_PATTERN_OFFSET, INPUT_DRAW_PATTERN_SPEED, INPUT_DRAW_PATTERN_TYPE,
    INPUT_LFO_1_ACTIVE, INPUT_LFO_1_AMP, INPUT_LFO_1_DEST, INPUT_LFO_1_FREQ, INPUT_LFO_1_PHASE,
    INPUT_LFO_1_SHAPE, INPUT_LFO_2_ACTIVE, INPUT_LFO_2_AMP, INPUT_LFO_2_DEST, INPUT_LFO_2_FREQ,
    INPUT_LFO_2_PHASE, INPUT_LFO_2_SHAPE, INPUT_LFO_3_ACTIVE, INPUT_LFO_3_AMP, INPUT_LFO_3_DEST,
    INPUT_LFO_3_FREQ, INPUT_LFO_3_PHASE, INPUT_LFO_3_SHAPE, INPUT_LFO_4_ACTIVE, INPUT_LFO_4_AMP,
    INPUT_LFO_4_DEST, INPUT_LFO_4_FREQ, INPUT_LFO_4_PHASE, INPUT_LFO_4_SHAPE, INPUT_MOUSE_TRACKING,
    INPUT_PRESET, INPUT_RADIUS_BASE, INPUT_RADIUS_STEP, INPUT_SHAPES, INPUT_TRANSFORM_ORDER,
    INPUT_TRANSLATION_X_BASE, INPUT_TRANSLATION_X_SPREAD, INPUT_TRANSLATION_Y_BASE,
    INPUT_TRANSLATION_Y_SPREAD, INPUT_TRANSLATION_Z_BASE, INPUT_TRANSLATION_Z_SPREAD,
    INPUT_X_AXIS_X_ROT_COEFF, INPUT_X_AXIS_Y_ROT_COEFF, INPUT_X_AXIS_Z_ROT_COEFF, INPUT_X_ROT_BASE,
    INPUT_X_ROT_SPREAD, INPUT_Y_AXIS_X_ROT_COEFF, INPUT_Y_AXIS_Y_ROT_COEFF,
    INPUT_Y_AXIS_Z_ROT_COEFF, INPUT_Y_ROT_BASE, INPUT_Y_ROT_SPREAD, INPUT_Z_ROT_BASE,
    INPUT_Z_ROT_SPREAD,
};
use crate::JsValue;

pub const EMPTY_COLORS: Colors = [[0.0; 4]; CACHE_CAPACITY];
pub const PRESET_CAPACITY: usize = 64;
// pub const PRESETS_DEFAULT: [Settings; PRESET_CAPACITY] = [Settings::new(); PRESET_CAPACITY];

#[derive(Clone, Copy, Debug)]
pub struct UiBuffer {
    pub settings: Settings,
    pub presets: [Settings; PRESET_CAPACITY],
}

impl Default for UiBuffer {
    fn default() -> UiBuffer {
        UiBuffer {
            settings: Settings::new(),
            // TODO: default presets
            presets: [Settings::new(); PRESET_CAPACITY],
        }
    }
}

impl UiBuffer {
    pub fn new() -> UiBuffer {
        UiBuffer {
            settings: Settings::new(),
            presets: [Settings::new(); PRESET_CAPACITY],
        }
    }

    pub fn copy(&self) -> UiBuffer {
        UiBuffer {
            settings: Settings { ..self.settings },
            presets: self.presets,
        }
    }

    pub fn from(settings: JsValue, presets: JsValue) -> UiBuffer {
        // log(&format!("{:?}", prev_settings));
        let settings: Settings = match serde_wasm_bindgen::from_value(settings) {
            Ok(res) => {
                log("SUCCESSFUL SETTINGS PARSE");
                // log(&format!("{:?}", res));
                res
            }
            Err(e) => {
                log(&format!("{:?}", e));
                Settings::new()
            }
        };

        let presets_vec: Vec<Settings> = match serde_wasm_bindgen::from_value(presets) {
            Ok(res) => {
                log("SUCCESSFUL PRESETS PARSE");
                res
            }
            Err(e) => {
                log(&format!("{:?}", e));
                [Settings::new(); PRESET_CAPACITY].to_vec()
            }
        };

        let mut presets: [Settings; PRESET_CAPACITY] = [Settings::new(); PRESET_CAPACITY];
        for (idx, p) in presets.iter_mut().enumerate() {
            *p = presets_vec[idx];
        }

        UiBuffer { settings, presets }
    }

    pub fn color_gradient_at_step(&self, step: u8, idx_a: usize, idx_b: usize) -> Rgba {
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

    pub fn set_color_gradient(&mut self, ig: IOGradient) {
        let mut idx: usize = 1;
        while ig.idx_a + idx < ig.idx_b {
            self.settings.colors[ig.idx_a + idx] =
                self.color_gradient_at_step(idx as u8, ig.idx_a, ig.idx_b);
            idx += 1
        }
    }

    pub fn update(
        &mut self,
        input_id: String,
        val: String,
        _geometry_cache: &mut Cache,
        // animation: &mut Animation,
    ) {
        // log(&input_id);
        // log(&val);
        //
        // TODO:
        //  // more try_into methods
        //  // store ranges in a json
        //  load json here
        //  use json to hydrate ui
        match input_id.as_str() {
            INPUT_COLOR_DIRECTION => {
                if let Ok(val) = Validate::try_into_color_direction(val) {
                    self.settings.color_direction = val
                }
            }
            INPUT_COLOR_SPEED => {
                if let Ok(val) = val.parse::<u8>() {
                    self.settings.color_speed = val
                }
            }
            INPUT_COLORS => {
                if let Ok(io_color) = Validate::try_into_io_color(val) {
                    self.settings.colors[io_color.idx] = io_color.rgba;
                }
            }
            INPUT_SHAPES => {
                if let Ok(io_shape) = Validate::try_into_io_shape(val) {
                    self.settings.shapes[io_shape.index] = io_shape.shape;
                }
            }
            INPUT_DRAW_PATTERN_TYPE => {
                if let Ok(draw_pattern_type) = Validate::try_into_draw_pattern_type(val) {
                    log(&format!("{:?}", draw_pattern_type));
                    self.settings.draw_pattern_type = draw_pattern_type;
                }
            }
            INPUT_DRAW_PATTERN_COUNT => {
                if let Ok(count) = val.parse::<i32>() {
                    if count > 0 && count < 17 {
                        self.settings.draw_pattern_count = count;
                    }
                }
            }
            INPUT_DRAW_PATTERN_OFFSET => {
                if let Ok(offset) = val.parse::<i32>() {
                    if offset > -1 && offset < CACHE_CAPACITY as i32 {
                        self.settings.draw_pattern_offset = offset;
                    }
                }
            }
            INPUT_DRAW_PATTERN_SPEED => {
                if let Ok(speed) = val.parse::<i32>() {
                    if speed > -1 && speed < 21 {
                        self.settings.draw_pattern_speed = speed;
                    }
                }
            }
            INPUT_MOUSE_TRACKING => {
                if let Ok(mouse_traking) = Validate::try_into_mouse_tracking(val) {
                    self.settings.mouse_tracking = mouse_traking
                }
            }
            INPUT_RADIUS_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.radius_base = val;
                }
            }
            INPUT_RADIUS_STEP => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.radius_step = val;
                }
            }
            INPUT_TRANSFORM_ORDER => {
                if let Ok(val) = Validate::try_into_transform_order(val) {
                    self.settings.transform_order = val;
                }
            }
            INPUT_PRESET => {
                if let Ok(io_preset) = Validate::try_into_io_preset(val) {
                    self.settings.preset = io_preset.preset;
                    match io_preset.action {
                        IOPresetAction::Save => {
                            self.presets[io_preset.preset] = Settings { ..self.settings };
                            let presets_string: String = js_sys::JSON::stringify(
                                &serde_wasm_bindgen::to_value(&self.presets.to_vec()).unwrap(),
                            )
                            .unwrap()
                            .into();
                            let local_storage =
                                web_sys::window().unwrap().local_storage().unwrap().unwrap();
                            local_storage
                                .set_item("magic_square_presets", &presets_string)
                                .unwrap();

                            log(&format!("{:?}", self.presets));
                        }
                        IOPresetAction::Set => {
                            self.settings = self.presets[io_preset.preset];
                        }
                    }
                }
            }
            INPUT_X_ROT_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.x_rot_base = val;
                }
            }
            INPUT_Y_ROT_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.y_rot_base = val;
                }
            }
            INPUT_Z_ROT_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.z_rot_base = val;
                }
            }
            INPUT_X_ROT_SPREAD => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.x_rot_spread = val;
                }
            }
            INPUT_Y_ROT_SPREAD => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.y_rot_spread = val;
                }
            }
            INPUT_Z_ROT_SPREAD => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.z_rot_spread = val;
                }
            }
            INPUT_X_AXIS_X_ROT_COEFF => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.x_axis_x_rot_coeff = val;
                }
            }
            INPUT_X_AXIS_Y_ROT_COEFF => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.x_axis_y_rot_coeff = val;
                }
            }
            INPUT_X_AXIS_Z_ROT_COEFF => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.x_axis_z_rot_coeff = val;
                }
            }
            INPUT_Y_AXIS_X_ROT_COEFF => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.y_axis_x_rot_coeff = val;
                }
            }
            INPUT_Y_AXIS_Y_ROT_COEFF => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.y_axis_y_rot_coeff = val;
                }
            }
            INPUT_Y_AXIS_Z_ROT_COEFF => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.y_axis_z_rot_coeff = val;
                }
            }
            INPUT_TRANSLATION_X_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_x_base = val;
                }
            }
            INPUT_TRANSLATION_X_SPREAD => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_x_spread = val;
                }
            }
            INPUT_TRANSLATION_Y_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_y_base = val;
                }
            }
            INPUT_TRANSLATION_Y_SPREAD => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_y_spread = val;
                }
            }
            INPUT_TRANSLATION_Z_BASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_z_base = val;
                }
            }
            INPUT_TRANSLATION_Z_SPREAD => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.translation_z_spread = val;
                }
            }
            INPUT_LFO_1_ACTIVE => {
                if let Ok(val) = val.parse::<bool>() {
                    self.settings.lfo_1_active = val;
                }
            }
            INPUT_LFO_1_AMP => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_1_amp = val;
                }
            }
            INPUT_LFO_1_DEST => {
                if let Ok(dest) = Validate::try_into_lfo_destination(val) {
                    self.settings.lfo_1_dest = dest;
                }
            }
            INPUT_LFO_1_FREQ => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_1_freq = val;
                }
            }
            INPUT_LFO_1_PHASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_1_phase = val;
                }
            }
            INPUT_LFO_1_SHAPE => {
                if let Ok(shape) = Validate::try_into_lfo_shape(val) {
                    self.settings.lfo_1_shape = shape;
                }
            }
            INPUT_LFO_2_ACTIVE => {
                if let Ok(val) = val.parse::<bool>() {
                    self.settings.lfo_2_active = val;
                }
            }
            INPUT_LFO_2_AMP => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_2_amp = val;
                }
            }
            INPUT_LFO_2_DEST => {
                if let Ok(dest) = Validate::try_into_lfo_destination(val) {
                    self.settings.lfo_2_dest = dest;
                }
            }
            INPUT_LFO_2_FREQ => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_2_freq = val;
                }
            }
            INPUT_LFO_2_PHASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_2_phase = val;
                }
            }
            INPUT_LFO_2_SHAPE => {
                if let Ok(shape) = Validate::try_into_lfo_shape(val) {
                    self.settings.lfo_2_shape = shape;
                }
            }
            INPUT_LFO_3_ACTIVE => {
                if let Ok(val) = val.parse::<bool>() {
                    self.settings.lfo_3_active = val;
                }
            }
            INPUT_LFO_3_AMP => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_3_amp = val;
                }
            }
            INPUT_LFO_3_DEST => {
                if let Ok(dest) = Validate::try_into_lfo_destination(val) {
                    self.settings.lfo_3_dest = dest;
                }
            }
            INPUT_LFO_3_FREQ => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_3_freq = val;
                }
            }
            INPUT_LFO_3_PHASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_3_phase = val;
                }
            }
            INPUT_LFO_3_SHAPE => {
                if let Ok(shape) = Validate::try_into_lfo_shape(val) {
                    self.settings.lfo_3_shape = shape;
                }
            }
            INPUT_LFO_4_ACTIVE => {
                if let Ok(val) = val.parse::<bool>() {
                    self.settings.lfo_4_active = val;
                }
            }
            INPUT_LFO_4_AMP => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_4_amp = val;
                }
            }
            INPUT_LFO_4_DEST => {
                if let Ok(dest) = Validate::try_into_lfo_destination(val) {
                    self.settings.lfo_4_dest = dest;
                }
            }
            INPUT_LFO_4_FREQ => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_4_freq = val;
                }
            }
            INPUT_LFO_4_PHASE => {
                if let Ok(val) = val.parse::<f32>() {
                    self.settings.lfo_4_phase = val;
                }
            }
            INPUT_LFO_4_SHAPE => {
                if let Ok(shape) = Validate::try_into_lfo_shape(val) {
                    self.settings.lfo_4_shape = shape;
                }
            }
            _ => {}
        }
    }
}
