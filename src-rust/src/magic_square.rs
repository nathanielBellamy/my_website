use wasm_bindgen::prelude::*;

// Animator exposes the buffer to JS
// JS writes to the buffer
// Animator reads from the buffer and renders to canvas

#[wasm_bindgen]
pub struct MagicSquare {
    buffer: MagicSquareBuffer
}

#[wasm_bindgen]
impl MagicSquare { 
    #[wasm_bindgen(constructor)]
    pub fn new() -> MagicSquare {
        MagicSquare { buffer: MagicSquareBuffer::new() }
    }

    pub fn write_to_buffer(&mut self, x: i32, y: i32) -> String {
        let test_string = self.buffer.write(x, y);
        self.read_buffer();
        test_string
    }
}

impl MagicSquare {
    pub fn read_buffer(&self) {
        if self.buffer.idx == 7 {
            // batch process
        }
    }
}

pub struct MagicSquareBuffer {
    pub x_0: i32,
    pub x_1: i32,
    pub x_2: i32,
    pub x_3: i32,
    pub x_4: i32,
    pub x_5: i32,
    pub x_6: i32,
    pub x_7: i32,
    pub y_0: i32,
    pub y_1: i32,
    pub y_2: i32,
    pub y_3: i32,
    pub y_4: i32,
    pub y_5: i32,
    pub y_6: i32,
    pub y_7: i32,
    pub idx: u8,
}

impl MagicSquareBuffer {
    pub fn new() -> MagicSquareBuffer {
        MagicSquareBuffer {
            x_0: -1,
            x_1: -1,
            x_2: -1,
            x_3: -1,
            x_4: -1,
            x_5: -1,
            x_6: -1,
            x_7: -1,
            y_0: -1,
            y_1: -1,
            y_2: -1,
            y_3: -1,
            y_4: -1,
            y_5: -1,
            y_6: -1,
            y_7: -1,
            idx: 0,
        }
    }

    pub fn write(&mut self, x: i32, y: i32) -> String {
        let idx = self.idx;
        match idx {
            0 => {
                self.x_0 = x;
                self.y_0 = y;
            },
            1 => {
                self.x_1 = x;
                self.y_1 = y;
            },
            2 => {
                self.x_2 = x;
                self.y_2 = y;
            },
            3 => {
                self.x_3 = x;
                self.y_3 = y;
            },
            4 => {
                self.x_4 = x;
                self.y_4 = y;
            },
            5 => {
                self.x_5 = x;
                self.y_5 = y;
            },
            6 => {
                self.x_6 = x;
                self.y_6 = y;
            },
            7 => {
                self.x_7 = x;
                self.y_7 = y;
            },
            _ => (),
        }
        self.idx = (idx + 1) % 8;
        format!(
            "0.({}, {}), 1.({}, {}), 2.({}, {}), 3.({}, {}), 4.({}, {}), 5.({}, {}), 6.({}, {}), 7.({}, {})",
            self.x_0,
            self.y_0,
            self.x_1,
            self.y_1,
            self.x_2,
            self.y_2,
            self.x_3,
            self.y_3,
            self.x_4,
            self.y_4,
            self.x_5,
            self.y_5,
            self.x_6,
            self.y_6,
            self.x_7,
            self.y_7,
        )
    }
}


