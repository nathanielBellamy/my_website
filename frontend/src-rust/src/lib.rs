#[allow(unused)]
use give_me_a_sine;
use wasm_bindgen::prelude::*; // ensure wasm-bindgen creates bindings for sub-module

#[macro_use]
pub mod magic_square;
pub mod public_square;
pub mod websocket;

#[wasm_bindgen]
pub fn rust_init_message(message: String) {
    log(&format!("Hello, {message}! -From the Rust of RustWasm!"))
}

#[wasm_bindgen]
extern "C" {
    #[wasm_bindgen(js_namespace = console)]
    pub fn log(s: &str);

    #[wasm_bindgen(js_name = "performance")]
    pub static PERFORMANCE: web_sys::Performance;
}

#[cfg(test)]
mod tests {

    #[test]
    fn it_works() {
        //
    }
}
