use super::Shape;
use super::hexagon::Hexagon;
use super::icosahedron::Icosahedron;
use super::vertices::VERTEX_ARRAY_SIZE;

pub struct Geom;

pub type OffsetVc = (i32, i32);

type GeomVertArrays = [[f32; VERTEX_ARRAY_SIZE]; 2];
impl Geom {
    pub fn f32_array() -> [f32; VERTEX_ARRAY_SIZE * 2] {
        const TOTAL_LEN: usize = VERTEX_ARRAY_SIZE * 2;
        let mut result: [f32; TOTAL_LEN] = [0.0; TOTAL_LEN];
        let arrays: GeomVertArrays = [
            Hexagon::f32_array(),
            Icosahedron::f32_array()
        ];
        for (idx_arr, array) in arrays.iter().enumerate() {
            for (idx_el, el) in array.iter().enumerate() {
                result[idx_arr * VERTEX_ARRAY_SIZE + idx_el] = *el;
            }
        }
        result
    }

    pub fn into_offset_vc(shape: Shape) -> OffsetVc {
        match shape {
            Shape::Hexagon => (0, 14),
            Shape::Icosahedron => (100, 100),
            Shape::None => (0, 0),
            _ => (0, 0)
        }
    }
}

// #[derive(Clone, Copy, Debug)]
// struct DictEntry {
//     pub idx: i32,
//     pub shape: Shape,
//     pub vert_count: i32
// }

// const GEOM_DICT: [DictEntry; 3] = [
//     DictEntry { idx: 0, shape: Shape::Hexagon, vert_count: 7 },
//     DictEntry { idx: 1, shape: Shape::Icosahedron, vert_count: 100 },
//     DictEntry { idx: 2, shape: Shape::None, vert_count: 0 },
// ];
