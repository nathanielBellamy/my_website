declare namespace wasm_bindgen {
	/* tslint:disable */
	/* eslint-disable */
	/**
	* @param {number} addr
	*/
	export function worker_entry_point(addr: number): void;
	/**
	* @param {string} message
	* @returns {string}
	*/
	export function init_message(message: string): string;
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
	
}

declare type InitInput = RequestInfo | URL | Response | BufferSource | WebAssembly.Module;

declare interface InitOutput {
  readonly memory: WebAssembly.Memory;
  readonly magicsquare_run: (a: number) => void;
  readonly __wbg_magicsquare_free: (a: number) => void;
  readonly magicbanner_run: (a: number) => void;
  readonly __wbg_magicbanner_free: (a: number) => void;
  readonly worker_entry_point: (a: number) => void;
  readonly init_message: (a: number, b: number, c: number) => void;
  readonly __wbindgen_malloc: (a: number) => number;
  readonly __wbindgen_realloc: (a: number, b: number, c: number) => number;
  readonly __wbindgen_export_2: WebAssembly.Table;
  readonly _dyn_core__ops__function__FnMut__A____Output___R_as_wasm_bindgen__closure__WasmClosure___describe__invoke__h0c996e74fda0591a: (a: number, b: number, c: number) => void;
  readonly __wbindgen_add_to_stack_pointer: (a: number) => number;
  readonly __wbindgen_free: (a: number, b: number) => void;
  readonly __wbindgen_exn_store: (a: number) => void;
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
