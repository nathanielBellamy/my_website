use super::cool_s::{CoolS, MISC_IDX_COOL_S, VERTEX_COUNT_COOL_S};
use super::cube::VERTEX_COUNT_CUBE;
use super::dodecahedron::{Dodecahedron, VERTEX_COUNT_DODECAHEDRON};
use super::icosahedron::{Icosahedron, VERTEX_COUNT_ICOSAHEDRON};
use super::ngon::Ngon;
use super::octahedron::{Octahedron, VERTEX_COUNT_OCTAHEDRON};
use super::star_five::{StarFive, MISC_IDX_STAR_FIVE, VERTEX_COUNT_STAR_FIVE};
use super::tetrahedron::{Tetrahedron, VERTEX_COUNT_TETRAHEDRON};
use super::vertices::VERTEX_ARRAY_SIZE;
use super::Shape;
use crate::magic_square::geometry::cube::Cube;

pub struct Geom;

pub const PI: f32 = std::f32::consts::PI;

pub type OffsetVc = (i32, i32);

const SHAPE_COUNT: usize = 35;
const TOTAL_LEN: usize = VERTEX_ARRAY_SIZE * SHAPE_COUNT;

type GeomVertArrays = [[f32; VERTEX_ARRAY_SIZE]; SHAPE_COUNT];
impl Geom {
    pub fn f32_array() -> [f32; TOTAL_LEN] {
        let mut result: [f32; TOTAL_LEN] = [0.0; TOTAL_LEN];
        let arrays: GeomVertArrays = [
            Ngon::f32_array(3),
            Ngon::f32_array(4),
            Ngon::f32_array(5),
            Ngon::f32_array(6),
            Ngon::f32_array(7),
            Ngon::f32_array(8),
            Ngon::f32_array(9),
            Ngon::f32_array(10),
            Ngon::f32_array(11),
            Ngon::f32_array(12),
            Ngon::f32_array(13),
            Ngon::f32_array(14),
            Ngon::f32_array(15),
            Ngon::f32_array(16),
            Ngon::f32_array(17),
            Ngon::f32_array(18),
            Ngon::f32_array(19),
            Ngon::f32_array(20),
            Ngon::f32_array(21),
            Ngon::f32_array(22),
            Ngon::f32_array(23),
            Ngon::f32_array(24),
            Ngon::f32_array(25),
            Ngon::f32_array(26),
            Ngon::f32_array(27),
            Ngon::f32_array(28),
            Ngon::f32_array(29),
            Ngon::f32_array(30),
            Tetrahedron::f32_array(),
            Cube::f32_array(),
            Octahedron::f32_array(),
            Dodecahedron::f32_array(),
            Icosahedron::f32_array(),
            StarFive::f32_array(),
            CoolS::f32_array(),
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
            Shape::Misc(idx) => match idx {
                MISC_IDX_STAR_FIVE => (3300, VERTEX_COUNT_STAR_FIVE),
                MISC_IDX_COOL_S => (3400, VERTEX_COUNT_COOL_S),
                _ => (0, 0),
            },
            Shape::Ngon(n) => (100 * (n - 3) as i32, 2 * n as i32),
            Shape::PlatoThree(n) => match n {
                4 => (2800, VERTEX_COUNT_TETRAHEDRON),
                6 => (2900, VERTEX_COUNT_CUBE),
                8 => (3000, VERTEX_COUNT_OCTAHEDRON),
                12 => (3100, VERTEX_COUNT_DODECAHEDRON),
                20 => (3200, VERTEX_COUNT_ICOSAHEDRON),
                _ => (0, 0),
            },
            Shape::None => (0, 0),
        }
    }
}
