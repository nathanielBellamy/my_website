import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'
import type { StorageSettings } from '../MagicSquare/StorageSettings'

export const magicSquareSettings: Writable<StorageSettings | null> = writable(null)
