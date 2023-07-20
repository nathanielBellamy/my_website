import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'
import type { FeedMessage } from '../PublicSquare/FeedMessage'

export const FEED_LENGTH: number = 926
export const psFeed: Writable<FeedMessage[]> = writable([])
