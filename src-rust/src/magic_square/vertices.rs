use std::ops::{Index, IndexMut};
use std::convert::From;
use ndarray::prelude::*;
use ndarray::Array;
use crate::magic_square::transformations::RotationSequence;
use super::traits::VertexStore;

// pub type Vertex = [f32; 3];
pub type VertexArr = [f32; 21000];

const ORIGIN: Vertex = Vertex { arr: [0.0, 0.0, 0.0] };
const VERTEX_ARRAY_SIZE: usize = 21_000;

#[derive(Clone, Copy)]
pub struct Vertex {
    arr: [f32; 3]
}

impl From<[f32; 3]> for Vertex {
    fn from(arr: [f32; 3]) -> Self {
        Vertex { arr }
    }
}

impl Vertex {
    pub fn new(x: f32, y: f32, z: f32) -> Vertex {
        Vertex { arr: [x, y, z] }
    }

    pub fn lh_mult(&self, matrix: Array<f32, Ix2>) -> Vertex {
        Vertex {
            arr: [
                self[0] * matrix[[0,0]] + self[1] * matrix[[0, 1]] + self[2] * matrix[[0,2]],
                self[0] * matrix[[1,0]] + self[1] * matrix[[1, 1]] + self[2] * matrix[[1,2]],
                self[0] * matrix[[2,0]] + self[1] * matrix[[2, 1]] + self[2] * matrix[[2,2]],
            ]
        }
    }

    pub fn rot(&self, rotation: RotationSequence) ->  Vertex {
        self.lh_mult(rotation.matrix())
    }
}

impl Index<usize> for Vertex {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Vertex {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

pub struct Vertices {
    pub arr: VertexArr,
    idx: usize
}

impl Index<usize> for Vertices {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Vertices {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<Vertices> for Vertices {
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

impl Vertices {
    pub fn new() -> Vertices {
        Vertices { 
            arr: [0.0; VERTEX_ARRAY_SIZE], 
            idx: 0 
        }
    }

    pub fn add_geometry(&mut self) {
        // TODO
        self.arr;
    }

    pub fn set_slice<T>(geometry: T) {
        //
    }
}
