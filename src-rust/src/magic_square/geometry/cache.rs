use crate::magic_square::main::Rgba;

pub const CACHE_CAPACITY: usize = 50;

//
//  Slices were a nice idea but they won't let us cache.
//  The computation takes place in the event listener closure
//  Using slices, this means that a new slice set to the cache
//  will immediately be an invalid reference at the end of the closure.
//  We need to pass the actual arrays out of the closure into the cache
//
//  But in order to do this, all of the arrays need to be the same length
//  So we need to:
//      1. determine what a suitable array capacity that can fit all of the geometry we want
//      2. provide a method to pass the appropriate length slice to GL for each render call
//
//
//
pub struct Cache<'a> {
    pub idx: usize, // get & set
    pub max_idx: usize, // <= CACHE_CAPACITY
    pub vertices: [&'a [f32]; CACHE_CAPACITY],
    pub rgbas: [&'a [f32; 4]; CACHE_CAPACITY],
}

impl<'a> Cache<'a> {
    pub fn new(max_idx: usize, vertices: [&'a [f32]; CACHE_CAPACITY], rgbas: [&'a Rgba; CACHE_CAPACITY]) -> Cache<'a> {
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

    pub fn set_next(&mut self, vertices: &'a [f32], rgbas: &'a Rgba) {
        self.vertices[self.idx] = vertices;
        self.rgbas[self.idx] = rgbas;
        self.idx = (self.idx + 1) % self.max_idx;
    }
}
