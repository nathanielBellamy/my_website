use super::Shape;
use super::hexagon::Hexagon;

pub struct Geom;

pub type IdxVc = (i32, i32);

impl Geom {
    pub fn f32_array() -> [f32; 42] {
        Hexagon::f32_array()
    }

    pub fn into_idx_vc(shape: Shape) -> IdxVc {
        match shape {
            Shape::Hexagon => (0, 14),
            Shape::Icosahedron => (1, 100),
            Shape::None => (-1, -1),
            _ => (-1, -1)
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
