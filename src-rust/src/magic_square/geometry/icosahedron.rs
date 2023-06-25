use std::ops::{Index, IndexMut};
use crate::magic_square::geometry::vertex_store::VertexStore;
use crate::magic_square::geometry::vertices::Vertex;

use super::vertices::VERTEX_ARRAY_SIZE;

const PI: f32 = std::f32::consts::PI;

pub struct Icosahedron {
    pub arr: [f32; 300], // # coordinates needed to define hexagon
    idx: usize,
}

impl Index<usize> for Icosahedron {
    type Output = f32;
    fn index<'a>(&'a self, i: usize) -> &'a f32 {
        &self.arr[i]
    }
}

impl IndexMut<usize> for Icosahedron {
    fn index_mut<'a>(&'a mut self, i: usize) -> &'a mut f32 {
        &mut self.arr[i]
    }
}

impl VertexStore<Icosahedron> for Icosahedron {
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

impl Icosahedron {
    fn init() -> Icosahedron {
        Icosahedron { arr: [0.0; VERTEX_ARRAY_SIZE], idx: 0 }
    }
    // write to vertices
    // return array to be cached 
    pub fn f32_array() -> [f32; VERTEX_ARRAY_SIZE] {
        let mut icosahedron = Icosahedron::init();

        let h_angle: f32 = PI / 180.0 * 72.0; // 72 degrees = 360/5
        let v_angle: f32 = 0.5_f32.atan(); // elevation = 26.565 degrees
        
        let z: f32 = v_angle.sin();
        let xy: f32 = v_angle.cos();

        let mut h_angle_top: f32 = -PI / 2.0 - h_angle / 2.0; // start from -126 degree 
        let mut h_angle_bottom: f32 = -PI / 2.0; // start from -90deg

        // draw icosahedron edges
        // connect top and bottom row vertices
        // each iteration draws a V - top -> bottom -> top
        for _ in 1..6 {
            let next_angle_top = h_angle_top + h_angle;
            icosahedron.set_next(
                Vertex::new(xy * h_angle_top.cos(), z, xy * h_angle_top.sin())
            );
            icosahedron.set_next(
                Vertex::new(xy * h_angle_bottom.cos(), -z, xy * h_angle_bottom.sin())
            );

            icosahedron.set_next(
                Vertex::new(xy * h_angle_bottom.cos(), -z, xy * h_angle_bottom.sin())
            );
            icosahedron.set_next(
                Vertex::new(xy * next_angle_top.cos(), z, xy * next_angle_top.sin())
            );
            
            h_angle_top += h_angle;
            h_angle_bottom += h_angle;
        }

        // connect bottom vertex to five bottom-row vertices
        for _ in 1..6 {
            icosahedron.set_next(
                Vertex::new(0.0, -1.0, 0.0)
            );
            icosahedron.set_next(
                Vertex::new(xy * h_angle_bottom.cos(), -z, xy * h_angle_bottom.sin())
            );

            
            h_angle_bottom += h_angle;
        }

        // connect bottom row vertices
        for _ in 1..6 {
            let next_angle = h_angle_bottom + h_angle;
            icosahedron.set_next(
                Vertex::new(xy * h_angle_bottom.cos(), -z, xy * h_angle_bottom.sin())
            );
            icosahedron.set_next(
                Vertex::new(xy * next_angle.cos(), -z, xy * next_angle.sin())
            );

            
            h_angle_bottom += h_angle;
        }

        // connect top vertex to five top-row vertices
        for _ in 1..6 {
            icosahedron.set_next(
                Vertex::new(0.0, 1.0, 0.0)
            );
            icosahedron.set_next(
                Vertex::new(xy * h_angle_top.cos(), z, xy * h_angle_top.sin())
            );
            
            h_angle_top += h_angle;
        }

        // connect top-row vertices
        for _ in 1..6 {
            let next_angle = h_angle_top + h_angle;
            icosahedron.set_next(
                Vertex::new(xy * h_angle_top.cos(), z, xy * h_angle_top.sin())
            );
            icosahedron.set_next(
                Vertex::new(xy * next_angle.cos(), z, xy * next_angle.sin())
            );
            
            h_angle_top += h_angle;
        }

        icosahedron.arr
    }
}


