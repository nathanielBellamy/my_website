import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'

export enum SquareType {
  magic = "magic",
  public = "public",
  none = "none"
}

export const currSquare: Writable<SquareType> = writable(SquareType.none)
