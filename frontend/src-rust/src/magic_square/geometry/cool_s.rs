
use crate::magic_square::geometry::vertex_store::VertexStore;
use crate::magic_square::geometry::vertices::Vertex;
use std::ops::{Index, IndexMut};

use super::vertices::VERTEX_ARRAY_SIZE;

pub const MISC_IDX_COOL_S: i32 = 1;

pub const VERTEX_COUNT_COOL_S: i32 = 28;

const R: f32 = 1.41421356237;// 2.0_f32.sqrt();

const Y0: f32 = R;
const Y1: f32 = 0.66 * R;
const Y2: f32 = 0.25 * R;
const Y3: f32 = 0.0 * R;
const Y4: f32 = -0.25 * R;
const Y5: f32 = -0.66 * R;
const Y6: f32 = - R;

const X0: f32 = -0.5;
const X1: f32 = -0.25;
const X2: f32 = 0.0;
const X3: f32 = 0.25;
const X4: f32 = 0.5;

pub struct CoolS {
    pub arr: [f32; 300], // # coordinates needed to define hexagon
    pub vertex_count: i32,
    idx: usize,
}

impl Index<usize> for CoolS {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for CoolS {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<CoolS> for CoolS {
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

impl CoolS {
    fn init() -> CoolS {
        CoolS {
            arr: [0.0; VERTEX_ARRAY_SIZE],
            idx: 0,
            vertex_count: VERTEX_COUNT_COOL_S,
        }
    }

    fn vertex(id: usize) -> Vertex {
        match id {
            0 => Vertex::new(X2, Y0, 0.0),
            1 => Vertex::new(X0, Y1, 0.0),
            2 => Vertex::new(X2, Y1, 0.0),
            3 => Vertex::new(X4, Y1, 0.0),
            4 => Vertex::new(X0, Y2, 0.0),
            5 => Vertex::new(X2, Y2, 0.0),
            6 => Vertex::new(X4, Y2, 0.0),
            7 => Vertex::new(X1, Y3, 0.0),
            8 => Vertex::new(X3, Y3, 0.0),
            9 => Vertex::new(X0, Y4, 0.0),
            10 => Vertex::new(X2, Y4, 0.0),
            11 => Vertex::new(X4, Y4, 0.0),
            12 => Vertex::new(X0, Y5, 0.0),
            13 => Vertex::new(X2, Y5, 0.0),
            14 => Vertex::new(X4, Y5, 0.0),
            15 => Vertex::new(X2, Y6, 0.0),
            _ => Vertex::new(0.0, 0.0, 0.0)
        }
    }
    // write to vertices
    // return array to be cached
    pub fn f32_array() -> [f32; VERTEX_ARRAY_SIZE] {
        let mut cool_s = CoolS::init();

        cool_s.set_next(CoolS::vertex(0));
        cool_s.set_next(CoolS::vertex(1));

        cool_s.set_next(CoolS::vertex(0));
        cool_s.set_next(CoolS::vertex(3));

        cool_s.set_next(CoolS::vertex(1));
        cool_s.set_next(CoolS::vertex(4));

        cool_s.set_next(CoolS::vertex(2));
        cool_s.set_next(CoolS::vertex(5));

        cool_s.set_next(CoolS::vertex(3));
        cool_s.set_next(CoolS::vertex(6));

        cool_s.set_next(CoolS::vertex(4));
        cool_s.set_next(CoolS::vertex(10));

        cool_s.set_next(CoolS::vertex(5));
        cool_s.set_next(CoolS::vertex(11));
        
        cool_s.set_next(CoolS::vertex(6));
        cool_s.set_next(CoolS::vertex(8));
 
        cool_s.set_next(CoolS::vertex(7));
        cool_s.set_next(CoolS::vertex(9));
        
        cool_s.set_next(CoolS::vertex(9));
        cool_s.set_next(CoolS::vertex(12));

        cool_s.set_next(CoolS::vertex(10));
        cool_s.set_next(CoolS::vertex(13));


        cool_s.set_next(CoolS::vertex(11));
        cool_s.set_next(CoolS::vertex(14));

        cool_s.set_next(CoolS::vertex(12));
        cool_s.set_next(CoolS::vertex(15));

        cool_s.set_next(CoolS::vertex(14));
        cool_s.set_next(CoolS::vertex(15));

        cool_s.arr
    }
}
