use serde::{Serialize, Deserialize};
use crate::magic_square::geometry::empty::Empty;
use crate::magic_square::geometry::hexagon::Hexagon;
use crate::magic_square::geometry::icosahedron::Icosahedron;

use self::cache::CACHE_SHAPE_WIDTH;

// store
pub mod cache;

// shapes
pub mod empty;
pub mod hexagon;
pub mod icosahedron;


#[derive(Clone, Copy, Debug)]
pub struct Geometry { shape: Shape }

#[derive(Clone, Copy, Debug, Default, Eq, PartialEq, Serialize, Deserialize)]
pub enum Shape {
    Triangle,
    Square,
    Pentagon,
    Hexagon,
    Icosahedron,
    #[default]
    None
}

impl Geometry {
    pub fn new(shape: Shape) -> Geometry {
        Geometry { shape }
    }

    pub fn arr(&self) -> [f32; CACHE_SHAPE_WIDTH] {
        match self.shape {
            Shape::Hexagon => Hexagon::new().arr,
            Shape::Icosahedron => Icosahedron::new().arr,
            _ => Empty::new().arr
        }
    }

    pub fn to_vertex_count(shape: Shape) -> usize {
        match shape {
            Shape::Hexagon => 16,
            Shape::Icosahedron => 100,
            _ => 100,
        }
    }
}
