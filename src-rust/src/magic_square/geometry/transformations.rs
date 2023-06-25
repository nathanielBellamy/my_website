use crate::magic_square::main::Axis;

pub type Mat4 = [f32; 16]; 
pub const MAT4_ID: Mat4 = [
    1.0, 0.0, 0.0, 0.0,
    0.0, 1.0, 0.0, 0.0,
    0.0, 0.0, 1.0, 0.0,
    0.0, 0.0, 0.0, 1.0,
];
pub const MAT4_ZERO: Mat4 = [
    0.0, 0.0, 0.0, 0.0,
    0.0, 0.0, 0.0, 0.0,
    0.0, 0.0, 0.0, 0.0,
    0.0, 0.0, 0.0, 0.0,
];

pub struct RotationSequence {
    pub arr: [Rotation; 3] // all rotations acheivable in three
}

impl RotationSequence {
    pub fn new(r1: Rotation, r2: Rotation, r3: Rotation) -> RotationSequence {
        RotationSequence { arr: [r1, r2, r3] }
    }

    pub fn default() -> RotationSequence {
        RotationSequence {
            arr: [
                Rotation { axis: Axis::X, theta: 0.0 },
                Rotation { axis: Axis::Y, theta: 0.0 },
                Rotation { axis: Axis::Z, theta: 0.0 },
            ]
        }
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
    
    pub fn matrix(&self) -> Mat4 {
        match self.axis {
            Axis::X => {
                [
                    1.0, 0.0, 0.0, 0.0,
                    0.0, self.theta.cos(), -self.theta.sin(), 0.0,
                    0.0, self.theta.sin(), self.theta.cos(), 0.0,
                    0.0, 0.0, 0.0, 1.0,
                ]
            },
            Axis::Y => {
                [
                    self.theta.cos(), 0.0, self.theta.sin(), 0.0,
                    0.0, 1.0, 0.0, 0.0,
                    -self.theta.sin(), 0.0, self.theta.cos(), 0.0,
                    0.0, 0.0, 0.0, 1.0,
                ]
            },           
            Axis::Z => {
                [
                    self.theta.cos(), self.theta.sin(), 0.0, 0.0,
                    -self.theta.sin(), self.theta.cos(), 0.0, 0.0,
                    0.0, 0.0, 1.0, 0.0,
                    0.0, 0.0, 0.0, 1.0,
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

impl Translation {
    pub fn default() -> Translation {
        Translation {
            x: 0.0,
            y: 0.0,
            z: 0.0
        }
    }

    pub fn arr(&self) -> [f32; 3] {
        [self.x, self.y, self.z]
    }

    pub fn matrix(&self) -> Mat4 {
        [
            1.0,  0.0,  0.0,  0.0,
            0.0,  1.0,  0.0,  0.0,
            0.0,  0.0,  1.0,  0.0,
            self.x, self.y, self.z, 1.0,
        ]
    }
}

pub struct Projection;

impl Projection {
    pub fn z_zero() -> Mat4 {
        [
            1.0,  0.0,  0.0,  0.0,
            0.0,  1.0,  0.0,  0.0,
            0.0,  0.0,  0.0,  0.0,
            0.0,  0.0,  0.0,  1.0,
        ]
    }
}
