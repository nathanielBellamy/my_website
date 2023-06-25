use serde::{Serialize, Deserialize};
use self::cache::CACHE_CAPACITY;

pub mod cache;
pub mod empty;
pub mod geom;
pub mod hexagon;
pub mod icosahedron;
pub mod vertex_store;
pub mod vertices;



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

pub type Shapes = [Shape; CACHE_CAPACITY];

