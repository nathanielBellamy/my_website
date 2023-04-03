use wasm_bindgen::prelude::*;


#[wasm_bindgen]
pub struct AppState {
    pub point: Point
}


#[wasm_bindgen]
impl AppState {
    pub fn new() -> AppState {
        AppState { point: Point::new(0,0)}
    }

    pub fn foo(x: String) -> String {
        format!("WASM: {x}")
    }

    pub fn set_point(&mut self, x: usize, y: usize) -> String {
        let before_x = self.point.x;
        let before_y = self.point.y;
        self.point = Point::new(x,y);
        format!("before_x: {before_x}, before_y: {before_y} after: ({}, {})", self.point.x, self.point.y)
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
