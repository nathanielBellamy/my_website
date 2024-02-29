import { writable } from 'svelte/store'
import type { Writable } from 'svelte/store'
import type { GithbRepos } from '../integrations/github/GithubTypes'

export const githubRepos: Writable<GithubRepos> = writable([])
