use wasm_bindgen::prelude::*;
use std::collections::BTreeMap;

#[wasm_bindgen]
pub struct AppBuffer {
    pub x: i32,
    pub y: i32,
    pub r: u8,
    pub g: u8,
    pub b: u8,
}

#[wasm_bindgen]
impl AppBuffer {
    #[wasm_bindgen(constructor)]
    pub fn new() -> AppBuffer {
        AppBuffer {
            x: 0,
            y: 0,
            r: 0,
            g: 0,
            b: 0
        }
    }
}

pub struct Point {
    pub x: i32, 
    pub y: i32
}

impl Point {
    pub fn new(x: i32, y: i32) -> Point {
        Point { x, y }
    }
}

#[derive(Copy, Clone)]
pub struct Rgb {
    pub r: u8,
    pub g: u8,
    pub b: u8,
}

impl Rgb {
    pub fn new(r: u8, g: u8, b: u8) -> Rgb {
        Rgb { r, g, b }
    }

    pub fn default() -> Rgb {
        Rgb { r: 0, g: 0, b: 0 }
    }
}

pub type PointStore = BTreeMap<u8, Point>;

const HD_WIDTH: usize = 1920;
const HD_RGB_WIDTH: usize = HD_WIDTH * 3;
const HD_HEIGHT: usize = 1080;
const RASTER_LEN: usize = HD_HEIGHT * HD_RGB_WIDTH;

pub type Raster = [u8; RASTER_LEN];
const RASTER_BLANK: Raster = [0_u8; RASTER_LEN];

pub struct AppState {
    pub points: PointStore,
    pub raster: Raster,
}

impl AppState {
    pub fn new() -> AppState {
        AppState { 
            points: PointStore::new(),
            raster: RASTER_BLANK,
        }
    }

    pub fn get_index(x: u16, y: u16) -> u32 {
        0
    }
}

#[cfg(test)]
mod tests {

    #[test]
    fn it_works() {
        //
    }
}
