use crate::magic_square::geometry::hexagon::Hexagon;
use crate::magic_square::geometry::icosohedron::Icosohedron;
use crate::magic_square::transformations::Transformation;

// store
pub mod cache;

// shapes
pub mod hexagon;
pub mod icosohedron;



pub struct Geometry;

#[derive(Clone, Copy)]
pub enum Shape {
    Triangle,
    Square,
    Pentagon,
    Hexagon,
    Icosohedron,
    None
}

impl Geometry {
    // per shape:
    //  shape -> accepts &mut Vertices, writes directly to array that will be passed to GL
    //  shape_cached -> Returns ShapeCache, array of vertices need to define the shape
    pub fn hexagon(
        radius: f32,
        transformation: Transformation
    ) -> Hexagon {
        Hexagon::new(radius, transformation)
    }

    pub fn icosohedron(
        radius: f32,
        transformation: Transformation
    ) -> Icosohedron {
        Icosohedron::new(radius, transformation)
    }
}

// struct used to write geometry to vertex array that will be passed to GL


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
//

