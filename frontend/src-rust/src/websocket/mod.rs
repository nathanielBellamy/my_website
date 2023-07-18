use js_sys::ArrayBuffer;
use wasm_bindgen::JsCast;
use web_sys::{BinaryType, MessageEvent};

#[derive(Clone, Debug)]
pub struct Websocket {
    conn: web_sys::WebSocket
}

#[derive(Clone, Copy, Debug)]
pub struct WebsocketConnError;

#[derive(Clone, Copy, Debug)]
pub struct WebsocketSendError;

pub enum WebsocketError {
    WebsocketConnError,
    WebsocketSendError
}

impl Websocket {
    pub fn new(url: String) -> Result<Websocket, WebsocketConnError>  {
        match web_sys::WebSocket::new(&url) {
            Ok(conn) => {
                conn.set_binary_type(BinaryType::Arraybuffer);
                Ok(Websocket { conn })
            },
            Err(_) => Err(WebsocketConnError)
        }
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
}

unsafe fn any_as_u8_slice<T: Sized>(p: &T) -> &[u8] {
    ::core::slice::from_raw_parts(
        (p as *const T) as *const u8,
        ::core::mem::size_of::<T>(),
    )
}

fn main() {
    struct MyStruct {
        id: u8,
        data: [u8; 1024],
    }
    let my_struct = MyStruct { id: 0, data: [1; 1024] };
    let bytes: &[u8] = unsafe { any_as_u8_slice(&my_struct) };
    // tcp_stream.write(bytes);
    println!("{:?}", bytes);
}
