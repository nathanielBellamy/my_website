use js_sys::ArrayBuffer;
use serde::{Deserialize, Serialize};
use wasm_bindgen::{JsCast, JsValue};
use web_sys::{BinaryType, MessageEvent};

#[derive(Debug)]
pub struct Websocket {
    pub conn: web_sys::WebSocket
}

#[derive(Clone, Copy, Debug, Deserialize, Serialize)]
pub struct WebsocketConnError;

#[derive(Clone, Copy, Debug, Deserialize, Serialize)]
pub struct WebsocketSendError;

#[derive(Clone, Copy, Debug, Deserialize, Serialize)]
pub enum WebsocketError {
    WebsocketConnError,
    WebsocketSendError
}

impl Websocket {
    pub fn new(url: String) -> Result<Websocket, WebsocketConnError>  {
        match web_sys::WebSocket::new(&url) {
            Ok(conn) => {
                conn.set_binary_type(BinaryType::Blob);
                Ok(Websocket { conn })
            },
            Err(_) => Err(WebsocketConnError)
        }
    }

    pub fn clone(&self) -> Websocket {
        Websocket { conn: self.conn.clone() }
    }

    pub fn send(&self, message: String) -> Result<(), WebsocketSendError> {
        let message_event = MessageEvent::new(&message).unwrap();
        match message_event.data().dyn_into::<ArrayBuffer>() {
            Ok(arr_buff) => {
                self.conn.send_with_array_buffer(&arr_buff).unwrap();
                return Ok(())
            },
            Err(_) => Err(WebsocketSendError)
        }
    }

    pub fn close(&self) -> Result<(), JsValue> {
        self.conn.close()
    }
}

