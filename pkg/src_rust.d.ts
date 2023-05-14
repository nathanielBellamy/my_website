declare namespace wasm_bindgen {
	/* tslint:disable */
	/* eslint-disable */
	/**
	* @param {string} message
	* @returns {string}
	*/
	export function init_message(message: string): string;
	/**
	* @param {number} addr
	*/
	export function worker_entry_point(addr: number): void;
	/**
	*/
	export class GmasWasm {
	  free(): void;
	/**
	*/
	  constructor();
	}
	/**
	*/
	export class MagicBanner {
	  free(): void;
	/**
	*/
	  static run(): void;
	}
	/**
	*/
	export class MagicSquare {
	  free(): void;
	/**
	*/
	  static run(): void;
	}
	/**
	*/
	export class Wasm {
	  free(): void;
	/**
	*/
	  static run(): void;
	}
	
}

declare type InitInput = RequestInfo | URL | Response | BufferSource | WebAssembly.Module;

declare interface InitOutput {
  readonly magicsquare_run: (a: number) => void;
  readonly __wbg_magicsquare_free: (a: number) => void;
  readonly magicbanner_run: (a: number) => void;
  readonly __wbg_magicbanner_free: (a: number) => void;
  readonly init_message: (a: number, b: number, c: number) => void;
  readonly worker_entry_point: (a: number) => void;
  readonly gmaswasm_new: () => void;
  readonly __wbg_gmaswasm_free: (a: number) => void;
  readonly wasm_run: () => void;
  readonly __wbg_wasm_free: (a: number) => void;
  readonly memory: WebAssembly.Memory;
  readonly __wbindgen_malloc: (a: number) => number;
  readonly __wbindgen_realloc: (a: number, b: number, c: number) => number;
  readonly __wbindgen_export_3: WebAssembly.Table;
  readonly _dyn_core__ops__function__FnMut__A____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__hf243b7ee20a6e774: (a: number, b: number, c: number) => void;
  readonly _dyn_core__ops__function__FnMut__A____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__h50920110bf6927a2: (a: number, b: number, c: number) => void;
  readonly __wbindgen_add_to_stack_pointer: (a: number) => number;
  readonly __wbindgen_free: (a: number, b: number) => void;
  readonly __wbindgen_exn_store: (a: number) => void;
  readonly __wbindgen_thread_destroy: (a: number, b: number) => void;
  readonly __wbindgen_start: () => void;
}

/**
* If `module_or_path` is {RequestInfo} or {URL}, makes a request and
* for everything else, calls `WebAssembly.instantiate` directly.
*
* @param {InitInput | Promise<InitInput>} module_or_path
* @param {WebAssembly.Memory} maybe_memory
*
* @returns {Promise<InitOutput>}
*/
declare function wasm_bindgen (module_or_path?: InitInput | Promise<InitInput>, maybe_memory?: WebAssembly.Memory): Promise<InitOutput>;
