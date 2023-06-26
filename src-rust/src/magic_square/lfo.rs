use serde::{Deserialize, Serialize};
use super::settings::Settings;
// use crate::magic_square::main::log;

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub enum LfoDestination {
    // rotation
    PitchBase,
    PitchSpread,
    PitchX,
    PitchY,
    RollBase,
    RollSpread,
    RollX,
    RollY,
    YawBase,
    YawSpread,
    YawX,
    YawY,

    // radius
    RadiusBase,
    RadiusStep,

    // translation
    TranslationXBase,
    TranslationXSpread,
    TranslationYBase,
    TranslationYSpread,
    #[default]
    None
}

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub enum LfoShape {
    Linear,
    // Sawtooth,
    #[default]
    Sine,
    // Square
}

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub struct Lfo {
    pub active: bool,
    pub amp: f32,
    pub dest: LfoDestination,
    pub freq: f32,
    pub phase: f32,
    pub shape: LfoShape
}

impl Lfo {
    pub fn new(active: bool, amp: f32, dest: LfoDestination, freq: f32, phase: f32, shape: LfoShape) -> Lfo {
        Lfo {
            active, amp, dest, freq, phase, shape
        }
    }

    pub fn eval(&self, x: f32) -> f32 {
        // log(&format!("{:?}", self.shape));
        match self.shape {
            LfoShape::Sine => self.amp * (((self.freq / 2.0)  * x) + self.phase).sin(),
            LfoShape::Linear => {
                // | LfoShape::Square
                // | LfoShape::Sawtooth => {
                let mut result: f32 = 0.0;
                // determine the sub-intervals
                let subinterval_count: usize = (self.freq * 6.28).floor() as usize;
                let subinterval_length: f32 = 6.28 / (subinterval_count as f32);
                
                // determine in which sub-interval x falls
                // and whether that interval has positive or negative slope
                // by construction, the left-most interval will have positive slope
                let mut x_subinterval_parity: bool = true; // true = pos, false = neg
                
                for idx in 0..subinterval_count {
                    let x_left = -3.14 + ((idx as f32) * subinterval_length);
                    let x_right = -3.14 + ((idx + 1) as f32 * subinterval_length);
                    if x >= x_left && x < x_right  {// [x_left, x_right), half-open interval
                        // x is in sub-interval
                        result = match self.shape {
                                LfoShape::Linear => {
                                    let par: f32 = if x_subinterval_parity { -1.0 } else { 1.0 };
                                    (par * (2.0 * self.amp) / subinterval_length) * (x - x_left) - par * self.amp
                                },
                                _ => 1.0
                            };
                        break;
                    } else {
                        // move on to next interval
                        // record parity flip
                        x_subinterval_parity = !x_subinterval_parity;
                    }
                }
                // log(&format!("{result}"));
                return result;
            }
        }
    }

    pub fn modify(&self, t: f32, settings: &mut Settings) {
        if self.active {
            match self.dest {
                LfoDestination::PitchBase => {
                    settings.y_rot_base = settings.y_rot_base 
                                                        + self.eval(t) * 3.14
                },
                LfoDestination::PitchSpread => {
                    settings.y_rot_spread = settings.y_rot_spread
                                                        + self.eval(t) * 0.3
                },
                LfoDestination::PitchX => {
                    settings.x_axis_y_rot_coeff = settings.x_axis_y_rot_coeff
                                                        + self.eval(t)
                },
                LfoDestination::PitchY => {
                    settings.y_axis_y_rot_coeff = settings.y_axis_y_rot_coeff
                                                        + self.eval(t)
                },
                LfoDestination::RollBase => {
                    settings.x_rot_base = settings.x_rot_base 
                                                        + self.eval(t) * 3.14
                },
                LfoDestination::RollSpread => {
                    settings.x_rot_spread = settings.x_rot_spread
                                                        + self.eval(t) * 0.3
                },
                LfoDestination::RollX => {
                    settings.x_axis_x_rot_coeff = settings.x_axis_x_rot_coeff
                                                                + self.eval(t)
                },
                LfoDestination::RollY => {
                    settings.y_axis_x_rot_coeff = settings.y_axis_x_rot_coeff
                                                                + self.eval(t)
                },
                LfoDestination::YawBase => {
                    settings.z_rot_base = settings.z_rot_base 
                                                        + self.eval(t) * 3.14
                },
                LfoDestination::YawSpread => {
                    settings.z_rot_spread = settings.z_rot_spread
                                                        + self.eval(t) * 0.3
                },
                LfoDestination::YawX => {
                    settings.x_axis_z_rot_coeff = settings.x_axis_z_rot_coeff
                                                                + self.eval(t)
                },
                LfoDestination::YawY => {
                    settings.y_axis_z_rot_coeff = settings.y_axis_z_rot_coeff
                                                                + self.eval(t)
                },
                LfoDestination::RadiusBase => {
                    settings.radius_base = settings.radius_base
                                                        + self.eval(t)
                },
                LfoDestination::RadiusStep => {
                    settings.radius_step = settings.radius_step
                                                        + self.eval(t)
                },
                LfoDestination::TranslationXBase => {
                    settings.translation_x_base =  settings.translation_x_base 
                                                                + self.eval(t);
                },
                LfoDestination::TranslationXSpread => {
                    settings.translation_x_spread =  settings.translation_x_spread 
                                                                    + self.eval(t);
                },
                LfoDestination::TranslationYBase => {
                    settings.translation_y_base =  settings.translation_y_base 
                                                                + self.eval(t);
                },
                LfoDestination::TranslationYSpread => {
                    settings.translation_y_spread =  settings.translation_y_spread + self.eval(t);
                },
                LfoDestination::None => {}
            }
        }
    }
}
