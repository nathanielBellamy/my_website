import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

export enum NewOrLoad {
  load = "load",
  new = "new",
  uninit = "uninit",
}

export const newOrLoad: Writable<NewOrLoad> = writable(NewOrLoad.uninit)
