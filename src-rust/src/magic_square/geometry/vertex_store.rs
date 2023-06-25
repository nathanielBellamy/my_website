use std::ops::{Index, IndexMut};
use crate::magic_square::geometry::vertices::Vertex;

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
        for i in 0..3 {
            arr[idx + i] = vertex[i]
        }
        self.set_idx(idx + 3);
    }

    fn set_next_slice(&mut self, slice: &[f32]) {
        let idx: usize = self.idx();
        let arr: &mut [f32] = self.arr();
        
        let new_idx: usize = idx + slice.len();
        if new_idx > arr.len() - 1 { return; }
        
        for (i, coord) in slice.iter().enumerate() {
            arr[idx + i] = *coord;
        }

        self.set_idx(idx + slice.len());
    }

    // allow partial clear of vertices
    // we will have to see how performance goes
    fn zero(&mut self, clear_to_idx: Option<usize>) {
        let arr: &mut [f32] = self.arr();
        let max_idx: usize = match clear_to_idx {
            Some(int) => int,
            None => arr.len() - 1
        };

        for i in 1..max_idx {
            arr[i] = 0.0;
        }

        self.set_idx(0);
    }
}
