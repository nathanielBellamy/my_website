//
//  handle async wasm import once
//  and for all 
//
import("./main.ts")
  .catch((e) => console.error("Error importing src/main.ts", e))
