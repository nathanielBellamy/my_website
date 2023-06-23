use crate::magic_square::geometry::{Geometry, Shape};

pub const CACHE_CAPACITY: usize = 16;
pub const CACHE_SHAPE_WIDTH: usize = 300;
pub const CACHE_VERTICES_LEN: usize = CACHE_CAPACITY * CACHE_SHAPE_WIDTH;

pub type CacheVertices = [f32; CACHE_VERTICES_LEN];
pub type CacheShapes = [Shape; CACHE_CAPACITY];

#[derive(Clone, Copy, Debug)]
pub struct Cache {
    pub idx: usize, // get & set
    pub vertices: CacheVertices, // TODO: refactor into an array of Struc vertices,rgba, shape
    pub shapes: CacheShapes,
}

impl Cache {
    pub fn new(
        shapes: &CacheShapes
    ) -> Cache {
        let mut cache = Cache {
            idx: 0,
            vertices: [0.0; CACHE_VERTICES_LEN],
            shapes: [Shape::None; CACHE_CAPACITY]
        };
        cache.update_from_shapes(shapes); 
        cache
    }

    pub fn cache_idx_to_vert_idx(&self, cache_idx: usize) -> usize {
        cache_idx * CACHE_SHAPE_WIDTH
    }

    pub fn set_shape_in_slice(&mut self, cache_idx: usize) {
        let vert_l_idx = self.cache_idx_to_vert_idx(cache_idx);
        let vert_r_idx = self.cache_idx_to_vert_idx(cache_idx + 1);
        let new_shape = self.shapes[cache_idx];
        let new_shape_vertices: [f32; 300] = Geometry { shape: new_shape }.arr(); 

        for (idx, val) in self.vertices[vert_l_idx..vert_r_idx].iter_mut().enumerate() {
            *val = new_shape_vertices[idx];
        }
    }

    pub fn update_from_shapes(&mut self, new_shapes: &CacheShapes) {
        let mut indices_to_update: Vec<usize> = vec![];
        for (idx, _shape) in self.shapes.iter().enumerate() {
            if self.shapes[idx] != new_shapes[idx] {
                indices_to_update.push(idx)
            }
        }
        self.shapes = *new_shapes;
        for cache_idx in indices_to_update.iter() {
            self.set_shape_in_slice(*cache_idx);
            
        }

    }

    // pub fn gl_vertices(&self, idx: usize) -> &[f32] {
    //     if idx > CACHE_CAPACITY - 1 {
    //         return &self.vertices[0][0..1];
    //     }
    //     &self.vertices[idx][0..Cache::gl_vert_len_from_shape(self.shapes[idx])]
    // }

    // pub fn gl_vert_len_from_shape(shape: Shape) -> usize {
    //     match shape {
    //         Shape::Hexagon => 42,
    //         Shape::Icosahedron => 300,
    //         Shape::None => 0,
    //         _ => 0
    //     }
    // }
}
