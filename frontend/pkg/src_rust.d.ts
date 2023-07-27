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
/**
*/
export class PubSq {
  free(): void;
/**
* @param {Function} set_all_settings
* @param {any} touch_screen
* @returns {Promise<any>}
*/
  static run(set_all_settings: Function, touch_screen: any): Promise<any>;
}

export type InitInput = RequestInfo | URL | Response | BufferSource | WebAssembly.Module;

export interface InitOutput {
  readonly memory: WebAssembly.Memory;
  readonly magicsquare_run: (a: number, b: number, c: number) => number;
  readonly __wbg_magicsquare_free: (a: number) => void;
  readonly pubsq_run: (a: number, b: number) => number;
  readonly rust_init_message: (a: number, b: number) => void;
  readonly __wbg_pubsq_free: (a: number) => void;
  readonly worker_entry_point: (a: number) => void;
  readonly gmaswasm_run: () => number;
  readonly __wbg_gmaswasm_free: (a: number) => void;
  readonly __wbindgen_malloc: (a: number, b: number) => number;
  readonly __wbindgen_realloc: (a: number, b: number, c: number, d: number) => number;
  readonly __wbindgen_export_2: WebAssembly.Table;
  readonly _dyn_core__ops__function__FnMut_____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__hf3ae5327e736add0: (a: number, b: number) => void;
  readonly _dyn_core__ops__function__FnMut__A____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__h41ef1178afbc5db3: (a: number, b: number, c: number) => void;
  readonly _dyn_core__ops__function__FnMut__A____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__h578b14ae1e225664: (a: number, b: number, c: number, d: number) => void;
  readonly _dyn_core__ops__function__FnMut__A____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__h9917f84f5e50eb0c: (a: number, b: number, c: number) => void;
  readonly _dyn_core__ops__function__FnMut__A____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__h8160dcae764b395b: (a: number, b: number, c: number) => void;
  readonly __wbindgen_exn_store: (a: number) => void;
  readonly wasm_bindgen__convert__closures__invoke2_mut__h8a993d1da9b2edc3: (a: number, b: number, c: number, d: number) => void;
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
