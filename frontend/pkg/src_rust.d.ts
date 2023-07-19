/* tslint:disable */
/* eslint-disable */
/**
* @param {string} message
*/
export function rust_init_message(message: string): void;
/**
* @param {number} addr
*/
export function worker_entry_point(addr: number): void;
/**
*/
export class GmasWasm {
  free(): void;
/**
* @returns {any}
*/
  static run(): any;
}
/**
*/
export class MagicSquare {
  free(): void;
/**
* @param {any} settings
* @param {any} presets
* @param {any} touch_screen
* @returns {Promise<any>}
*/
  static run(settings: any, presets: any, touch_screen: any): Promise<any>;
}

export type InitInput = RequestInfo | URL | Response | BufferSource | WebAssembly.Module;

export interface InitOutput {
  readonly memory: WebAssembly.Memory;
  readonly magicsquare_run: (a: number, b: number, c: number) => number;
  readonly __wbg_magicsquare_free: (a: number) => void;
  readonly rust_init_message: (a: number, b: number) => void;
  readonly worker_entry_point: (a: number) => void;
  readonly gmaswasm_run: () => number;
  readonly __wbg_gmaswasm_free: (a: number) => void;
  readonly __wbindgen_malloc: (a: number) => number;
  readonly __wbindgen_realloc: (a: number, b: number, c: number) => number;
  readonly __wbindgen_export_2: WebAssembly.Table;
  readonly wasm_bindgen__convert__closures__invoke1_mut__h3f969557b3fb004c: (a: number, b: number, c: number) => void;
  readonly wasm_bindgen__convert__closures__invoke0_mut__hc4b9a43a165e4945: (a: number, b: number) => void;
  readonly _dyn_core__ops__function__FnMut__A____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__h232fbe164ebb195a: (a: number, b: number, c: number) => void;
  readonly _dyn_core__ops__function__FnMut__A____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__ha8c8ec19f24713a8: (a: number, b: number, c: number) => void;
  readonly __wbindgen_exn_store: (a: number) => void;
  readonly wasm_bindgen__convert__closures__invoke2_mut__h37cfe3e274912e80: (a: number, b: number, c: number, d: number) => void;
}

export type SyncInitInput = BufferSource | WebAssembly.Module;
/**
* Instantiates the given `module`, which can either be bytes or
* a precompiled `WebAssembly.Module`.
*
* @param {SyncInitInput} module
*
* @returns {InitOutput}
*/
export function initSync(module: SyncInitInput): InitOutput;

/**
* If `module_or_path` is {RequestInfo} or {URL}, makes a request and
* for everything else, calls `WebAssembly.instantiate` directly.
*
* @param {InitInput | Promise<InitInput>} module_or_path
*
* @returns {Promise<InitOutput>}
*/
export default function __wbg_init (module_or_path?: InitInput | Promise<InitInput>): Promise<InitOutput>;
