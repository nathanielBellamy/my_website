use serde::{Serialize, Deserialize};

#[derive(Clone, Copy, Debug, Serialize, Deserialize)]
pub struct Settings;

impl Settings {
    pub fn new() -> Settings {
        Settings
    }
}
