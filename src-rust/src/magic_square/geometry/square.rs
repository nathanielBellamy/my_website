use std::ops::{Index, IndexMut};
use crate::magic_square::geometry::vertex_store::VertexStore;
use crate::magic_square::geometry::vertices::Vertex;

use super::vertices::VERTEX_ARRAY_SIZE;

pub const VERTEX_COUNT_SQUARE: i32 = 8;

pub struct Square {
    pub arr: [f32; VERTEX_ARRAY_SIZE],
    pub vert_count: i32,
    idx: usize
}

impl Index<usize> for Square {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Square {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<Square> for Square {
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

impl Square {
    fn init() -> Square {
        Square { arr: [0.0; VERTEX_ARRAY_SIZE], idx: 0, vert_count: VERTEX_COUNT_SQUARE }
    }
    // write to vertices
    // return array to be cached 
    pub fn f32_array() -> [f32; VERTEX_ARRAY_SIZE] {
        let x: f32 = 1.0;
        let y: f32 = 1.0;
        let z: f32 = 0.0;
        
        let mut square = Square::init();

        // draw square boundary
        // end back north east corner
        square.set_next(Vertex::new(-x, y, z));
        square.set_next(Vertex::new(x, y, z));
 
        // start back north east corner
        // end back south east corner
        square.set_next(Vertex::new(x, y, z));
        square.set_next(Vertex::new(x, -y, z));
    
        // start back south eash corner
        // end back south west corner
        square.set_next(Vertex::new(x, -y, z));
        square.set_next(Vertex::new(-x, -y, z));
        
        // start back south west corner
        // end back north west corner
        square.set_next(Vertex::new(-x, -y, z));
        square.set_next(Vertex::new(-x, y, z));
        
        square.arr
    }
}



