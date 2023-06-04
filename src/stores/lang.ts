import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'
import { Lang } from '../I18n'

export const lang: Writable<Lang | null> = writable(null)
