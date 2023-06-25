
use std::ops::{Index, IndexMut};
use crate::magic_square::geometry::vertex_store::VertexStore;
use crate::magic_square::geometry::vertices::Vertex;

use super::vertices::VERTEX_ARRAY_SIZE;

pub const VERTEX_COUNT_CUBE: i32 = 24;

pub struct Cube {
    pub arr: [f32; VERTEX_ARRAY_SIZE],
    pub vert_count: i32,
    idx: usize
}

impl Index<usize> for Cube {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Cube {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<Cube> for Cube {
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

impl Cube {
    fn init() -> Cube {
        Cube { arr: [0.0; VERTEX_ARRAY_SIZE], idx: 0, vert_count: VERTEX_COUNT_CUBE }
    }
    // write to vertices
    // return array to be cached 
    pub fn f32_array() -> [f32; VERTEX_ARRAY_SIZE] {
        let x: f32 = 1.0;
        let y: f32 = 1.0;
        let z: f32 = 1.0;
        
        let mut cube = Cube::init();

        // draw cube boundary
        //
        // BACK SQUARE
        // start back north west corner
        // end back north east corner
        cube.set_next(Vertex::new(-x, y, -z));
        cube.set_next(Vertex::new(x, y, -z));
 
        // start back north east corner
        // end back south east corner
        cube.set_next(Vertex::new(x, y, -z));
        cube.set_next(Vertex::new(x, -y, -z));
    
        // start back south eash corner
        // end back south west corner
        cube.set_next(Vertex::new(x, -y, -z));
        cube.set_next(Vertex::new(-x, -y, -z));
        

        // start back south west corner
        // end back north west corner
        cube.set_next(Vertex::new(-x, -y, -z));
        cube.set_next(Vertex::new(-x, y, -z));

        // FRONT SQUARE
        // start back north west corner
        // end back north east corner
        cube.set_next(Vertex::new(-x, y, z));
        cube.set_next(Vertex::new(x, y, z));
 
        // start back north east corner
        // end back south east corner
        cube.set_next(Vertex::new(x, y, z));
        cube.set_next(Vertex::new(x, -y, z));
    
        // start back south eash corner
        // end back south west corner
        cube.set_next(Vertex::new(x, -y, z));
        cube.set_next(Vertex::new(-x, -y, z));
        

        // start back south west corner
        // end back north west corner
        cube.set_next(Vertex::new(-x, -y, z));
        cube.set_next(Vertex::new(-x, y, z));

        // CONNECT CORNERS
        // start back north west corner
        // end front north west corner
        cube.set_next(Vertex::new(-x, y, -z));
        cube.set_next(Vertex::new(-x, y, z));

        // start back south west corner
        // end front south west corner
        cube.set_next(Vertex::new(-x, -y, -z));
        cube.set_next(Vertex::new(-x, -y, z));

        // start back south east corner
        // end front south east corner
        cube.set_next(Vertex::new(x, -y, -z));
        cube.set_next(Vertex::new(x, -y, z));

        // start back north east corner
        // end front north east corner
        cube.set_next(Vertex::new(x, y, -z));
        cube.set_next(Vertex::new(x, y, z));
        
        cube.arr
    }
}



