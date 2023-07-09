import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

export const touchScreen: Writable<boolean | null> = writable(null)
