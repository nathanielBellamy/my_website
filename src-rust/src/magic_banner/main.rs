use wasm_bindgen::prelude::*;
use std::cell::RefCell;
use std::rc::Rc;
use crate::magic_banner::buffer::Buffer;
use web_sys::Element;

#[wasm_bindgen]
pub struct MagicBanner {
    buffer: Buffer
}

#[wasm_bindgen]
impl MagicBanner { 
    #[wasm_bindgen(constructor)]
    pub fn new() -> MagicBanner {
        MagicBanner { buffer: Buffer::new() }
    }

    pub fn write_to_buffer(&mut self, x: i32, y: i32) {
        self.buffer.write(x, y);
        self.read_buffer();
    }

    pub fn run(&self) -> Result<(), JsValue> {
        // https://rustwasm.github.io/wasm-bindgen/examples/request-animation-frame.html
        // Here we want to call `requestAnimationFrame` in a loop, but only a fixed
        // number of times. After it's done we want all our resources cleaned up. To
        // achieve this we're using an `Rc`. The `Rc` will eventually store the
        // closure we want to execute on each frame, but to start out it contains
        // `None`.
        //
        // After the `Rc` is made we'll actually create the closure, and the closure
        // will reference one of the `Rc` instances. The other `Rc` reference is
        // used to store the closure, request the first frame, and then is dropped
        // by this function.
        //
        // Inside the closure we've got a persistent `Rc` reference, which we use
        // for all future iterations of the loop
        let f = Rc::new(RefCell::new(None));
        let g = f.clone();

        let mut i = 0;
        *g.borrow_mut() = Some(Closure::new(move || {
            if i > 300 {

                // Drop our handle to this closure so that it will get cleaned
                // up once we return.
                let _ = f.borrow_mut().take();
                return;
            }

            // Set the body's text content to how many times this
            // requestAnimationFrame callback has fired.
            self.buffer;
            i += 1;
            // let text = format!("requestAnimationFrame has been called {} times.", i);
            // self.canvas().set_text_content(Some(&text));

            // Schedule ourself for another requestAnimationFrame callback.
            // request_animation_frame(f.borrow().as_ref().unwrap());
        }));

        // request_animation_frame(g.borrow().as_ref().unwrap());
        Ok(())
    }
}

impl MagicBanner {
    pub fn read_buffer(&self) {
        if self.buffer.idx == 7 {
            // batch process
        }
    }

    pub fn window(&self) -> web_sys::Window {
        web_sys::window().expect("no global `window` exists")
    }

    pub fn document(&self) -> web_sys::Document {
        self.window()
            .document()
            .expect("should have a document on window")
    }

    pub fn canvas(&self) -> Element {
        self.document().get_element_by_id("canvas").expect("unable to find canvas element")
    }
}


use wasm_bindgen::prelude::*;

#[wasm_bindgen]
extern "C" {
    fn requestAnimationFrame(closure: &Closure<dyn FnMut()>) -> u32;
    fn cancelAnimationFrame(id: u32);

    #[wasm_bindgen(js_namespace = console)]
    fn log(s: &str);
}

#[wasm_bindgen]
pub struct AnimationFrameHandle {
    animation_id: u32,
    _closure: Closure<dyn FnMut()>,
}

impl Drop for AnimationFrameHandle {
    fn drop(&mut self) {
        cancelAnimationFrame(self.animation_id);
    }
}

// A type that will log a message when it is dropped.
struct LogOnDrop(&'static str);
impl Drop for LogOnDrop {
    fn drop(&mut self) {
        log(self.0);
    }
}

#[wasm_bindgen]
pub fn run() -> AnimationFrameHandle {
    // We are using `Closure::once` which takes a `FnOnce`, so the function
    // can drop and/or move things that it closes over.
    let fired = LogOnDrop("animation frame fired or canceled");
    let cb = Closure::once(move || drop(fired));

    // Schedule the animation frame!
    let animation_id = requestAnimationFrame(&cb);

    // Again, return a handle to JS, so that the closure is not dropped
    // immediately and JS can decide whether to cancel the animation frame.
    AnimationFrameHandle {
        animation_id,
        _closure: cb,
    }
}





// => keep buffer in RC
// => one event listener has mutable reference to write
// => another closure has the animation loop with an immutable reference
