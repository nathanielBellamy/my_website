use ndarray::prelude::*;
use ndarray::Array;

pub struct Rotate;

impl Rotate {
   fn rotx_matrix(theta: f32) -> Array<f32, Ix2> {
        array![
            [1.0, 0.0, 0.0, 0.0],
            [0.0, theta.cos(), theta.sin(), 0.0],
            [0.0, -theta.sin(), theta.cos(), 0.0],
            [0.0, 0.0, 0.0, 1.0],
        ]
    }

    fn roty_matrix(theta: f32) -> Array<f32, Ix2> {
        array![
            [theta.cos(), 0.0, -theta.sin(), 0.0],
            [0.0, 1.0, 0.0, 0.0],
            [theta.sin(), 0.0, theta.cos(), 0.0],
            [0.0, 0.0, 0.0, 0.0]
        ]
    }

    fn rotz_matrix(theta: f32) -> Array<f32, Ix2> {
        array![
            [theta.cos(), theta.sin(), 0.0, 0.0],
            [-theta.sin(), theta.cos(), 0.0, 0.0],
            [0.0, 0.0, 1.0, 0.0],
            [0.0, 0.0, 0.0, 1.0]
        ]
    }
}
