import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

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
