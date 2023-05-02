use std::ops::{Index, IndexMut};
use std::convert::From;
use ndarray::prelude::*;
use ndarray::Array;
use crate::magic_square::transformations::RotationSequence;



// pub type Vertex = [f32; 3];
pub type VertexArr = [f32; 400];

const ORIGIN: Vertex = Vertex { arr: [0.0, 0.0, 0.0] };

#[derive(Clone, Copy)]
pub struct Vertex {
    arr: [f32; 3]
}

impl From<[f32; 3]> for Vertex {
    fn from(arr: [f32; 3]) -> Self {
        Vertex { arr }
    }
}

impl Vertex {
    pub fn new(x: f32, y: f32, z: f32) -> Vertex {
        Vertex { arr: [x, y, z] }
    }

    pub fn lh_mult(&self, matrix: Array<f32, Ix2>) -> Vertex {
        Vertex {
            arr: [
                self[0] * matrix[[0,0]] + self[1] * matrix[[0, 1]] + self[2] * matrix[[0,2]],
                self[0] * matrix[[1,0]] + self[1] * matrix[[1, 1]] + self[2] * matrix[[1,2]],
                self[0] * matrix[[2,0]] + self[1] * matrix[[2, 1]] + self[2] * matrix[[2,2]],
            ]
        }
    }

    pub fn rot(&self, rotation: RotationSequence) ->  Vertex {
        self.lh_mult(rotation.matrix())
    }
}

impl Index<usize> for Vertex {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Vertex {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

pub struct Vertices {
    pub arr: VertexArr,
    idx: usize
}

impl Vertices {
    pub fn new() -> Vertices {
        Vertices { 
            arr: [0.0; 400], 
            idx: 0 
        }
    }

    pub fn add_geometry(&mut self) {
        self.arr;
    }

    pub fn set_next(&mut self, vertex: Vertex) {
        if self.idx > self.arr.len() - 1 { return; }
        for i in 0..2 {
            self.arr[self.idx + i] = vertex[i]
        }
        self.idx += 3;
    }

    pub fn hexagon(buffer: &[f32; 2], radius: f32, rotation: RotationSequence) -> Vertices {
        let mut vertices = Vertices::new();

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

        vertices
    }

    // pub fn icosahedron(buffer: &[f32;2], radius: f32) -> Vertices {
    //     // adopted from C++ code found here
    //     // https://www.songho.ca/opengl/gl_sphere.html
    //     let mut vertices = Vertices::new();

    //     let center_x = buffer[0];
    //     let center_y = buffer[1];
    //     let center = Vertex::new(buffer[0], buffer[1], 0.0);
    //      
    //     // first add top vertex (center_x, center_y, r)
    //     vertices.set_next(center);
    //     vertices.set_next([center_x, center_y, radius]);


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
}

    // fn get_vertices(buffer: &[i32; 2], idx: usize, axis: char, height: i32, width: i32) -> [f32; 8] {
    //     let mut result: [f32; 8] = [0.0; 8];
    //     
    //     let clip_x: f32 = (2.0 * (buffer[0] as f32) / width as f32) - 1.0;
    //     let clip_y: f32 = 1.0 - ((2.0 * buffer[1] as f32) / height as f32);

    //     let line_base: Array<f32, _> = array![
    //         [clip_x, clip_y],
    //         [0.0, 0.0],
    //         [0.0, 0.0],
    //         [0.0, 0.0],
    //     ];
    //     
    //     let theta: f32 = buffer[0] as f32 / (100.0 * idx as f32);
    //     let rot_matrix = match axis {
    //         'y' => MagicSquare::roty_matrix(theta),
    //         'z' => MagicSquare::rotz_matrix(theta),
    //         _ => MagicSquare::rotx_matrix(theta),
    //     };
    //     let rotated_line: Array<f32, _> = rot_matrix.dot(&line_base);
    //     // let rotated_line = line_base;

    //     // flatten
    //     let mut counter: usize = 0;
    //     for coord in rotated_line.iter() {
    //         result[counter] = *coord;
    //         counter += 1;
    //     }

    //     result
    // }
