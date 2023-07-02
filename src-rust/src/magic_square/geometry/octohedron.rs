
use crate::magic_square::geometry::vertex_store::VertexStore;
use crate::magic_square::geometry::vertices::Vertex;
use std::ops::{Index, IndexMut};

use super::geom::PI;
use super::vertices::VERTEX_ARRAY_SIZE;

pub const VERTEX_COUNT_OCTOHEDRON: i32 = 18;

const R: f32 = 1.41421356237;// 2.0_f32.sqrt();
const THETA: [f32; 5] = [0.0, 1.57, 1.57, 1.57, PI]; //ARCCOS_NEG_ONE_THIRD, ARCCOS_NEG_ONE_THIRD, ARCCOS_NEG_ONE_THIRD];
const PHI: [f32; 5] = [0.0, 0.0, 2.0*PI/3.0, 4.0*PI/3.0, PI];

pub struct Octohedron {
    pub arr: [f32; 300], // # coordinates needed to define hexagon
    pub vertex_count: i32,
    idx: usize,
}

impl Index<usize> for Octohedron {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Octohedron {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<Octohedron> for Octohedron {
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

impl Octohedron {
    fn init() -> Octohedron {
        Octohedron {
            arr: [0.0; VERTEX_ARRAY_SIZE],
            idx: 0,
            vertex_count: VERTEX_COUNT_OCTOHEDRON,
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
        let mut octohedron = Octohedron::init();

        octohedron.set_next(Octohedron::vertex(0));
        octohedron.set_next(Octohedron::vertex(1));
        
        octohedron.set_next(Octohedron::vertex(0));
        octohedron.set_next(Octohedron::vertex(2));

        octohedron.set_next(Octohedron::vertex(0));
        octohedron.set_next(Octohedron::vertex(3));

        octohedron.set_next(Octohedron::vertex(1));
        octohedron.set_next(Octohedron::vertex(2));

        octohedron.set_next(Octohedron::vertex(2));
        octohedron.set_next(Octohedron::vertex(3));

        octohedron.set_next(Octohedron::vertex(3));
        octohedron.set_next(Octohedron::vertex(1));

        octohedron.set_next(Octohedron::vertex(4));
        octohedron.set_next(Octohedron::vertex(3));

        octohedron.set_next(Octohedron::vertex(4));
        octohedron.set_next(Octohedron::vertex(2));

        octohedron.set_next(Octohedron::vertex(4));
        octohedron.set_next(Octohedron::vertex(1));
 
        octohedron.arr
    }
}
