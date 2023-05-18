use crate::magic_square::main::Rgba;

#[derive(Clone, Copy)]
pub enum DrawPattern {
    All,
    One,
    Two,
    Three,
    Four,
    Five,
    Six,
    Seven,
    Eight,
    Out1,
    Out2,
    Out3,
    Out4,
    Out5,
    Out6,
    Out7,
    Out8,
    In1,
    In2,
    In3,
    In4,
    In5,
    In6,
    In7,
    In8,
    Conv,
    Div,
    Random
}

#[derive(Clone, Copy)]
pub struct Settings {
    pub draw_pattern: DrawPattern,

    pub color_1: Rgba,
    pub color_2: Rgba,
    pub color_3: Rgba,
    pub color_4: Rgba,
    pub color_5: Rgba,
    pub color_6: Rgba,
    pub color_7: Rgba,
    pub color_8: Rgba,

    // mouse settings
    // MouseFollow - Always, Click + Drag, DoubleClick On/Off

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
            draw_pattern: DrawPattern::Three,

            color_1: [255.0, 0.0, 255.0, 1.0],
            color_2: [0.0, 1.0, 1.0, 1.0],
            color_3: [0.0, 1.0, 1.0, 1.0],
            color_4: [0.0, 1.0, 1.0, 1.0],
            color_5: [0.0, 1.0, 1.0, 1.0],
            color_6: [0.0, 1.0, 1.0, 1.0],
            color_7: [0.0, 1.0, 1.0, 1.0],
            color_8: [0.0, 1.0, 1.0, 1.0],
        }
    }

    pub fn draw_pattern_from_string(pattern: String) -> DrawPattern {
        match pattern.as_str() {
            "All" => DrawPattern::All,
            "One" => DrawPattern::One,
            "Two" => DrawPattern::Two,
            "Three" => DrawPattern::Three,
            "Four" => DrawPattern::Four,
            "Five" => DrawPattern::Five,
            "Six" => DrawPattern::Six,
            "Seven" => DrawPattern::Seven,
            "Eight" => DrawPattern::Eight,
            "Out1" => DrawPattern::Out1,
            "Out2" => DrawPattern::Out2,
            "Out3" => DrawPattern::Out3,
            "Out4" => DrawPattern::Out4,
            "Out5" => DrawPattern::Out5,
            "Out6" => DrawPattern::Out6,
            "Out7" => DrawPattern::Out7,
            "Out8" => DrawPattern::Out8,
            "In1" => DrawPattern::In1,
            "In2" => DrawPattern::In2,
            "In3" => DrawPattern::In3,
            "In4" => DrawPattern::In4,
            "In5" => DrawPattern::In5,
            "In6" => DrawPattern::In6,
            "In7" => DrawPattern::In7,
            "In8" => DrawPattern::In8,
            "Conv" => DrawPattern::Conv,
            "Div" => DrawPattern::Div,
            "Random" => DrawPattern::Random,
            _ => DrawPattern::Three
        }
    }

    pub fn rgba_string(arr: Rgba) -> String {
        format!(
            "{},{},{},{}",
            arr[0],
            arr[1],
            arr[2],
            arr[3]
        )
    }
}

