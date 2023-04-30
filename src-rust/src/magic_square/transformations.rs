use ndarray::prelude::*;
use ndarray::Array;
use crate::magic_square::main::Axis;

pub type RotationSequence = [Rotation; 3]; // all rotations acheivable in three

pub struct Rotation {
    axis: Axis,
    theta: f32
}

impl Rotation {
    pub fn matrix(&self) -> Array<f32, Ix2> {
        match self.axis {
            Axis::X => {
                array![

                    // dim 3
                    [1.0, 0.0, 0.0],
                    [0.0, self.theta.cos(), self.theta.sin()],
                    [0.0, -self.theta.sin(), self.theta.cos()],
                    
                    // dim 4
                    // [1.0, 0.0, 0.0, 0.0],
                    // [0.0, theta.cos(), theta.sin(), 0.0],
                    // [0.0, -theta.sin(), theta.cos(), 0.0],
                    // [0.0, 0.0, 0.0, 1.0],
                ]
            },
            Axis::Y => {
                array![

                    // dim 3
                    [1.0, 0.0, 0.0],
                    [0.0, self.theta.cos(), self.theta.sin()],
                    [0.0, -self.theta.sin(), self.theta.cos()],
                    
                    // dim 4
                    // [1.0, 0.0, 0.0, 0.0],
                    // [0.0, theta.cos(), theta.sin(), 0.0],
                    // [0.0, -theta.sin(), theta.cos(), 0.0],
                    // [0.0, 0.0, 0.0, 1.0],
                ]
            },           
            Axis::Z => {
                array![

                    // dim 3
                    [1.0, 0.0, 0.0],
                    [0.0, self.theta.cos(), self.theta.sin()],
                    [0.0, -self.theta.sin(), self.theta.cos()],
                    
                    // dim 4
                    // [1.0, 0.0, 0.0, 0.0],
                    // [0.0, theta.cos(), theta.sin(), 0.0],
                    // [0.0, -theta.sin(), theta.cos(), 0.0],
                    // [0.0, 0.0, 0.0, 1.0],
                ]
            },
        }
    } 
}
