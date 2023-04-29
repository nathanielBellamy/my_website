use std::f32::consts::PI;

pub type Vertex = [f32; 4];
pub type VertexArr = [f32; 400];

const ORIGIN: Vertex = [0.0, 0.0, 0.0, 0.0];

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

    pub fn set_next(&mut self, vertex: Vertex) {
        if self.idx > self.arr.len() - 1 { return; }
        for i in 0..3 {
            self.arr[self.idx + i] = vertex[i]
        }
        self.idx += 4;
    }

    pub fn icosahedron(buffer: &[f32;2], radius: f32) -> Vertices {
        // adopted from C++ code found here
        // https://www.songho.ca/opengl/gl_sphere.html
        let mut vertices = Vertices::new();

        let center_x = buffer[0];
        let center_y = buffer[1];
        let center: Vertex = [buffer[0], buffer[1], 0.0, 0.0];


         
        // first add top vertex (center_x, center_y, r)
        vertices.set_next([center_x, center_y, radius, 0.0]);
        vertices.set_next(center);


        let h_angle: f32 = PI / 180.0 * 72.0; // 72 degrees = 360/5
        let v_angle: f32 = 0.5_f32.atan(); // elevation = 26.565 degrees
        
        let mut h_angle_1: f32 = -PI / 2.0 - h_angle / 2.0; // start from -126 degree 
        let mut h_angle_2: f32 = -PI / 2.0; // start from -90 deg at 2nd row
        
        // coordinates
        let mut z: f32;
        let mut xy: f32;

        // compute 10 vertices on 1st and 2nd rows
        for _i in 1..5 {
            let mut v1: Vertex = center;
            let mut v2: Vertex = center;
            
            z = radius * v_angle.sin();
            xy = radius * v_angle.cos();

            v1[0] = xy * h_angle_1.cos() + center_x;
            v2[0] = xy * h_angle_2.cos() + center_x;
            v1[1] = xy * h_angle_1.sin() + center_y;
            v2[1] = xy * h_angle_2.sin() + center_y;
            v1[2] = z + center_x;
            v2[2] = -z + center_y;

            vertices.set_next(v1);
            vertices.set_next(center);
            vertices.set_next(v2);
            vertices.set_next(center);

            h_angle_1 += h_angle;
            h_angle_2 += h_angle;
        }

        // add bottom vertex (0, 0, -r)
        vertices.set_next([center_x, center_y, -radius, 0.0]);
        vertices.set_next(center);

        vertices

        // TODO: subdivide the sides + edges
    }
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
