use crate::magic_square::geometry::cube::Cube;
use super::Shape;
use super::cube::VERTEX_COUNT_CUBE;
use super::hexagon::{Hexagon, VERTEX_COUNT_HEXAGON};
use super::icosahedron::{Icosahedron, VERTEX_COUNT_ICOSAHEDRON};
use super::square::{Square, VERTEX_COUNT_SQUARE};
use super::triangle::{Triangle, VERTEX_COUNT_TRIANGLE};
use super::vertices::VERTEX_ARRAY_SIZE;

pub struct Geom;

pub type OffsetVc = (i32, i32);

const SHAPE_COUNT: usize = 5;
const TOTAL_LEN: usize = VERTEX_ARRAY_SIZE * SHAPE_COUNT;

type GeomVertArrays = [[f32; VERTEX_ARRAY_SIZE]; SHAPE_COUNT];
impl Geom {
    pub fn f32_array() -> [f32; TOTAL_LEN] {
        let mut result: [f32; TOTAL_LEN] = [0.0; TOTAL_LEN];
        let arrays: GeomVertArrays = [
            Hexagon::f32_array(),
            Cube::f32_array(),
            Icosahedron::f32_array(),
            Square::f32_array(),
            Triangle::f32_array(),
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
            Shape::Hexagon => (0, VERTEX_COUNT_HEXAGON),
            Shape::Cube => (100, VERTEX_COUNT_CUBE),
            Shape::Icosahedron => (200, VERTEX_COUNT_ICOSAHEDRON),
            Shape::Square => (300, VERTEX_COUNT_SQUARE),
            Shape::Triangle => (400, VERTEX_COUNT_TRIANGLE),
            Shape::None => (0, 0),
            _ => (0, 0)
        }
    }
}
