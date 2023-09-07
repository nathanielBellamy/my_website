import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

// TODO:
// - constrol modules subscribe to this store directly
// - this means that these values will be remembered/shared between PS + MS
// - at the moment, control modules are container-agnostic
// - we should figure out how we want to track the container
//  - pass down a prop into the control modules
//  vs.
//  - subscribe to a currentContainer store that flips in the onMount + onDestroy actions
//  of the square containers

// these settings are non-input settings
// a.k.a. settings that do NOT appear in UiBuffer.settings in RustWasm
export interface MsStoreSettings {
  colorIdxA: number,
  colorIdxB: number,
  geometryIdxA: number,
  geometryIdxB: number,
  geometryShapeIdx: number,
  presetBank: number
}

const defaultMsStoreSettings: MsStoreSettings = {
  colorIdxA: 0,
  colorIdxB: 15,
  geometryIdxA: 0,
  geometryIdxB: 15,
  geometryShapeIdx: 0,
  presetBank: 0
}

export const msStoreSettings: Writable<MsStoreSettings> = writable(defaultMsStoreSettings)
