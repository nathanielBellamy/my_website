use wasm_bindgen::prelude::*;
use give_me_a_sine::GmasWasm;

pub mod magic_banner;
pub mod magic_square;

#[wasm_bindgen]
pub fn init_message(message: String) -> String {
    format!("RUST RUST RUST {message}")
}

#[cfg(test)]
mod tests {

    #[test]
    fn it_works() {
        //
    }
}
