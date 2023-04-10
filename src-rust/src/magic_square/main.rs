use wasm_bindgen::prelude::*;
use crate::magic_square::buffer::Buffer;

#[wasm_bindgen]
pub struct MagicSquare {
    buffer: Buffer
}

#[wasm_bindgen]
impl MagicSquare { 
    #[wasm_bindgen(constructor)]
    pub fn new() -> MagicSquare {
        MagicSquare { buffer: Buffer::new() }
    }

    pub fn write_to_buffer(&mut self, x: i32, y: i32) {
        self.buffer.write(x, y);
        self.read_buffer();
    }
}

impl MagicSquare {
    pub fn read_buffer(&self) {
        if self.buffer.idx == 7 {
            // batch process
        }
    }
}

