use crate::magic_square::geometry::vertex_store::VertexStore;
use crate::magic_square::geometry::vertices::Vertex;
use std::ops::{Index, IndexMut};

use super::geom::PI;
use super::vertices::VERTEX_ARRAY_SIZE;

pub const VERTEX_COUNT_OCTAHEDRON: i32 = 18;

const R: f32 = 1.41421356237;// 2.0_f32.sqrt();
const THETA: [f32; 5] = [0.0, 1.57, 1.57, 1.57, PI]; //ARCCOS_NEG_ONE_THIRD, ARCCOS_NEG_ONE_THIRD, ARCCOS_NEG_ONE_THIRD];
const PHI: [f32; 5] = [0.0, 0.0, 2.0*PI/3.0, 4.0*PI/3.0, PI];

pub struct Octahedron {
    pub arr: [f32; 300], // # coordinates needed to define hexagon
    pub vertex_count: i32,
    idx: usize,
}

impl Index<usize> for Octahedron {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Octahedron {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<Octahedron> for Octahedron {
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

impl Octahedron {
    fn init() -> Octahedron {
        Octahedron {
            arr: [0.0; VERTEX_ARRAY_SIZE],
            idx: 0,
            vertex_count: VERTEX_COUNT_OCTAHEDRON,
        }
    }

    fn vertex(id: usize) -> Vertex {
        Vertex::new(
            R * THETA[id].sin() * PHI[id].cos(),
            R * THETA[id].sin() * PHI[id].sin(),
            R * THETA[id].cos(),
        )
    }
    // write to vertices
    // return array to be cached
    pub fn f32_array() -> [f32; VERTEX_ARRAY_SIZE] {
        let mut octahedron = Octahedron::init();

        octahedron.set_next(Octahedron::vertex(0));
        octahedron.set_next(Octahedron::vertex(1));
        
        octahedron.set_next(Octahedron::vertex(0));
        octahedron.set_next(Octahedron::vertex(2));

        octahedron.set_next(Octahedron::vertex(0));
        octahedron.set_next(Octahedron::vertex(3));

        octahedron.set_next(Octahedron::vertex(1));
        octahedron.set_next(Octahedron::vertex(2));

        octahedron.set_next(Octahedron::vertex(2));
        octahedron.set_next(Octahedron::vertex(3));

        octahedron.set_next(Octahedron::vertex(3));
        octahedron.set_next(Octahedron::vertex(1));

        octahedron.set_next(Octahedron::vertex(4));
        octahedron.set_next(Octahedron::vertex(3));

        octahedron.set_next(Octahedron::vertex(4));
        octahedron.set_next(Octahedron::vertex(2));

        octahedron.set_next(Octahedron::vertex(4));
        octahedron.set_next(Octahedron::vertex(1));
 
        octahedron.arr
    }
}
