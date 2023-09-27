import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

export const psConnected: Writable<boolean> = writable(false)
