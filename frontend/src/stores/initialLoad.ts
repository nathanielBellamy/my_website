import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

export const initialLoad: Writable<boolean> = writable(true)
