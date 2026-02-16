use super::settings::{DrawPatternType, Settings};
use crate::magic_square::geometry::cache::CACHE_CAPACITY;
use crate::magic_square::geometry::{Shape, Shapes};
// use super::main::log;

pub type Frame = [Shape; CACHE_CAPACITY];
pub type Reel = [Frame; CACHE_CAPACITY];

const EMPTY_FRAME: Frame = [Shape::None; CACHE_CAPACITY];
const EMPTY_REEL: Reel = [[Shape::None; CACHE_CAPACITY]; CACHE_CAPACITY];

#[derive(Clone, Copy, Debug)]
pub struct Animation {
    pub reel: Reel,
    idx: usize,
}

pub fn offset_plus(offset: usize, summand: usize) -> usize {
    (offset + summand) % CACHE_CAPACITY
}

impl Animation {
    pub fn new() -> Animation {
        Animation {
            reel: EMPTY_REEL,
            idx: 0,
        }
    }

    pub fn inc(&mut self) {
        self.idx = (self.idx + 1) % CACHE_CAPACITY;
    }

    pub fn curr_shapes(&self) -> Shapes {
        self.reel[self.idx]
    }

    pub fn set_from(&mut self, settings: &Settings) {
        match settings.draw_pattern_type {
            DrawPatternType::In => self.set_in(
                settings.draw_pattern_count,
                settings.shapes,
                settings.draw_pattern_offset as usize,
            ),
            DrawPatternType::Out => self.set_out(
                settings.draw_pattern_count,
                settings.shapes,
                settings.draw_pattern_offset as usize,
            ),
            DrawPatternType::Fix => self.set_fix(
                settings.draw_pattern_count,
                settings.shapes,
                settings.draw_pattern_offset as usize,
            ),
        }
    }

    fn set_fix(&mut self, count: i32, shapes: Shapes, offset: usize) {
        for frame in self.reel.iter_mut() {
            *frame = EMPTY_FRAME;

            for i in 0..count {
                let i_u = i as usize;
                frame[offset_plus(offset, i_u)] = shapes[offset_plus(offset, i_u)]
            }
        }
    }

    fn set_out(&mut self, count: i32, shapes: Shapes, offset: usize) {
        for (i, frame) in self.reel.iter_mut().enumerate() {
            *frame = EMPTY_FRAME;
            for c in 0..count {
                let c_u = c as usize;
                frame[offset_plus(i + offset, c_u)] = shapes[offset_plus(i + offset, c_u)];
            }
        }
    }

    fn set_in(&mut self, count: i32, shapes: Shapes, offset: usize) {
        for (i, frame) in self.reel.iter_mut().enumerate() {
            *frame = EMPTY_FRAME;
            for c in 0..count {
                let c_u = c as usize;
                frame[(CACHE_CAPACITY - offset_plus(i + offset, c_u)) % CACHE_CAPACITY] =
                    shapes[offset_plus(i + offset, c_u)];
            }
        }
    }
}
