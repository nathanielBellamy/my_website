use wasm_bindgen::prelude::*;
use js_sys::Array;
use gloo_utils::format::JsValueSerdeExt;


#[wasm_bindgen]
pub struct AppState {
    pub points: Array,
}


#[wasm_bindgen]
impl AppState {
    pub fn new() -> AppState {
        AppState { points: vec![] }
    }

    pub fn foo(x: String) -> String {
        format!("WASM: {x}")
    }
 
    pub fn add_point(&mut self, x: usize, y: usize) -> JsValue {
        self.points.push(Point::new(x,y));
        JsValue::from_serde(self.points).unwrap()
    }
}

#[wasm_bindgen]
#[derive(Copy, Clone)]
pub struct Point {
    pub x: usize, 
    pub y: usize
}

#[wasm_bindgen]
impl Point {
    #[wasm_bindgen(constructor)]
    pub fn new(x: usize, y: usize) -> Point {
        Point { x, y }
    }
}


#[cfg(test)]
mod tests {

    #[test]
    fn it_works() {
        //
    }
}
