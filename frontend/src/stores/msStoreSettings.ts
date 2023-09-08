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
  msColorIdxA: number,
  msColorIdxB: number,
  msGeometryIdxA: number,
  msGeometryIdxB: number,
  msGeometryShapeIdx: number,
  msPresetBank: number
  psColorIdxA: number,
  psColorIdxB: number,
  psGeometryIdxA: number,
  psGeometryIdxB: number,
  psGeometryShapeIdx: number,
  psPresetBank: number

}

const defaultMsStoreSettings: MsStoreSettings = {
  msColorIdxA: 0,
  msColorIdxB: 15,
  msGeometryIdxA: 0,
  msGeometryIdxB: 15,
  msGeometryShapeIdx: 0,
  msPresetBank: 0,
  psColorIdxA: 0,
  psColorIdxB: 15,
  psGeometryIdxA: 0,
  psGeometryIdxB: 15,
  psGeometryShapeIdx: 0,
  psPresetBank: 0
}

export const msStoreSettings: Writable<MsStoreSettings> = writable(defaultMsStoreSettings)
