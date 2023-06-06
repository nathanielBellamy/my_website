// COLOR
pub const INPUT_COLOR_1: &'static str = "magic_square_input_color_1";
pub const INPUT_COLOR_2: &'static str = "magic_square_input_color_2";
pub const INPUT_COLOR_3: &'static str = "magic_square_input_color_3";
pub const INPUT_COLOR_4: &'static str = "magic_square_input_color_4";
pub const INPUT_COLOR_5: &'static str = "magic_square_input_color_5";
pub const INPUT_COLOR_6: &'static str = "magic_square_input_color_6";
pub const INPUT_COLOR_7: &'static str = "magic_square_input_color_7";
pub const INPUT_COLOR_8: &'static str = "magic_square_input_color_8";

// LFO
pub const INPUT_LFO_1_AMP: &'static str = "magic_square_input_lfo_1_amp";
pub const INPUT_LFO_1_DEST: &'static str = "magic_square_input_lfo_1_dest";
pub const INPUT_LFO_1_FREQ: &'static str = "magic_square_input_lfo_1_freq";
pub const INPUT_LFO_1_PHASE: &'static str = "magic_square_input_lfo_1_phase";
pub const INPUT_LFO_1_SHAPE: &'static str = "magic_square_input_lfo_1_shape";

// PATTERN
pub const INPUT_DRAW_PATTERN: &'static str = "magic_square_input_draw_pattern";

// RADIUS
pub const INPUT_RADIUS_MIN: &'static str = "magic_square_input_radius_min";
pub const INPUT_RADIUS_STEP: &'static str = "magic_square_input_radius_step";

pub const INPUT_X_ROT_BASE: &'static str = "magic_square_input_x_rot_base";
pub const INPUT_Y_ROT_BASE: &'static str = "magic_square_input_y_rot_base";
pub const INPUT_Z_ROT_BASE: &'static str = "magic_square_input_z_rot_base";

pub const INPUT_X_ROT_SPREAD: &'static str = "magic_square_input_x_rot_spread";
pub const INPUT_Y_ROT_SPREAD: &'static str = "magic_square_input_y_rot_spread";
pub const INPUT_Z_ROT_SPREAD: &'static str = "magic_square_input_z_rot_spread";

pub const INPUT_X_AXIS_X_ROT_COEFF: &'static str = "magic_square_input_x_axis_x_rot_coeff";
pub const INPUT_X_AXIS_Y_ROT_COEFF: &'static str = "magic_square_input_x_axis_y_rot_coeff";
pub const INPUT_X_AXIS_Z_ROT_COEFF: &'static str = "magic_square_input_x_axis_z_rot_coeff";

pub const INPUT_Y_AXIS_X_ROT_COEFF: &'static str = "magic_square_input_y_axis_x_rot_coeff";
pub const INPUT_Y_AXIS_Y_ROT_COEFF: &'static str = "magic_square_input_y_axis_y_rot_coeff";
pub const INPUT_Y_AXIS_Z_ROT_COEFF: &'static str = "magic_square_input_y_axis_z_rot_coeff";

// TRANSLATION
pub const INPUT_TRANSLATION_X: &'static str = "magic_square_input_translation_x";
pub const INPUT_TRANSLATION_Y: &'static str = "magic_square_input_translation_y";
pub const INPUT_TRANSLATION_Z: &'static str = "magic_square_input_translation_z";
pub const INPUT_MOUSE_TRACKING: &'static str = "magic_square_input_mouse_tracking";

// whether or not this array is used
// it is useful within an IDE
// that checks the length defined in the type
// agains the length of the explicit array
// to ensure that all inputs are accounted for
// 
// also handy for copying and pasting when writing matches
pub const INPUT_IDS: [&'static str; 32] = [
    // COLOR
    INPUT_COLOR_1,
    INPUT_COLOR_2,
    INPUT_COLOR_3,
    INPUT_COLOR_4,
    INPUT_COLOR_5,
    INPUT_COLOR_6,
    INPUT_COLOR_7,
    INPUT_COLOR_8,

    // LFO
    INPUT_LFO_1_AMP,
    INPUT_LFO_1_DEST,
    INPUT_LFO_1_FREQ,
    INPUT_LFO_1_PHASE,
    INPUT_LFO_1_SHAPE,

    // PATTERN
    INPUT_DRAW_PATTERN,

    // RADIUS
    INPUT_RADIUS_MIN,
    INPUT_RADIUS_STEP,

    // ROTATION
    INPUT_X_ROT_BASE,
    INPUT_Y_ROT_BASE,
    INPUT_Z_ROT_BASE,
    INPUT_X_ROT_SPREAD,
    INPUT_Y_ROT_SPREAD,
    INPUT_Z_ROT_SPREAD,
    INPUT_X_AXIS_X_ROT_COEFF,
    INPUT_X_AXIS_Y_ROT_COEFF,
    INPUT_X_AXIS_Z_ROT_COEFF,
    INPUT_Y_AXIS_X_ROT_COEFF,
    INPUT_Y_AXIS_Y_ROT_COEFF,
    INPUT_Y_AXIS_Z_ROT_COEFF,

    // TRANSLATION
    INPUT_TRANSLATION_X,
    INPUT_TRANSLATION_Y,
    INPUT_TRANSLATION_Z,
    INPUT_MOUSE_TRACKING,
];
