use crate::magic_square::main::Rgba;

#[derive(Clone, Copy)]
pub struct Settings {
    // exchanging f32's with JS is easy 
    pub color_origin_r: f32,
    pub color_origin_g: f32,
    pub color_origin_b: f32,
    pub color_origin_a: f32,
    
    pub color_ne_r: f32,
    pub color_ne_g: f32,
    pub color_ne_b: f32,
    pub color_ne_a: f32,

    pub color_nw_r: f32,
    pub color_nw_g: f32,
    pub color_nw_b: f32,
    pub color_nw_a: f32,

    pub color_se_r: f32,
    pub color_se_g: f32,
    pub color_se_b: f32,
    pub color_se_a: f32,

    pub color_sw_r: f32,
    pub color_sw_g: f32,
    pub color_sw_b: f32,
    pub color_sw_a: f32,


    // x_rot_coeff: f32,
    // y_rot_coeff: f32,
    // z_rot_coeff: f32,

    // // rotation sensitivity
    // x_axis_x_rot_corr: f32,
    // x_axis_y_rot_corr: f32,
    // x_axis_z_rot_corr: f32,

    // y_axis_x_rot_corr: f32,
    // y_axis_y_rot_corr: f32,
    // y_axis_z_rot_corr: f32,


    // color_gradient: f32,

    // radius_init: f32,
    // radius_step: f32,

    // // cache
    // cache_max_idx: usize, // 0..50
    // cache_per: usize,
}

impl Settings {
    pub fn new() -> Settings {
        Settings {
            color_origin_r: 0.5,
            color_origin_g: 0.5,
            color_origin_b: 0.5,
            color_origin_a: 1.0,
            
            color_ne_r: 1.0,
            color_ne_g: 0.0,
            color_ne_b: 0.0,
            color_ne_a: 1.0,

            color_nw_r: 0.0,
            color_nw_g: 1.0,
            color_nw_b: 0.0,
            color_nw_a: 1.0,

            color_se_r: 0.0,
            color_se_g: 0.0,
            color_se_b: 1.0,
            color_se_a: 1.0,

            color_sw_r: 1.0,
            color_sw_g: 0.0,
            color_sw_b: 1.0,
            color_sw_a: 1.0,
        }
    }

    pub fn origin_rgba(&self) -> Rgba {
        [
            self.color_origin_r,
            self.color_origin_g,
            self.color_origin_b,
            self.color_origin_a
        ]
    }

    pub fn origin_rgba_string(&self) -> String {
        format!(
            "rgba({}, {}, {}, {})",
            self.color_origin_r,
            self.color_origin_g,
            self.color_origin_b,
            self.color_origin_a
        )
    }

    pub fn nw_rgba(&self) -> Rgba {
        [
            self.color_nw_r,
            self.color_nw_g,
            self.color_nw_b,
            self.color_nw_a
        ]
    }

    pub fn nw_rgba_string(&self) -> String {
        format!(
            "rgba({}, {}, {}, {})",
            self.color_nw_r,
            self.color_nw_g,
            self.color_nw_b,
            self.color_nw_a
        )
    }

    pub fn ne_rgba(&self) -> Rgba {
        [
            self.color_ne_r,
            self.color_ne_g,
            self.color_ne_b,
            self.color_ne_a
        ]
    }

    pub fn ne_rgba_string(&self) -> String {
        format!(
            "rgba({}, {}, {}, {})",
            self.color_ne_r,
            self.color_ne_g,
            self.color_ne_b,
            self.color_ne_a
        )
    }

    pub fn se_rgba(&self) -> Rgba {
        [
            self.color_se_r,
            self.color_se_g,
            self.color_se_b,
            self.color_se_a
        ]
    }

    pub fn se_rgba_string(&self) -> String {
        format!(
            "rgba({}, {}, {}, {})",
            self.color_se_r,
            self.color_se_g,
            self.color_se_b,
            self.color_se_a
        )
    }

    pub fn sw_rgba(&self) -> Rgba {
        [
            self.color_sw_r,
            self.color_sw_g,
            self.color_sw_b,
            self.color_sw_a
        ]
    }

    pub fn sw_rgba_string(&self) -> String {
        format!(
            "rgba({}, {}, {}, {})",
            self.color_sw_r,
            self.color_sw_g,
            self.color_sw_b,
            self.color_sw_a
        )
    }
}
