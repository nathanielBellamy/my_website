use wasm_bindgen::prelude::*;

pub mod magic_banner;
#[macro_use] pub mod magic_square;

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

