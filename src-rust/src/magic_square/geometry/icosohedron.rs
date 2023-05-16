use std::ops::{Index, IndexMut};
use crate::magic_square::traits::VertexStore;
use crate::magic_square::vertices::{Vertex, VERTEX_ARRAY_SIZE};
use crate::magic_square::transformations::{RotationSequence, Translation};

const PI: f32 = std::f32::consts::PI;

// const H_ANGLE: f32 = PI / (180.0 * 72.0); // 72 degrees = 2pi/5; rotate horizontal
// const V_ANGLE: f32 = 0.5_f32.atan(); // elevation = 26.565 degrees
// const V_ANGLE_SIN: f32 = V_ANGLE.sin();
// const V_ANGLE_COS: f32 = V_ANGLE.cos();


pub struct Icosohedron {
    pub arr: [f32; VERTEX_ARRAY_SIZE], // # coordinates needed to define hexagon
    idx: usize
}

impl Icosohedron {
    fn init() -> Icosohedron {
        Icosohedron { arr: [0.0; VERTEX_ARRAY_SIZE], idx: 0 }
    }
    // write to vertices
    // return array to be cached 
    pub fn new(
        radius: f32, 
        rotation: RotationSequence,
        translation: Translation
    ) -> Icosohedron {
        let mut icosohedron = Icosohedron::init();

        let h_angle: f32 = PI / 180.0 * 72.0; // 72 degrees = 360/5
        let v_angle: f32 = 0.5_f32.atan(); // elevation = 26.565 degrees
        
        let z: f32 = radius * v_angle.sin();
        let xy: f32 = radius * v_angle.cos();

        let mut h_angle_top: f32 = -PI / 2.0 - h_angle / 2.0; // start from -126 degree 
        let mut h_angle_bottom: f32 = -PI / 2.0; // start from -90deg

        // draw icosahedron edges
        // connect top and bottom row vertices
        // each iteration draws a V - top -> bottom -> top
        for _ in 1..6 {
            let next_angle_top = h_angle_top + h_angle;
            icosohedron.set_next(
                Vertex::new(xy * h_angle_top.cos(), z, xy * h_angle_top.sin())
                    .rot(rotation)
                    .translate(translation)
            );
            icosohedron.set_next(
                Vertex::new(xy * h_angle_bottom.cos(), -z, xy * h_angle_bottom.sin())
                    .rot(rotation)
                    .translate(translation)
            );

            icosohedron.set_next(
                Vertex::new(xy * h_angle_bottom.cos(), -z, xy * h_angle_bottom.sin())
                    .rot(rotation)
                    .translate(translation)
            );
            icosohedron.set_next(
                Vertex::new(xy * next_angle_top.cos(), z, xy * next_angle_top.sin())
                    .rot(rotation)
                    .translate(translation)
            );
            
            h_angle_top += h_angle;
            h_angle_bottom += h_angle;
        }

        // connect bottom vertex to five bottom-row vertices
        for _ in 1..6 {
            icosohedron.set_next(
                Vertex::new(0.0, -radius, 0.0)
                    .rot(rotation)
                    .translate(translation)
            );
            icosohedron.set_next(
                Vertex::new(xy * h_angle_bottom.cos(), -z, xy * h_angle_bottom.sin())
                    .rot(rotation)
                    .translate(translation)
            );

            
            h_angle_bottom += h_angle;
        }

        // connect bottom row vertices
        for _ in 1..6 {
            let next_angle = h_angle_bottom + h_angle;
            icosohedron.set_next(
                Vertex::new(xy * h_angle_bottom.cos(), -z, xy * h_angle_bottom.sin())
                    .rot(rotation)
                    .translate(translation)
            );
            icosohedron.set_next(
                Vertex::new(xy * next_angle.cos(), -z, xy * next_angle.sin())
                    .rot(rotation)
                    .translate(translation)
            );

            
            h_angle_bottom += h_angle;
        }

        // connect top vertex to five top-row vertices
        for _ in 1..6 {
            icosohedron.set_next(
                Vertex::new(0.0, radius, 0.0)
                    .rot(rotation)
                    .translate(translation)
            );
            icosohedron.set_next(
                Vertex::new(xy * h_angle_top.cos(), z, xy * h_angle_top.sin())
                    .rot(rotation)
                    .translate(translation)
            );
            
            h_angle_top += h_angle;
        }

        // connect top-row vertices
        for _ in 1..6 {
            let next_angle = h_angle_top + h_angle;
            icosohedron.set_next(
                Vertex::new(xy * h_angle_top.cos(), z, xy * h_angle_top.sin())
                    .rot(rotation)
                    .translate(translation)
            );
            icosohedron.set_next(
                Vertex::new(xy * next_angle.cos(), z, xy * next_angle.sin())
                    .rot(rotation)
                    .translate(translation)
            );

            
            h_angle_top += h_angle;
        }

        icosohedron
    }
}

impl Index<usize> for Icosohedron {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Icosohedron {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<Icosohedron> for Icosohedron {
    fn idx(&self) -> usize {
        self.idx
    }

    fn set_idx(&mut self, new_idx: usize) -> usize {
        self.idx = new_idx;
        self.idx
    }

    fn arr(&mut self) -> &mut [f32] {
        &mut self.arr
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
