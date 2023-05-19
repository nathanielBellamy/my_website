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
            draw_pattern: DrawPattern::Seven,

            color_1: [255.0, 0.0, 255.0, 1.0],
            color_2: [0.0, 255.0, 255.0, 1.0],
            color_3: [255.0, 255.0, 1.0, 1.0],
            color_4: [100.0, 1.0, 101.0, 1.0],
            color_5: [0.0, 120.0, 140.0, 1.0],
            color_6: [0.0, 1.0, 1.0, 1.0],
            color_7: [150.0, 140.0, 225.0, 1.0],
            color_8: [110.0, 1.0, 1.0, 1.0],
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

    pub fn max_idx_from_draw_pattern(pattern: DrawPattern) -> usize {
        match pattern {
                DrawPattern::All => 16,
                DrawPattern::One => 1,
                DrawPattern::Two => 2,
                DrawPattern::Three => 3,
                DrawPattern::Four => 4,
                DrawPattern::Five => 5,
                DrawPattern::Six => 6,
                DrawPattern::Seven => 7,
                DrawPattern::Eight => 8,
                DrawPattern::Out1 => 1,
                DrawPattern::Out2 => 2,
                DrawPattern::Out3 => 3,
                DrawPattern::Out4 => 4,
                DrawPattern::Out5 => 5,
                DrawPattern::Out6 => 6,
                DrawPattern::Out7 => 7,
                DrawPattern::Out8 => 8,
                DrawPattern::In1 => 1,
                DrawPattern::In2 => 2,
                DrawPattern::In3 => 3,
                DrawPattern::In4 => 4,
                DrawPattern::In5 => 5,
                DrawPattern::In6 => 6,
                DrawPattern::In7 => 7,
                DrawPattern::In8 => 8,
                DrawPattern::Conv => 4,
                DrawPattern::Div => 4,
                DrawPattern::Random => 3
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

