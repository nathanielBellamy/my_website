import type { ColorDirection, ColorMode } from './ControlModules/Color'
import type { DrawPattern } from "./ControlModules/DrawPattern"
import type { LfoDestination } from "./ControlModules/LfoDestination"
import type { LfoShape } from "./ControlModules/LfoShape"
import type { MouseTracking } from "./ControlModules/MouseTracking"
import type { TransformOrder } from "./ControlModules/TransformOrder"

export interface StorageSettings {
  // COLOR
  color_direction: ColorDirection,
  color_mode: ColorMode,
  color_1: number[],
  color_2: number[],
  color_3: number[],
  color_4: number[],
  color_5: number[],
  color_6: number[],
  color_7: number[],
  color_8: number[],

  // GEOMETRY
  radius_base: number,
  radius_step: number,
  transform_order: TransformOrder

  // lfo_1
  lfo_1_active: boolean,
  lfo_1_amp: number,
  lfo_1_dest: LfoDestination,
  lfo_1_freq: number,
  lfo_1_phase: number,
  lfo_1_shape: LfoShape,

  // lfo_2
  lfo_2_active: boolean,
  lfo_2_amp: number,
  lfo_2_dest: LfoDestination,
  lfo_2_freq: number,
  lfo_2_phase: number,
  lfo_2_shape: LfoShape,

  // lfo_3
  lfo_3_active: boolean,
  lfo_3_amp: number,
  lfo_3_dest: LfoDestination,
  lfo_3_freq: number,
  lfo_3_phase: number,
  lfo_3_shape: LfoShape,

  // lfo_4
  lfo_4_active: boolean,
  lfo_4_amp: number,
  lfo_4_dest: LfoDestination,
  lfo_4_freq: number,
  lfo_4_phase: number,
  lfo_4_shape: LfoShape,

  // PATTERN
  draw_pattern: DrawPattern,


  // ROTATION
  x_rot_base: number,
  y_rot_base: number,
  z_rot_base: number,

  x_rot_spread: number,
  y_rot_spread: number,
  z_rot_spread: number,

  // rotation sensitivity to mouse movement
  x_axis_x_rot_coeff: number,
  x_axis_y_rot_coeff: number,
  x_axis_z_rot_coeff: number,

  y_axis_x_rot_coeff: number,
  y_axis_y_rot_coeff: number,
  y_axis_z_rot_coeff: number,

  // TRANSLATION
  translation_x_base: number,
  translation_x_spread: number,
  translation_y_base: number,
  translation_y_spread: number,
  translation_z_base: number,
  translation_z_spread: number,
  mouse_tracking: MouseTracking,
}
