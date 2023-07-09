// see src/MagicSquare/WasmInputId.ts

// COLOR
pub const INPUT_COLORS: &'static str = "magic_square_input_colors";
pub const INPUT_COLOR_DIRECTION: &'static str = "magic_square_input_color_direction";
pub const INPUT_COLOR_SPEED: &'static str = "magic_square_input_color_speed";

// DRAW PATTERN
pub const INPUT_DRAW_PATTERN_TYPE: &'static str = "magic_square_input_draw_pattern_type";
pub const INPUT_DRAW_PATTERN_COUNT: &'static str = "magic_square_input_draw_pattern_count";
pub const INPUT_DRAW_PATTERN_OFFSET: &'static str = "magic_square_input_draw_pattern_offset";
pub const INPUT_DRAW_PATTERN_SPEED: &'static str = "magic_square_input_draw_pattern_speed";
pub const INPUT_TRANSFORM_ORDER: &'static str = "magic_square_input_transform_order";

// GEOMETRY
pub const INPUT_SHAPES: &'static str = "magic_square_input_shapes";
pub const INPUT_RADIUS_BASE: &'static str = "magic_square_input_radius_base";
pub const INPUT_RADIUS_STEP: &'static str = "magic_square_input_radius_step";
pub const INPUT_RADIUS_OFFSET: &'static str = "magic_square_input_radius_offset";

// LFO 1
pub const INPUT_LFO_1_ACTIVE: &'static str = "magic_square_input_lfo_1_active";
pub const INPUT_LFO_1_AMP: &'static str = "magic_square_input_lfo_1_amp";
pub const INPUT_LFO_1_DEST: &'static str = "magic_square_input_lfo_1_dest";
pub const INPUT_LFO_1_FREQ: &'static str = "magic_square_input_lfo_1_freq";
pub const INPUT_LFO_1_PHASE: &'static str = "magic_square_input_lfo_1_phase";
pub const INPUT_LFO_1_SHAPE: &'static str = "magic_square_input_lfo_1_shape";

// LFO 2
pub const INPUT_LFO_2_ACTIVE: &'static str = "magic_square_input_lfo_2_active";
pub const INPUT_LFO_2_AMP: &'static str = "magic_square_input_lfo_2_amp";
pub const INPUT_LFO_2_DEST: &'static str = "magic_square_input_lfo_2_dest";
pub const INPUT_LFO_2_FREQ: &'static str = "magic_square_input_lfo_2_freq";
pub const INPUT_LFO_2_PHASE: &'static str = "magic_square_input_lfo_2_phase";
pub const INPUT_LFO_2_SHAPE: &'static str = "magic_square_input_lfo_2_shape";

// LFO 3
pub const INPUT_LFO_3_ACTIVE: &'static str = "magic_square_input_lfo_3_active";
pub const INPUT_LFO_3_AMP: &'static str = "magic_square_input_lfo_3_amp";
pub const INPUT_LFO_3_DEST: &'static str = "magic_square_input_lfo_3_dest";
pub const INPUT_LFO_3_FREQ: &'static str = "magic_square_input_lfo_3_freq";
pub const INPUT_LFO_3_PHASE: &'static str = "magic_square_input_lfo_3_phase";
pub const INPUT_LFO_3_SHAPE: &'static str = "magic_square_input_lfo_3_shape";

// LFO 4
pub const INPUT_LFO_4_ACTIVE: &'static str = "magic_square_input_lfo_4_active";
pub const INPUT_LFO_4_AMP: &'static str = "magic_square_input_lfo_4_amp";
pub const INPUT_LFO_4_DEST: &'static str = "magic_square_input_lfo_4_dest";
pub const INPUT_LFO_4_FREQ: &'static str = "magic_square_input_lfo_4_freq";
pub const INPUT_LFO_4_PHASE: &'static str = "magic_square_input_lfo_4_phase";
pub const INPUT_LFO_4_SHAPE: &'static str = "magic_square_input_lfo_4_shape";

// PRESET
pub const INPUT_PRESET: &'static str = "magic_square_input_preset";

//ROTATION

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
pub const INPUT_TRANSLATION_X_BASE: &'static str = "magic_square_input_translation_x_base";
pub const INPUT_TRANSLATION_X_SPREAD: &'static str = "magic_square_input_translation_x_spread";
pub const INPUT_TRANSLATION_Y_BASE: &'static str = "magic_square_input_translation_y_base";
pub const INPUT_TRANSLATION_Y_SPREAD: &'static str = "magic_square_input_translation_y_spread";
pub const INPUT_TRANSLATION_Z_BASE: &'static str = "magic_square_input_translation_z_base";
pub const INPUT_TRANSLATION_Z_SPREAD: &'static str = "magic_square_input_translation_z_spread";
pub const INPUT_MOUSE_TRACKING: &'static str = "magic_square_input_mouse_tracking";

// whether or not this array is used
// it is useful within an IDE
// that checks the length defined in the type
// agains the length of the explicit array
// to ensure that all inputs are accounted for
//
// also handy for copying and pasting when writing matches
pub const INPUT_IDS: [&'static str; 55] = [
    // COLOR
    INPUT_COLORS,
    INPUT_COLOR_DIRECTION,
    INPUT_COLOR_SPEED,
    // DRAW PATTERN
    INPUT_DRAW_PATTERN_TYPE,
    INPUT_DRAW_PATTERN_COUNT,
    INPUT_DRAW_PATTERN_OFFSET,
    INPUT_DRAW_PATTERN_SPEED,
    // GEOMETRY
    INPUT_SHAPES,
    INPUT_RADIUS_BASE,
    INPUT_RADIUS_STEP,
    INPUT_TRANSFORM_ORDER,
    // LFO_1
    INPUT_LFO_1_ACTIVE,
    INPUT_LFO_1_AMP,
    INPUT_LFO_1_DEST,
    INPUT_LFO_1_FREQ,
    INPUT_LFO_1_PHASE,
    INPUT_LFO_1_SHAPE,
    // LFO_2
    INPUT_LFO_2_ACTIVE,
    INPUT_LFO_2_AMP,
    INPUT_LFO_2_DEST,
    INPUT_LFO_2_FREQ,
    INPUT_LFO_2_PHASE,
    INPUT_LFO_2_SHAPE,
    // LFO_3
    INPUT_LFO_3_ACTIVE,
    INPUT_LFO_3_AMP,
    INPUT_LFO_3_DEST,
    INPUT_LFO_3_FREQ,
    INPUT_LFO_3_PHASE,
    INPUT_LFO_3_SHAPE,
    // LFO_4
    INPUT_LFO_4_ACTIVE,
    INPUT_LFO_4_AMP,
    INPUT_LFO_4_DEST,
    INPUT_LFO_4_FREQ,
    INPUT_LFO_4_PHASE,
    INPUT_LFO_4_SHAPE,
    // PRESET
    INPUT_PRESET,
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
    INPUT_TRANSLATION_X_BASE,
    INPUT_TRANSLATION_X_SPREAD,
    INPUT_TRANSLATION_Y_BASE,
    INPUT_TRANSLATION_Y_SPREAD,
    INPUT_TRANSLATION_Z_BASE,
    INPUT_TRANSLATION_Z_SPREAD,
    INPUT_MOUSE_TRACKING,
];
