use serde::{Serialize, Deserialize};
use self::cache::CACHE_CAPACITY;

pub mod cache;
pub mod cube;
pub mod empty;
pub mod geom;
pub mod icosahedron;
pub mod ngon;
pub mod transformations;
pub mod vertex_store;
pub mod vertices;


#[derive(Clone, Copy, Debug, Default, Eq, PartialEq, Serialize, Deserialize)]
#[serde(tag = "t", content = "c")]
pub enum Shape {
    Ngon(u8),
    PlatoThree(u8),
    #[default]
    None
}

pub type Shapes = [Shape; CACHE_CAPACITY];

