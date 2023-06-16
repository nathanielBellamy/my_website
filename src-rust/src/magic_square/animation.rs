use crate::magic_square::geometry::Shape;
use crate::magic_square::geometry::cache::CACHE_CAPACITY;
use super::settings::{DrawPattern, Settings};

pub type Shapes = [Shape; CACHE_CAPACITY];
pub type Frame = [Shape; CACHE_CAPACITY];
pub type Reel = [Frame; CACHE_CAPACITY];

const EMPTY_FRAME: Frame = [Shape::None; CACHE_CAPACITY];
const EMPTY_REEL: Reel = [[Shape::None; CACHE_CAPACITY]; CACHE_CAPACITY];

#[derive(Clone, Copy, Debug)]
pub struct Animation {
    pub draw_pattern: DrawPattern,
    pub shapes: Shapes,
    pub reel: Reel,
    pub offset: usize, // 0 <= offset <= 15
    pub idx: usize,
}

pub fn offset_plus(offset: usize, summand: usize) -> usize {
    (offset + summand) % CACHE_CAPACITY
}

impl Animation {
    pub fn new(draw_pattern: DrawPattern, shapes: Shapes) -> Animation {
        let mut anim = Animation {
            draw_pattern,
            shapes,
            reel: EMPTY_REEL,
            offset: 0,
            idx: 0
        };
        anim.set_reel();
        anim
    }

    pub fn new_from_settings(settings: &Settings) -> Animation {
        Animation::new(settings.draw_pattern, settings.shapes)
    }

    pub fn next(&mut self) {
        self.idx = (self.idx + 1) % CACHE_CAPACITY;
    }

    pub fn set_reel(&mut self) {
        match self.draw_pattern {
            DrawPattern::Fix1 => {
               for frame in self.reel.iter_mut() {
                    *frame = EMPTY_FRAME;
                    frame[self.offset] = self.shapes[self.offset];
               }
            },
            DrawPattern::Fix2 => {
               for frame in self.reel.iter_mut() {
                    *frame = EMPTY_FRAME;
                    frame[self.offset] = self.shapes[self.offset];
                    frame[offset_plus(self.offset, 1)] = self.shapes[offset_plus(self.offset, 1)];
               }
            },
            DrawPattern::Fix3 => {
               for frame in self.reel.iter_mut() {
                    *frame = EMPTY_FRAME;
                    frame[self.offset] = self.shapes[self.offset];
                    frame[offset_plus(self.offset, 1)] = self.shapes[offset_plus(self.offset, 1)];
                    frame[offset_plus(self.offset, 2)] = self.shapes[offset_plus(self.offset, 2)];
               }
            },
            DrawPattern::Fix4 => {
               for frame in self.reel.iter_mut() {
                    *frame = EMPTY_FRAME;
                    frame[self.offset] = self.shapes[self.offset];
                    frame[offset_plus(self.offset, 1)] = self.shapes[offset_plus(self.offset, 1)];
                    frame[offset_plus(self.offset, 2)] = self.shapes[offset_plus(self.offset, 2)];
                    frame[offset_plus(self.offset, 3)] = self.shapes[offset_plus(self.offset, 3)];
               }
            },
            DrawPattern::Fix5 => {
               for frame in self.reel.iter_mut() {
                    *frame = EMPTY_FRAME;
                    frame[self.offset] = self.shapes[self.offset];
                    frame[offset_plus(self.offset, 1)] = self.shapes[offset_plus(self.offset, 1)];
                    frame[offset_plus(self.offset, 2)] = self.shapes[offset_plus(self.offset, 2)];
                    frame[offset_plus(self.offset, 3)] = self.shapes[offset_plus(self.offset, 3)];
                    frame[offset_plus(self.offset, 4)] = self.shapes[offset_plus(self.offset, 4)];
               }
            },
            DrawPattern::Fix6 => {
               for frame in self.reel.iter_mut() {
                    *frame = EMPTY_FRAME;
                    frame[self.offset] = self.shapes[self.offset];
                    frame[offset_plus(self.offset, 1)] = self.shapes[offset_plus(self.offset, 1)];
                    frame[offset_plus(self.offset, 2)] = self.shapes[offset_plus(self.offset, 2)];
                    frame[offset_plus(self.offset, 3)] = self.shapes[offset_plus(self.offset, 3)];
                    frame[offset_plus(self.offset, 4)] = self.shapes[offset_plus(self.offset, 4)];
                    frame[offset_plus(self.offset, 5)] = self.shapes[offset_plus(self.offset, 5)];
               }
            },
            DrawPattern::Fix7 => {
               for frame in self.reel.iter_mut() {
                    *frame = EMPTY_FRAME;
                    frame[self.offset] = self.shapes[self.offset];
                    frame[offset_plus(self.offset, 1)] = self.shapes[offset_plus(self.offset, 1)];
                    frame[offset_plus(self.offset, 2)] = self.shapes[offset_plus(self.offset, 2)];
                    frame[offset_plus(self.offset, 3)] = self.shapes[offset_plus(self.offset, 3)];
                    frame[offset_plus(self.offset, 4)] = self.shapes[offset_plus(self.offset, 4)];
                    frame[offset_plus(self.offset, 5)] = self.shapes[offset_plus(self.offset, 5)];
                    frame[offset_plus(self.offset, 6)] = self.shapes[offset_plus(self.offset, 6)];
               }
            },
            DrawPattern::Fix8 => {
               for frame in self.reel.iter_mut() {
                    *frame = EMPTY_FRAME;
                    frame[self.offset] = self.shapes[self.offset];
                    frame[offset_plus(self.offset, 1)] = self.shapes[offset_plus(self.offset, 1)];
                    frame[offset_plus(self.offset, 2)] = self.shapes[offset_plus(self.offset, 2)];
                    frame[offset_plus(self.offset, 3)] = self.shapes[offset_plus(self.offset, 3)];
                    frame[offset_plus(self.offset, 4)] = self.shapes[offset_plus(self.offset, 4)];
                    frame[offset_plus(self.offset, 5)] = self.shapes[offset_plus(self.offset, 5)];
                    frame[offset_plus(self.offset, 6)] = self.shapes[offset_plus(self.offset, 6)];
                    frame[offset_plus(self.offset, 7)] = self.shapes[offset_plus(self.offset, 7)];
               }
            },

            // OUT
            DrawPattern::Out1 => {
                for (i, frame) in self.reel.iter_mut().enumerate() {
                    *frame = EMPTY_FRAME;
                    frame[i] = self.shapes[self.offset];
                }
            }
            DrawPattern::Out2 => {
                for (i, frame) in self.reel.iter_mut().enumerate() {
                    *frame = EMPTY_FRAME;
                    frame[i] = self.shapes[self.offset];
                    frame[offset_plus(i, 1)] = self.shapes[offset_plus(i, 1)];
                }
            }
            _ => self.reel = EMPTY_REEL
        }
    }
}
