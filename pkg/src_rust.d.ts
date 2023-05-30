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
	* @returns {any}
	*/
	  static run(): any;
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
	* @returns {Promise<void>}
	*/
	  static run(): Promise<void>;
	}
	
}

declare type InitInput = RequestInfo | URL | Response | BufferSource | WebAssembly.Module;

declare interface InitOutput {
  readonly memory: WebAssembly.Memory;
  readonly init_message: (a: number, b: number, c: number) => void;
  readonly magicsquare_run: () => number;
  readonly __wbg_magicsquare_free: (a: number) => void;
  readonly magicbanner_run: (a: number) => void;
  readonly __wbg_magicbanner_free: (a: number) => void;
  readonly worker_entry_point: (a: number) => void;
  readonly gmaswasm_run: () => number;
  readonly __wbg_gmaswasm_free: (a: number) => void;
  readonly __wbindgen_malloc: (a: number) => number;
  readonly __wbindgen_realloc: (a: number, b: number, c: number) => number;
  readonly __wbindgen_export_2: WebAssembly.Table;
  readonly _dyn_core__ops__function__FnMut__A____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__h4b856b0baa9ad843: (a: number, b: number, c: number) => void;
  readonly _dyn_core__ops__function__FnMut_____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__h846c1ff79c7a8c07: (a: number, b: number) => void;
  readonly _dyn_core__ops__function__FnMut__A____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__h232fbe164ebb195a: (a: number, b: number, c: number) => void;
  readonly _dyn_core__ops__function__FnMut__A____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__h8c076a919d0d8cfc: (a: number, b: number, c: number) => void;
  readonly __wbindgen_add_to_stack_pointer: (a: number) => number;
  readonly __wbindgen_free: (a: number, b: number) => void;
  readonly __wbindgen_exn_store: (a: number) => void;
  readonly wasm_bindgen__convert__closures__invoke2_mut__h37cfe3e274912e80: (a: number, b: number, c: number, d: number) => void;
}

/**
* If `module_or_path` is {RequestInfo} or {URL}, makes a request and
* for everything else, calls `WebAssembly.instantiate` directly.
*
* @param {InitInput | Promise<InitInput>} module_or_path
*
* @returns {Promise<InitOutput>}
*/
declare function wasm_bindgen (module_or_path?: InitInput | Promise<InitInput>): Promise<InitOutput>;
