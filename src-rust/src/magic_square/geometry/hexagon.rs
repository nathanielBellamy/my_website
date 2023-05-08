
use std::sync::Arc;
use std::ops::{Index, IndexMut};
use crate::magic_square::traits::VertexStore;
use crate::magic_square::vertices::{Vertex, VERTEX_ARRAY_SIZE};
use crate::magic_square::transformations::{RotationSequence, Translation};

pub struct Hexagon {
    pub arr: [f32; VERTEX_ARRAY_SIZE], // # coordinates needed to define hexagon
    idx: usize
}

impl Hexagon {
    fn init() -> Hexagon {
        Hexagon { arr: [0.0; VERTEX_ARRAY_SIZE], idx: 0 }
    }
    // write to vertices
    // return array to be cached 
    pub fn new(
        radius: f32, 
        rotation: RotationSequence,
        translation: Translation
    ) -> Hexagon {
        let xy = 0.0;
        let x_shift = radius * 0.5; // r cos(pi/3)
        let y_shift = radius * 0.86602540378; // r sin(pi/3)
        
        let mut hexagon = Hexagon::init();

        // draw hexagon boundary
        // start north east corner
        // end east corner
        hexagon.set_next(
            Vertex::new(x_shift, y_shift, xy)
                .rot(rotation)
                .translate(translation)
        );
        hexagon.set_next(
            Vertex::new(radius, 0.0, xy)
                .rot(rotation)
                .translate(translation)
        );

        // start east corner
        // end south east corner
        hexagon.set_next(
            Vertex::new(radius, 0.0, xy)
                .rot(rotation)
                .translate(translation)
        );
        hexagon.set_next(
            Vertex::new(x_shift, -y_shift, xy)
                .rot(rotation)
                .translate(translation)
        );

        // start east corner
        // end south east corner
        hexagon.set_next(
            Vertex::new(radius, 0.0, xy)
                .rot(rotation)
                .translate(translation)
        );
        hexagon.set_next(
            Vertex::new(x_shift, -y_shift, xy)
                .rot(rotation)
                .translate(translation)
        );

        // start south east corner
        // end south west corner
        hexagon.set_next(
            Vertex::new(x_shift, -y_shift, xy)
                .rot(rotation)
                .translate(translation)
        );
        hexagon.set_next(
            Vertex::new(-x_shift, -y_shift, xy)
                .rot(rotation)
                .translate(translation)
        );

        // start south west corner
        // end west corner
        hexagon.set_next(
            Vertex::new(-x_shift, -y_shift, xy)
                .rot(rotation)
                .translate(translation)
        );
        hexagon.set_next(
            Vertex::new(-radius, 0.0, xy)
                .rot(rotation)
                .translate(translation)
        );

        // start west corner
        // end north west corner
        hexagon.set_next(
            Vertex::new(-radius, 0.0, xy)
                .rot(rotation)
                .translate(translation)
        );
        hexagon.set_next(
            Vertex::new(-x_shift, y_shift, xy)
                .rot(rotation)
                .translate(translation)
        );


        // start north west corner
        // end north east corner
        hexagon.set_next( 
            Vertex::new(-x_shift, y_shift, xy)
                .rot(rotation)
                .translate(translation)
        );
        hexagon.set_next(
            Vertex::new(x_shift, y_shift, xy)
                .rot(rotation)
                .translate(translation)
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
