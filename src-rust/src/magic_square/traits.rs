use std::ops::{Index, IndexMut};
use crate::magic_square::vertices::Vertex;

pub trait CoordinateStore<T: Index<usize> + IndexMut<usize>> {
    // to be used on Structs with
    // pub arr: [f32; XXX]
    // idx: usize
    //
    // inteface for providing functionality to fill and read buffers
    // sharability is limited by varying array lengths
    fn idx(&self) -> usize;

    fn set_idx(&mut self, new_index: usize) -> usize;

    fn set_next(&mut self, vertex: Vertex);
}
