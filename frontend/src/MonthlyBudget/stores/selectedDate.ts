import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

export const selectedDate: Writable<Date> = writable(new Date())
