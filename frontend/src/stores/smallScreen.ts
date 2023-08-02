import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

export const smallScreen: Writable<boolean | null> = writable(null)
