#[allow(unused)]
use give_me_a_sine;
use wasm_bindgen::prelude::*; // ensure wasm-bindgen creates bindings for sub-module

#[macro_use]
pub mod magic_square;

#[wasm_bindgen]
pub fn init_message(message: String) -> String {
    format!("Hello! {message} From Rust->Wasm!")
}

#[cfg(test)]
mod tests {

    #[test]
    fn it_works() {
        //
    }
}
