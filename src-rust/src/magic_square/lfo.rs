use serde::{Deserialize, Serialize};
use crate::magic_square::ui_buffer::UiBuffer;
// use crate::magic_square::main::log;

#[derive(Serialize, Deserialize, Clone, Copy, Default, Debug)]
pub enum LfoDestination {
    TranslationX,
    TranslationY,
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
            LfoShape::Sine => self.amp * ((self.freq * x) + self.phase).sin(),
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

    pub fn modify(&self, t: f32, ui_buffer: &mut UiBuffer) {
        if self.active {
            match self.dest {
                LfoDestination::TranslationX => {
                    ui_buffer.settings.translation_x =  ui_buffer.settings.translation_x + self.eval(t);
                },
                LfoDestination::TranslationY => {
                    ui_buffer.settings.translation_y =  ui_buffer.settings.translation_y + self.eval(t);
                },
                LfoDestination::None => {}
            }
        }
    }
}
