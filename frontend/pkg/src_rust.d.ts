/* tslint:disable */
/* eslint-disable */

export class GmasWasm {
    private constructor();
    free(): void;
    [Symbol.dispose](): void;
    static run(): any;
}

export class MagicSquare {
    private constructor();
    free(): void;
    [Symbol.dispose](): void;
    static run(settings: any, presets: any, touch_screen: any): Promise<any>;
}

export class PubSq {
    private constructor();
    free(): void;
    [Symbol.dispose](): void;
    static run(base_url: any, set_all_settings: Function, touch_screen: any): Promise<any>;
}

export function rust_init_message(message: string): void;

export function worker_entry_point(addr: number): void;

export type InitInput = RequestInfo | URL | Response | BufferSource | WebAssembly.Module;

export interface InitOutput {
    readonly memory: WebAssembly.Memory;
    readonly __wbg_magicsquare_free: (a: number, b: number) => void;
    readonly magicsquare_run: (a: any, b: any, c: any) => any;
    readonly worker_entry_point: (a: number) => void;
    readonly __wbg_pubsq_free: (a: number, b: number) => void;
    readonly pubsq_run: (a: any, b: any, c: any) => any;
    readonly rust_init_message: (a: number, b: number) => void;
    readonly __wbg_gmaswasm_free: (a: number, b: number) => void;
    readonly gmaswasm_run: () => any;
    readonly wasm_bindgen__closure__destroy__h234764b74a55fcd8: (a: number, b: number) => void;
    readonly wasm_bindgen__closure__destroy__hf43dbe93a8793e3a: (a: number, b: number) => void;
    readonly wasm_bindgen__closure__destroy__h153965bb95497284: (a: number, b: number) => void;
    readonly wasm_bindgen__convert__closures_____invoke__hc2dfe96f5f250965: (a: number, b: number, c: number, d: number) => void;
    readonly wasm_bindgen__convert__closures_____invoke__h6999a6c48ba9ba5b: (a: number, b: number, c: any, d: any) => void;
    readonly wasm_bindgen__convert__closures_____invoke__h4da8c97a0c5a4e2e: (a: number, b: number, c: any) => void;
    readonly wasm_bindgen__convert__closures_____invoke__hd95d24573c051714: (a: number, b: number, c: any) => void;
    readonly wasm_bindgen__convert__closures_____invoke__he183312e5a13b41c: (a: number, b: number, c: any) => void;
    readonly wasm_bindgen__convert__closures_____invoke__h5994bde47de2401f: (a: number, b: number) => void;
    readonly __wbindgen_malloc: (a: number, b: number) => number;
    readonly __wbindgen_realloc: (a: number, b: number, c: number, d: number) => number;
    readonly __wbindgen_exn_store: (a: number) => void;
    readonly __externref_table_alloc: () => number;
    readonly __wbindgen_externrefs: WebAssembly.Table;
    readonly __wbindgen_start: () => void;
}

export type SyncInitInput = BufferSource | WebAssembly.Module;

/**
 * Instantiates the given `module`, which can either be bytes or
 * a precompiled `WebAssembly.Module`.
 *
 * @param {{ module: SyncInitInput }} module - Passing `SyncInitInput` directly is deprecated.
 *
 * @returns {InitOutput}
 */
export function initSync(module: { module: SyncInitInput } | SyncInitInput): InitOutput;

/**
 * If `module_or_path` is {RequestInfo} or {URL}, makes a request and
 * for everything else, calls `WebAssembly.instantiate` directly.
 *
 * @param {{ module_or_path: InitInput | Promise<InitInput> }} module_or_path - Passing `InitInput` directly is deprecated.
 *
 * @returns {Promise<InitOutput>}
 */
export default function __wbg_init (module_or_path?: { module_or_path: InitInput | Promise<InitInput> } | InitInput | Promise<InitInput>): Promise<InitOutput>;
