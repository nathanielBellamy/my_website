import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

// these settings are non-input settings
// a.k.a. settings that do NOT appear in UiBuffer.settings in RustWasm
export interface MsStoreSettings {
  msColorCurrIdx: number,
  msColorIdxA: number,
  msColorIdxB: number,
  msGeometryIdxA: number,
  msGeometryIdxB: number,
  msGeometryShapeIdx: number,
  msPresetBank: number
  psColorCurrIdx: number,
  psColorIdxA: number,
  psColorIdxB: number,
  psGeometryIdxA: number,
  psGeometryIdxB: number,
  psGeometryShapeIdx: number,
  psPresetBank: number

}

const defaultMsStoreSettings: MsStoreSettings = {
  msColorCurrIdx: 0,
  msColorIdxA: 0,
  msColorIdxB: 15,
  msGeometryIdxA: 0,
  msGeometryIdxB: 15,
  msGeometryShapeIdx: 0,
  msPresetBank: 0,
  psColorCurrIdx: 0,
  psColorIdxA: 0,
  psColorIdxB: 15,
  psGeometryIdxA: 0,
  psGeometryIdxB: 15,
  psGeometryShapeIdx: 0,
  psPresetBank: 0
}

export const msStoreSettings: Writable<MsStoreSettings> = writable(defaultMsStoreSettings)
