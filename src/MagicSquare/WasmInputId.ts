
// see src-rust/magic_square/ui_manifest.rs
export enum WasmInputId {
  // there should exist a class variable by the left-hand name for each of these
  drawPattern = "magic_square_input_draw_pattern",

  // color
  colorDirection = "magic_square_input_color_direction",
  colorMode = "magic_square_input_color_mode",
  colorSpeed = "magic_square_input_color_speed",
  color1 = "magic_square_input_color_1",
  color2 = "magic_square_input_color_2",
  color3 = "magic_square_input_color_3",
  color4 = "magic_square_input_color_4",
  color5 = "magic_square_input_color_5",
  color6 = "magic_square_input_color_6",
  color7 = "magic_square_input_color_7",
  color8 = "magic_square_input_color_8",

  // geometry
  radiusBase = "magic_square_input_radius_base",
  radiusStep = "magic_square_input_radius_step",
  transformOrder = "magic_square_input_transform_order",

  // lfo1
  lfo1Active= "magic_square_input_lfo_1_active",
  lfo1Amp = "magic_square_input_lfo_1_amp",
  lfo1Dest = "magic_square_input_lfo_1_dest",
  lfo1Freq = "magic_square_input_lfo_1_freq",
  lfo1Phase = "magic_square_input_lfo_1_phase",
  lfo1Shape = "magic_square_input_lfo_1_shape",

  // lfo2
  lfo2Active= "magic_square_input_lfo_2_active",
  lfo2Amp = "magic_square_input_lfo_2_amp",
  lfo2Dest = "magic_square_input_lfo_2_dest",
  lfo2Freq = "magic_square_input_lfo_2_freq",
  lfo2Phase = "magic_square_input_lfo_2_phase",
  lfo2Shape = "magic_square_input_lfo_2_shape",

  // lfo3
  lfo3Active= "magic_square_input_lfo_3_active",
  lfo3Amp = "magic_square_input_lfo_3_amp",
  lfo3Dest = "magic_square_input_lfo_3_dest",
  lfo3Freq = "magic_square_input_lfo_3_freq",
  lfo3Phase = "magic_square_input_lfo_3_phase",
  lfo3Shape = "magic_square_input_lfo_3_shape",

  // lfo4
  lfo4Active= "magic_square_input_lfo_4_active",
  lfo4Amp = "magic_square_input_lfo_4_amp",
  lfo4Dest = "magic_square_input_lfo_4_dest",
  lfo4Freq = "magic_square_input_lfo_4_freq",
  lfo4Phase = "magic_square_input_lfo_4_phase",
  lfo4Shape = "magic_square_input_lfo_4_shape",

  // rotation
  pitchBase = "magic_square_input_y_rot_base",
  pitchSpread = "magic_square_input_y_rot_spread",
  pitchX = "magic_square_input_x_axis_y_rot_coeff",
  pitchY = "magic_square_input_y_axis_y_rot_coeff",
  rollBase = "magic_square_input_x_rot_base",
  rollSpread = "magic_square_input_x_rot_spread",
  rollX = "magic_square_input_x_axis_x_rot_coeff",
  rollY = "magic_square_input_y_axis_x_rot_coeff",
  yawBase = "magic_square_input_z_rot_base",
  yawSpread = "magic_square_input_z_rot_spread",
  yawX = "magic_square_input_x_axis_z_rot_coeff",
  yawY = "magic_square_input_y_axis_z_rot_coeff",

  // translation
  mouseTracking = "magic_square_input_mouse_tracking",
  translationXBase = "magic_square_input_translation_x_base",
  translationXSpread = "magic_square_input_translation_x_spread",
  translationYBase = "magic_square_input_translation_y_base",
  translationYSpread = "magic_square_input_translation_y_spread",
  translationZBase = "magic_square_input_translation_z_base",
  translationZSpread = "magic_square_input_translation_z_spread"
}
