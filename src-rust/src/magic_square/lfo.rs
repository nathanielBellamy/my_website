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
    Sawtooth,
    #[default]
    Sine,
    Square,
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
       self.amp * ((self.freq * x) + self.phase).sin()
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
