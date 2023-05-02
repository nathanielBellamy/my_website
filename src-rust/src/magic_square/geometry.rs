
us crate::magic_square::vertices::{Vertex, Vertices};
use crate::magic_square::transformations::RotationSequence;

// struct used to write geometry to vertex array that will be passed to GL
pub struct Geometry;

const HEXAGON_CACHE_LEN: usize = 14;

struct HexagonCache {
    pub arr: [f32; HEXAGON_CACHE_LEN], // # coordinates needed to define hexagon
    idx: usize
}

// TODO: VertexArray trait
// share indexing/adding/set_next
// whatever function eats Geometry can impliment a trait bound
impl HexagonCache {
    pub fn new() -> HexagonCache {
        HexagonCache { arr: [0.0; HEXAGON_CACHE_LEN], idx: 0 }
    }
}

impl Geometry {
    // per shape:
    //  shape -> accepts &mut Vertices, writes directly to array that will be passed to GL
    //  shape_cached -> Returns ShapeCache, array of vertices need to define the shape
    pub fn hexagon_cache() -> HexagonCache {
        HexagonCache::new()
    }

    pub fn hexagon(buffer: &[f32; 2], radius: f32, rotation: RotationSequence, vertices: &mut Vertices) {
        let center_x = buffer[0];
        let center_y = buffer[1];
        let xy = 0.02 * center_x * center_y;

        let x_shift = radius * 0.5; // r cos(pi/3)
        let y_shift = radius * 0.86602540378; // r sin(pi/3)

        // draw hexagon boundary
        // start north east corner
        // end east corner
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x + radius, center_y, xy)
                .rot(rotation)
        );

        // start east corner
        // end south east corner
        vertices.set_next(
            Vertex::new(center_x + radius, center_y, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );

        // start east corner
        // end south east corner
        vertices.set_next(
            Vertex::new(center_x + radius, center_y, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );

        // start south east corner
        // end south west corner
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x - x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );

        // start south west corner
        // end west corner
        vertices.set_next(
            Vertex::new(center_x - x_shift, center_y - y_shift, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x - radius, center_y, xy)
                .rot(rotation)
        );

        // start west corner
        // end north west corner
        vertices.set_next(
            Vertex::new(center_x - radius, center_y, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x - x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );


        // start north west corner
        // end north east corner
        vertices.set_next( 
            Vertex::new(center_x - x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );
        vertices.set_next(
            Vertex::new(center_x + x_shift, center_y + y_shift, xy)
                .rot(rotation)
        );
    }
}



    //     let h_angle: f32 = PI / 180.0 * 72.0; // 72 degrees = 360/5
    //     let v_angle: f32 = 0.5_f32.atan(); // elevation = 26.565 degrees
    //     
    //     let mut h_angle_1: f32 = -PI / 2.0 - h_angle / 2.0; // start from -126 degree 
    //     let mut h_angle_2: f32 = -PI / 2.0; // start from -90 deg at 2nd row
    //     
    //     // coordinates
    //     let mut z: f32;
    //     let mut xy: f32;

    //     // compute 10 vertices on 1st and 2nd rows
    //     for _i in 1..5 {
    //         let mut v1: Vertex = center;
    //         let mut v2: Vertex = center;
    //         
    //         z = radius * v_angle.sin();
    //         xy = radius * v_angle.cos();

    //         v1[0] = xy * h_angle_1.cos() + center_x;
    //         v2[0] = xy * h_angle_2.cos() + center_x;
    //         v1[1] = xy * h_angle_1.sin() + center_y;
    //         v2[1] = xy * h_angle_2.sin() + center_y;
    //         v1[2] = z;
    //         v2[2] = -z;
    //         
    //         vertices.set_next(center);
    //         vertices.set_next(v1);
    //         vertices.set_next(center);
    //         vertices.set_next(v2);

    //         h_angle_1 += h_angle;
    //         h_angle_2 += h_angle;
    //     }

    //     // add bottom vertex (0, 0, -r)
    //     vertices.set_next(center);
    //     vertices.set_next([center_x, center_y, -radius]);

    //     vertices

        // TODO: subdivide the sides + edges
    // }
