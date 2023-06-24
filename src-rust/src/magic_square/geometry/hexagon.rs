const X_SHIFT: f32 = 0.5; // cos(pi/3)
const Y_SHIFT: f32 =  0.86602540378; // sin(pi/3)

pub const VERTICES_HEXAGON: [f32; 42] = [
    // start north east corner
    // end east corner
    X_SHIFT, Y_SHIFT, 0.0,
    1.0, 0.0, 0.0,
    
    // start east corner
    // end south east corner
    1.0, 0.0, 0.0,
    X_SHIFT, -Y_SHIFT, 0.0,

    // start east corner
    // end south east corner
    1.0, 0.0, 0.0,
    X_SHIFT, -Y_SHIFT, 0.0,

    // start south east corner
    // end south west corner
    X_SHIFT, -Y_SHIFT, 0.0,
    -X_SHIFT, -Y_SHIFT, 0.0,

    // start south west corner
    // end west corner
    -X_SHIFT, -Y_SHIFT, 0.0,
    -1.0, 0.0, 0.0,

    // start west corner
    // end north west corner
    -1.0, 0.0, 0.0,
    -X_SHIFT, Y_SHIFT, 0.0,


    // start north west corner
    // end north east corner
    -X_SHIFT, Y_SHIFT, 0.0,
    X_SHIFT, Y_SHIFT, 0.0,
];
