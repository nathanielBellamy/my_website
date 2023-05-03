
use std::ops::{Index, IndexMut};
use crate::magic_square::traits::CoordinateStore;
use crate::magic_square::vertices::{Vertex, Vertices};
use crate::magic_square::transformations::RotationSequence;


pub struct Hexagon;


impl Hexagon {
    pub fn new(buffer: &[f32; 2], radius: f32, rotation: RotationSequence, vertices: &mut Vertices) {
        let center_x = buffer[0];
        let center_y = buffer[1];
        let xy = 0.02 * center_x * center_y;

        let x_shift = radius * 0.5; // r cos(pi/3)
        let y_shift = radius * 0.86602540378; // r sin(pi/3)

        // draw hexagon boundary
        // start north east corner
        // end east corner
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x + radius, center_y, xy)
                .rot(rotation)
        );

        // start east corner
        // end south east corner
        vertices.set_next(
            Vertex::new(center_x + radius, center_y, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );

        // start east corner
        // end south east corner
        vertices.set_next(
            Vertex::new(center_x + radius, center_y, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );

        // start south east corner
        // end south west corner
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x - x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );

        // start south west corner
        // end west corner
        vertices.set_next(
            Vertex::new(center_x - x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x - radius, center_y, xy)
                .rot(rotation)
        );

        // start west corner
        // end north west corner
        vertices.set_next(
            Vertex::new(center_x - radius, center_y, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x - x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );


        // start north west corner
        // end north east corner
        vertices.set_next( 
            Vertex::new(center_x - x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );
    }
}

const HEXAGON_STORE_LEN: usize = 14;

struct HexagonStore {
    pub arr: [f32; HEXAGON_STORE_LEN], // # coordinates needed to define hexagon
    idx: usize
}

// TODO: VertexArray trait
// share indexing/adding/set_next
// whatever function eats Geometry can impliment a trait bound
impl HexagonStore {
    pub fn new() -> HexagonStore {
        HexagonStore { arr: [0.0; HEXAGON_STORE_LEN], idx: 0 }
    }
}

impl Index<usize> for HexagonStore {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for HexagonStore {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl CoordinateStore<HexagonStore> for HexagonStore {
    fn idx(&self) -> usize {
        self.idx
    }

    fn set_idx(&mut self, new_idx: usize) -> usize {
        self.idx = new_idx;
        self.idx
    }

    fn set_next(&mut self, vertex: Vertex) {
        if self.idx > self.arr.len() - 1 { return; }
        for i in 0..2 {
            self.arr[self.idx + i] = vertex[i]
        }
        self.idx += 3;
    }
}
