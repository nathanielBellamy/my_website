use crate::magic_square::geometry::vertex_store::VertexStore;
use std::ops::{Index, IndexMut};

use super::geom::PI;
use super::vertices::{Vertex, VERTEX_ARRAY_SIZE};

pub const VERTEX_COUNT_HEPTAGON: i32 = 14;

pub struct Ngon {
    pub arr: [f32; VERTEX_ARRAY_SIZE],
    pub vert_count: i32,
    pub n: u8,
    idx: usize,
}

impl Index<usize> for Ngon {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Ngon {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<Ngon> for Ngon {
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

impl Ngon {
    fn init(n: u8) -> Ngon {
        Ngon {
            arr: [0.0; VERTEX_ARRAY_SIZE],
            idx: 0,
            vert_count: 2 * n as i32,
            n,
        }
    }

    pub fn polar(r: f32, k: f32, n: f32) -> (f32, f32) {
        (r * (k * 2.0 * PI / n).cos(), r * (k * 2.0 * PI / n).sin())
    }

    pub fn polar_vertex(r: f32, k: f32, n: f32) -> Vertex {
        let res: (f32, f32) = Ngon::polar(r, k, n);
        Vertex::new(res.0, res.1, 0.0)
    }

    pub fn f32_array(n: u8) -> [f32; VERTEX_ARRAY_SIZE] {
        let r: f32 = 2.0_f32.sqrt();
        let mut ngon = Ngon::init(n);
        let iter_max: usize = (ngon.n as usize) + 1;

        for k in 0..iter_max {
            ngon.set_next(Ngon::polar_vertex(r, k as f32, ngon.n as f32));
            ngon.set_next(Ngon::polar_vertex(r, (k + 1) as f32, ngon.n as f32));
        }

        ngon.arr
    }
}
