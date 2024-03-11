import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'
import {
  type GithubStore,
  SortColumn,
  SortOrder,
} from '../integrations/github/GithubTypes'

const defaultStore: GithubStore = {
  repos: [],
  userLanguageSummary: [],
  sortColumn: SortColumn.PUSHED_AT,
  sortOrder: SortOrder.DESC
}
export const githubStore: Writable<GithubStore> = writable(defaultStore)
