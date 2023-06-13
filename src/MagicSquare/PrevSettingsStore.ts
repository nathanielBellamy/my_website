import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'
import type { StorageSettings } from './StorageSettings'

export const prevSettingsStore: Writable<StorageSettings | null> = writable(null)
