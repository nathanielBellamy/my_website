use crate::magic_square::geometry::{Shape, Shapes};
use super::hexagon::VERTICES_HEXAGON;

pub const CACHE_CAPACITY: usize = 16;
pub const CACHE_SHAPE_WIDTH: usize = 300;
pub const CACHE_VERTICES_LEN: usize = CACHE_CAPACITY * CACHE_SHAPE_WIDTH;

pub type CacheVertices = [f32; CACHE_VERTICES_LEN];

#[derive(Clone, Copy, Debug)]
pub struct Cache {
    pub idx: usize, // get & set
    pub shapes: Shapes,
}

impl Cache {
    pub fn f32_array() -> [f32; 42] {
        VERTICES_HEXAGON
    }

    pub fn new(shapes: &Shapes) -> Cache {
        Cache {
            idx: 0,
            shapes: [Shape::None; CACHE_CAPACITY]
        }
    }

    pub fn cache_idx_to_vert_idx(&self, cache_idx: usize) -> usize {
        cache_idx * CACHE_SHAPE_WIDTH
    }
}
