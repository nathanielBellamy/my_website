use crate::magic_square::geometry::vertex_store::VertexStore;
use crate::magic_square::geometry::vertices::Vertex;
use std::ops::{Index, IndexMut};

use super::geom::PI;
use super::vertices::VERTEX_ARRAY_SIZE;

pub const MISC_IDX_STAR_FIVE: i32 = 0;

pub const VERTEX_COUNT_STAR_FIVE: i32 = 10;

const R: f32 = 1.41421356237;// 2.0_f32.sqrt();
const THETA: [f32; 5] = [PI/8.0, PI/2.0, 7.0*PI/8.0, 10.0*PI/8.0, 14.0*PI/8.0];

pub struct StarFive {
    pub arr: [f32; 300], // # coordinates needed to define hexagon
    pub vertex_count: i32,
    idx: usize,
}

impl Index<usize> for StarFive {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for StarFive {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<StarFive> for StarFive {
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

impl StarFive {
    fn init() -> StarFive {
        StarFive {
            arr: [0.0; VERTEX_ARRAY_SIZE],
            idx: 0,
            vertex_count: VERTEX_COUNT_STAR_FIVE,
        }
    }

    fn vertex(id: usize) -> Vertex {
        Vertex::new(
            R * THETA[id].cos(),
            R * THETA[id].sin(),
            0.0,
        )
    }
    // write to vertices
    // return array to be cached
    pub fn f32_array() -> [f32; VERTEX_ARRAY_SIZE] {
        let mut star_five = StarFive::init();

        star_five.set_next(StarFive::vertex(3));
        star_five.set_next(StarFive::vertex(1));
 
        star_five.set_next(StarFive::vertex(1));
        star_five.set_next(StarFive::vertex(4));

        star_five.set_next(StarFive::vertex(4));
        star_five.set_next(StarFive::vertex(2));

        star_five.set_next(StarFive::vertex(2));
        star_five.set_next(StarFive::vertex(0));

        star_five.set_next(StarFive::vertex(0));
        star_five.set_next(StarFive::vertex(3));
 
        star_five.arr
    }
}
