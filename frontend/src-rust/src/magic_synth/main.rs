use wasm_bindgen::JsValue;

use super::settings::Settings;

pub struct MagicSynth;

impl MagicSynth {
    pub fn run(settings: JsValue) -> JsValue {
        let settings: Settings = serde_wasm_bindgen::from_value(settings)
            .unwrap_or(Settings::new());

        let to_js = serde_wasm_bindgen::to_value(&settings).unwrap();
        to_js
    }
}
