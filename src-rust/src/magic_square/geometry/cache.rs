use crate::magic_square::geometry::Shapes;

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
    pub fn new(shapes: &Shapes) -> Cache {
        Cache {
            idx: 0,
            shapes: *shapes,
        }
    }

    pub fn cache_idx_to_vert_idx(&self, cache_idx: usize) -> usize {
        cache_idx * CACHE_SHAPE_WIDTH
    }
}
