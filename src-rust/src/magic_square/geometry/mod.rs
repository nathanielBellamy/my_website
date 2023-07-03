use self::cache::CACHE_CAPACITY;
use serde::{Deserialize, Serialize};

pub mod cache;
pub mod cool_s;
pub mod cube;
pub mod dodecahedron;
pub mod empty;
pub mod geom;
pub mod icosahedron;
pub mod ngon;
pub mod octahedron;
pub mod star_five;
pub mod tetrahedron;
pub mod transformations;
pub mod vertex_store;
pub mod vertices;

#[derive(Clone, Copy, Debug, Default, Eq, PartialEq, Serialize, Deserialize)]
#[serde(tag = "t", content = "c")]
pub enum Shape {
    Misc(i32),
    Ngon(u8),
    PlatoThree(u8),
    #[default]
    None,
}

pub type Shapes = [Shape; CACHE_CAPACITY];
