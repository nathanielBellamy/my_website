use js_sys::{DataView, ArrayBuffer};
use wasm_bindgen::JsValue;

use crate::{magic_square::{settings::Settings, main::MagicSquare}, websocket::{Websocket, WebsocketConnError}};

// TODO:
//  run an instance of magic square that has access to the websocket
//  UiBuffer sends serialized JSON (see deser.rs) to backend through socket on update
//  UiBuffer is updated with new received Settings when message recevied

const ARRAY_BUFFER_CAPACITY: u32 = 5000;
const URL: &str = "ws://localhost:8080/public-square-wasm-ws";

pub struct PublicSquare {
    ab: ArrayBuffer,
    pub dv: DataView,
    pub settings: Settings,
    pub websocket: Websocket
}

impl PublicSquare {
    pub fn new(settings: Settings) -> Result<PublicSquare, WebsocketConnError> {
        let ab = ArrayBuffer::new(ARRAY_BUFFER_CAPACITY);
        let websocket = Websocket::new(URL.to_owned())?;
        Ok(PublicSquare {
            dv: DataView::new(&ab, 0, ARRAY_BUFFER_CAPACITY as usize),
            ab,
            settings,
            websocket,
        })
    }

    pub async fn run(settings: JsValue, client_id: JsValue, touch_screen: JsValue) -> Result<(), ()> {
        let clientId: u64 = serde_wasm_bindgen::from_value(client_id).unwrap();
        let touch_screen: bool = serde_wasm_bindgen::from_value(touch_screen).unwrap();
        let settings: Settings = serde_wasm_bindgen::from_value(settings).unwrap();
        
        MagicSquare::run_public();
        

        Ok(())
    }
}
