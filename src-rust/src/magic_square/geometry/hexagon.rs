use std::ops::{Index, IndexMut};
use crate::magic_square::geometry::vertex_store::VertexStore;
use crate::magic_square::geometry::vertices::Vertex;

use super::vertices::VERTEX_ARRAY_SIZE;

pub const VERTEX_COUNT_HEXAGON: i32 = 14;

pub struct Hexagon {
    pub arr: [f32; VERTEX_ARRAY_SIZE],
    pub vert_count: i32,
    idx: usize
}

impl Index<usize> for Hexagon {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Hexagon {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<Hexagon> for Hexagon {
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

impl Hexagon {
    fn init() -> Hexagon {
        Hexagon { arr: [0.0; VERTEX_ARRAY_SIZE], idx: 0, vert_count: VERTEX_COUNT_HEXAGON }
    }
    // write to vertices
    // return array to be cached 
    pub fn f32_array() -> [f32; VERTEX_ARRAY_SIZE] {
        let xy: f32 = 0.0;
        let x_shift: f32 = 0.5; // cos(pi/3)
        let y_shift: f32 = 0.86602540378; // sin(pi/3)
        
        let mut hexagon = Hexagon::init();

        // draw hexagon boundary
        // start north east corner
        // end east corner
        hexagon.set_next(Vertex::new(x_shift, y_shift, xy));
        hexagon.set_next(Vertex::new(1.0, 0.0, xy));

        // start east corner
        // end south east corner
        hexagon.set_next(Vertex::new(1.0, 0.0, xy));
        hexagon.set_next(Vertex::new(x_shift, -y_shift, xy));

        // start east corner
        // end south east corner
        hexagon.set_next(Vertex::new(1.0, 0.0, xy));
        hexagon.set_next(Vertex::new(x_shift, -y_shift, xy));

        // start south east corner
        // end south west corner
        hexagon.set_next(Vertex::new(x_shift, -y_shift, xy));
        hexagon.set_next(Vertex::new(-x_shift, -y_shift, xy));

        // start south west corner
        // end west corner
        hexagon.set_next(Vertex::new(-x_shift, -y_shift, xy));
        hexagon.set_next(Vertex::new(-1.0, 0.0, xy));

        // start west corner
        // end north west corner
        hexagon.set_next(Vertex::new(-1.0, 0.0, xy));
        hexagon.set_next(Vertex::new(-x_shift, y_shift, xy));


        // start north west corner
        // end north east corner
        hexagon.set_next(Vertex::new(-x_shift, y_shift, xy));
        hexagon.set_next(Vertex::new(x_shift, y_shift, xy));
        
        hexagon.arr
    }
}



