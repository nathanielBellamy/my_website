use crate::magic_square::vertices::{VertexArr, VERTEX_ARRAY_SIZE};
use crate::magic_square::geometry::Shape;

pub const CACHE_CAPACITY: usize = 16;

#[derive(Clone, Copy, Debug)]
pub struct Cache {
    pub idx: usize, // get & set
    pub max_idx: usize, // <= CACHE_CAPACITY
    pub vertices: [VertexArr; CACHE_CAPACITY], // TODO: refactor into an array of Struc vertices,rgba, shape
    pub shapes: [Shape; CACHE_CAPACITY],
}

impl Cache {
    pub fn new(
        max_idx: usize, 
        vertices: [VertexArr; CACHE_CAPACITY], 
        shapes: [Shape; CACHE_CAPACITY]
    ) -> Cache {
        let max_idx_loc: usize;
        if max_idx > CACHE_CAPACITY - 1 {
            max_idx_loc = CACHE_CAPACITY - 1;
        } else {
            max_idx_loc = max_idx;
        }

        Cache {
            idx: 0,
            max_idx: max_idx_loc,
            vertices,
            shapes
        }
    }

    pub fn clear(&mut self) {
        for arr in self.vertices.iter_mut() {
            *arr = [0.0; VERTEX_ARRAY_SIZE];
        }

        for shape in self.shapes.iter_mut() {
            *shape = Shape::None;
        }

        self.idx = 0;
        self.max_idx = 7;
    }

    pub fn set_next(&mut self, vertices: VertexArr, shape: Shape) {
        self.vertices[self.idx] = vertices;
        self.shapes[self.idx] = shape;
        self.idx = (self.idx + 1) % CACHE_CAPACITY;
    }

    pub fn gl_vertices(&self, idx: usize) -> &[f32] {
        if idx > CACHE_CAPACITY - 1 {
            return &self.vertices[0][0..1];
        }
        &self.vertices[idx][0..Cache::gl_vert_len_from_shape(self.shapes[idx])]
    }

    pub fn gl_vert_len_from_shape(shape: Shape) -> usize {
        match shape {
            Shape::Hexagon => 42,
            Shape::Icosahedron => 300,
            Shape::None => 0,
            _ => 0
        }
    }
}
