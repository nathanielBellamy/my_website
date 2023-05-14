

import("./main.ts")
  .then(() => wasm_bindgen())
  .catch((e) => console.error("Error importing src/main.ts", e))
