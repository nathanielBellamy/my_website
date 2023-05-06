use crate::magic_square::main::Rgba;

const CACHE_CAPACITY: usize = 50;

pub struct Cache<'a> {
    idx: usize, // get & set
    pub max_idx: usize, // <= CACHE_CAPACITY
    pub vertices: [&'a [f32]; CACHE_CAPACITY],
    pub rgbas: [&'a [f32]; CACHE_CAPACITY],
}

impl<'a> Cache<'a> {
    pub fn new(max_idx: usize, vertices: [&'a [f32]; CACHE_CAPACITY], rgbas: [&'a [f32]; CACHE_CAPACITY]) -> Cache<'a> {
        let max_idx_loc: usize;
        if max_idx > CACHE_CAPACITY {
            max_idx_loc = CACHE_CAPACITY;
        } else {
            max_idx_loc = max_idx;
        }

        Cache {
            idx: 0,
            max_idx: max_idx_loc,
            vertices,
            rgbas    
        }
     
    }

    pub fn current(&self) -> (&'a [f32], &'a [f32]) {
        (self.vertices[self.idx], self.rgbas[self.idx])
    }

    pub fn set_next(&mut self, vertices: &'a [f32], rgbas: &'a [f32; 50]) {
        self.vertices[self.idx] = vertices;
        self.rgbas[self.idx] = rgbas;
        self.idx = (self.idx + 1) % self.max_idx;
    }
}
