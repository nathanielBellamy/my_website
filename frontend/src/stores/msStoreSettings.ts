import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

// these settings are non-input settings
// a.k.a. settings that do NOT appear in UiBuffer.settings in RustWasm
export interface MsStoreSettings {
  presetBank: number
}

const defaultMsStoreSettings: MsStoreSettings = {
  presetBank: 0
}

export const msStoreSettings: Writable<MsStoreSettings> = writable(defaultMsStoreSettings)
