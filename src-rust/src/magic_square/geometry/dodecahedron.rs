
use crate::magic_square::geometry::vertex_store::VertexStore;
use crate::magic_square::geometry::vertices::Vertex;
use std::ops::{Index, IndexMut};

use super::vertices::VERTEX_ARRAY_SIZE;

pub const VERTEX_COUNT_DODECAHEDRON: i32 = 60;

const R: f32 = 1.41421356237;// 2.0_f32.sqrt();
const P: f32 = 1.61803398875; // 1/2 (1 + sqrt(5))

// (±1, ±1, ±1)
const V0: Vertex = Vertex { arr: [1.0, 1.0, 1.0] };
const V1: Vertex = Vertex { arr: [-1.0, 1.0, 1.0] };
const V2: Vertex = Vertex { arr: [1.0, -1.0, 1.0] };
const V3: Vertex = Vertex { arr: [1.0, 1.0, -1.0] };
const V4: Vertex = Vertex { arr: [-1.0, -1.0, 1.0] };
const V5: Vertex = Vertex { arr: [1.0, -1.0, -1.0] };
const V6: Vertex = Vertex { arr: [-1.0, 1.0, -1.0] };
const V7: Vertex = Vertex { arr: [-1.0, -1.0, -1.0] };

// (0, ±1/P, ±P }
const V8: Vertex = Vertex { arr: [0.0, 1.0 / P, P] };
const V9: Vertex = Vertex { arr: [0.0, -1.0 / P, P] };
const V10: Vertex = Vertex { arr: [0.0, 1.0 / P, -P] };
const V11: Vertex = Vertex { arr: [0.0, -1.0 / P, -P] };

// (±1/P, ±P, 0 }
const V12: Vertex = Vertex { arr: [1.0 / P, P, 0.0] };
const V13: Vertex = Vertex { arr: [-1.0 / P, P, 0.0] };
const V14: Vertex = Vertex { arr: [1.0 / P, -P, 0.0] };
const V15: Vertex = Vertex { arr: [-1.0 / P, -P, 0.0] };

// (±P, 0, ±1/P }
const V16: Vertex = Vertex { arr: [P, 0.0, 1.0 / P] };
const V17: Vertex = Vertex { arr: [-P, 0.0, 1.0 / P] };
const V18: Vertex = Vertex { arr: [P, 0.0, -1.0 / P] };
const V19: Vertex = Vertex { arr: [-P, 0.0, -1.0 / P] };


pub struct Dodecahedron {
    pub arr: [f32; 300], // # coordinates needed to define hexagon
    pub vertex_count: i32,
    idx: usize,
}

impl Index<usize> for Dodecahedron {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Dodecahedron {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<Dodecahedron> for Dodecahedron {
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

impl Dodecahedron {
    fn init() -> Dodecahedron {
        Dodecahedron {
            arr: [0.0; VERTEX_ARRAY_SIZE],
            idx: 0,
            vertex_count: VERTEX_COUNT_DODECAHEDRON,
        }
    }

    // write to vertices
    // return array to be cached
    pub fn f32_array() -> [f32; VERTEX_ARRAY_SIZE] {
        let mut dodecahedron = Dodecahedron::init();

        dodecahedron.set_next(V0);
        dodecahedron.set_next(V8);

        dodecahedron.set_next(V0);
        dodecahedron.set_next(V12);
        
        dodecahedron.set_next(V0);
        dodecahedron.set_next(V16);

        dodecahedron.set_next(V1);
        dodecahedron.set_next(V8);

        dodecahedron.set_next(V1);
        dodecahedron.set_next(V13);
        
        dodecahedron.set_next(V1);
        dodecahedron.set_next(V17);

        dodecahedron.set_next(V9);
        dodecahedron.set_next(V2);

        dodecahedron.set_next(V9);
        dodecahedron.set_next(V4);

        dodecahedron.set_next(V9);
        dodecahedron.set_next(V8);

        dodecahedron.set_next(V2);
        dodecahedron.set_next(V16);

        dodecahedron.set_next(V4);
        dodecahedron.set_next(V17);

        dodecahedron.set_next(V12);
        dodecahedron.set_next(V13);

        dodecahedron.set_next(V12);
        dodecahedron.set_next(V3);

        dodecahedron.set_next(V13);
        dodecahedron.set_next(V6);

        dodecahedron.set_next(V6);
        dodecahedron.set_next(V19);

        dodecahedron.set_next(V6);
        dodecahedron.set_next(V10);

        dodecahedron.set_next(V3);
        dodecahedron.set_next(V10);

        dodecahedron.set_next(V19);
        dodecahedron.set_next(V17);

        dodecahedron.set_next(V2);
        dodecahedron.set_next(V14);

        dodecahedron.set_next(V14);
        dodecahedron.set_next(V15);

        dodecahedron.set_next(V15);
        dodecahedron.set_next(V4);

        dodecahedron.set_next(V3);
        dodecahedron.set_next(V18);

        dodecahedron.set_next(V18);
        dodecahedron.set_next(V16);

        dodecahedron.set_next(V10);
        dodecahedron.set_next(V11);

        dodecahedron.set_next(V11);
        dodecahedron.set_next(V5);

        dodecahedron.set_next(V11);
        dodecahedron.set_next(V7);

        dodecahedron.set_next(V5);
        dodecahedron.set_next(V14);

        dodecahedron.set_next(V5);
        dodecahedron.set_next(V18);

        dodecahedron.set_next(V7);
        dodecahedron.set_next(V15);

        dodecahedron.set_next(V7);
        dodecahedron.set_next(V19);

        dodecahedron.arr
    }
}
