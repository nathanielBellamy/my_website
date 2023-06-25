use crate::magic_square::geometry::vertices::VERTEX_ARRAY_SIZE;

pub struct Empty {
    pub arr: [f32; VERTEX_ARRAY_SIZE],
}

impl Empty {
    pub fn new() -> Empty {
        Empty {
            arr: [0.0; VERTEX_ARRAY_SIZE]
        }
    }
}
