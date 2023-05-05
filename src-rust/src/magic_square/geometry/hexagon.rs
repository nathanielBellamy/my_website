
use std::ops::{Index, IndexMut};
use crate::magic_square::traits::VertexStore;
use crate::magic_square::vertices::{Vertex, Vertices};
use crate::magic_square::transformations::RotationSequence;

const HEXAGON_ARR_LEN: usize = 42;

pub struct Hexagon {
    pub arr: [f32; HEXAGON_ARR_LEN], // # coordinates needed to define hexagon
    idx: usize
}

impl Hexagon {
    fn init() -> Hexagon {
        Hexagon { arr: [0.0; HEXAGON_ARR_LEN], idx: 0 }
    }
    // write to vertices
    // return array to be cached 
    pub fn new(buffer: &[f32; 2], radius: f32, rotation: RotationSequence) -> Hexagon {
        let center_x = buffer[0];
        let center_y = buffer[1];
        let xy = 0.02 * center_x * center_y;

        let x_shift = radius * 0.5; // r cos(pi/3)
        let y_shift = radius * 0.86602540378; // r sin(pi/3)
        
        let mut hexagon = Hexagon::init();

        // draw hexagon boundary
        // start north east corner
        // end east corner
        hexagon.set_next(
            Vertex::new(center_x + x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );
        hexagon.set_next(
            Vertex::new(center_x + radius, center_y, xy)
                .rot(rotation)
        );

        // start east corner
        // end south east corner
        hexagon.set_next(
            Vertex::new(center_x + radius, center_y, xy)
                .rot(rotation)
        );
        hexagon.set_next(
            Vertex::new(center_x + x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );

        // start east corner
        // end south east corner
        hexagon.set_next(
            Vertex::new(center_x + radius, center_y, xy)
                .rot(rotation)
        );
        hexagon.set_next(
            Vertex::new(center_x + x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );

        // start south east corner
        // end south west corner
        hexagon.set_next(
            Vertex::new(center_x + x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );
        hexagon.set_next(
            Vertex::new(center_x - x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );

        // start south west corner
        // end west corner
        hexagon.set_next(
            Vertex::new(center_x - x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );
        hexagon.set_next(
            Vertex::new(center_x - radius, center_y, xy)
                .rot(rotation)
        );

        // start west corner
        // end north west corner
        hexagon.set_next(
            Vertex::new(center_x - radius, center_y, xy)
                .rot(rotation)
        );
        hexagon.set_next(
            Vertex::new(center_x - x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );


        // start north west corner
        // end north east corner
        hexagon.set_next( 
            Vertex::new(center_x - x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );
        hexagon.set_next(
            Vertex::new(center_x + x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );
        
        hexagon
    }
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
