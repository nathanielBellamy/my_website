import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'
import type { GithubStore } from '../integrations/github/GithubTypes'

const defaultStore: GithubStore = {repos: [], userLanguageSummary: []}
export const githubStore: Writable<GithubStore> = writable(defaultStore)
