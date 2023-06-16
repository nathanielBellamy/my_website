use ndarray::prelude::*;
use ndarray::Array;
use crate::magic_square::main::Axis;
use super::settings::TransformOrder;

#[derive(Clone, Copy, Debug)]
pub struct Transformation {
    pub order: TransformOrder,
    pub rot_seq: RotationSequence,
    pub translation: Translation
}

#[derive(Clone, Copy, Debug)]
pub struct RotationSequence {
    pub arr: [Rotation; 3] // all rotations acheivable in three
}

impl RotationSequence {
    pub fn new(r1: Rotation, r2: Rotation, r3: Rotation) -> RotationSequence {
        RotationSequence { arr: [r1, r2, r3] }
    }

    pub fn matrix(&self) -> Array<f32, Ix2> {
        self.arr[2].matrix()
            .dot(
                &self.arr[1].matrix().dot(&self.arr[0].matrix())
            )
    }
}


#[derive(Clone, Copy, Debug)]
pub struct Rotation {
    axis: Axis,
    theta: f32
}

impl Rotation {
    pub fn new(axis: Axis, theta: f32) -> Rotation {
        Rotation { axis, theta }
    }
    
    pub fn matrix(&self) -> Array<f32, Ix2> {
        match self.axis {
            Axis::X => {
                array![
                    [1.0, 0.0, 0.0],
                    [0.0, self.theta.cos(), -self.theta.sin()],
                    [0.0, self.theta.sin(), self.theta.cos()],
                ]
            },
            Axis::Y => {
                array![
                    [self.theta.cos(), 0.0, self.theta.sin()],
                    [0.0, 1.0, 0.0],
                    [-self.theta.sin(), 0.0, self.theta.cos()],
                ]
            },           
            Axis::Z => {
                array![
                    [self.theta.cos(), self.theta.sin(), 0.0],
                    [-self.theta.sin(), self.theta.cos(), 0.0],
                    [0.0, 0.0, 1.0],
                ]
            },
        }
    } 
}

#[derive(Clone, Copy, Debug)]
pub struct Translation {
    pub x: f32,
    pub y: f32,
    pub z: f32
}
