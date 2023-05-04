use std::ops::{Index, IndexMut};
use crate::magic_square::vertices::Vertex;

pub trait VertexStore<T: Index<usize> + IndexMut<usize>> {
    // to be used on Structs with
    // pub arr: [f32; XXX]
    // idx: usize
    //
    // inteface for providing functionality to fill and read buffers
    // sharability is limited by varying array lengths
    fn idx(&self) -> usize;

    fn set_idx(&mut self, new_index: usize) -> usize;

    fn arr(&mut self) -> &mut [f32];

    fn set_next(&mut self, vertex: Vertex) {
        let idx: usize = self.idx();
        let arr: &mut [f32] = self.arr();
        if idx > arr.len() - 1 { return; }
        for i in 0..2 {
            arr[idx + i] = vertex[i]
        }
        self.set_idx(idx + 3);
    }

    fn set_next_slice(&mut self, slice: &mut [f32]) {
        let idx: usize = self.idx();
        let arr: &mut [f32] = self.arr();
        
        let new_idx: usize = idx + slice.len();
        if new_idx > arr.len() - 1 { return; }
        
        for (i, coord) in slice.iter().enumerate() {
            arr[idx + i] = *coord;
        }

        self.set_idx(idx + slice.len());
    }
}
