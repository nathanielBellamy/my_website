

pub struct Settings {
    x_rot_coeff: f32,
    y_rot_coeff: f32,
    z_rot_coeff: f32,

    // rotation sensitivity
    x_axis_x_rot_corr: f32,
    x_axis_y_rot_corr: f32,
    x_axis_z_rot_corr: f32,

    y_axis_x_rot_corr: f32,
    y_axis_y_rot_corr: f32,
    y_axis_z_rot_corr: f32,

    color_orig: f32,
    color_ne: f32,
    color_nw: f32,
    color_se: f32,
    color_sw: f32,

    color_gradient: f32,

    radius_init: f32,
    radius_step: f32,

    // cache
    cache_max_idx: usize, // 0..50
    cache_per: usize,
}
