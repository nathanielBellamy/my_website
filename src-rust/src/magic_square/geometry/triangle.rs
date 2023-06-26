
use std::ops::{Index, IndexMut};
use crate::magic_square::geometry::vertex_store::VertexStore;
use crate::magic_square::geometry::vertices::Vertex;

use super::vertices::VERTEX_ARRAY_SIZE;

pub const VERTEX_COUNT_TRIANGLE: i32 = 6;

pub struct Triangle {
    pub arr: [f32; VERTEX_ARRAY_SIZE],
    pub vert_count: i32,
    idx: usize
}

impl Index<usize> for Triangle {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Triangle {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<Triangle> for Triangle {
    fn idx(&self) -> usize {
        self.idx
    }

    fn set_idx(&mut self, new_idx: usize) -> usize {
        self.idx = new_idx;
        self.idx
    }

    fn arr(&mut self) -> &mut [f32] {
        &mut self.arr
    }
}

impl Triangle {
    fn init() -> Triangle {
        Triangle { arr: [0.0; VERTEX_ARRAY_SIZE], idx: 0, vert_count: VERTEX_COUNT_TRIANGLE }
    }
    // write to vertices
    // return array to be cached 
    pub fn f32_array() -> [f32; VERTEX_ARRAY_SIZE] {
        let x: f32 = 1.1547; // 2/tan(pi/3)
        let y: f32 = 1.0;
        let z: f32 = 0.0;
        
        let mut triangle = Triangle::init();

        // draw triangle boundary
        triangle.set_next(Vertex::new(0.0, y, z));
        triangle.set_next(Vertex::new(-x, -y, z));
 
        triangle.set_next(Vertex::new(-x, -y, z));
        triangle.set_next(Vertex::new(x, -y, z));
    
        triangle.set_next(Vertex::new(x, -y, z));
        triangle.set_next(Vertex::new(0.0, y, z));
        
        triangle.arr
    }
}



