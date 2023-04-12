use wasm_bindgen::prelude::*;
use crate::magic_buffer::buffer::Buffer;

#[wasm_bindgen]
pub struct MagicBanner {
    buffer: Buffer
}

#[wasm_bindgen]
impl MagicBanner { 
    #[wasm_bindgen(constructor)]
    pub fn new() -> MagicBanner {
        MagicBanner { buffer: Buffer::new() }
    }

    pub fn write_to_buffer(&mut self, x: i32, y: i32) {
        self.buffer.write(x, y);
        self.read_buffer();
    }
}

impl MagicBanner {
    pub fn read_buffer(&self) {
        if self.buffer.idx == 7 {
            // batch process
        }
    }
}

