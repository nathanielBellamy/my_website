import reposFixture from './reposFixture.json'
import {
  SortOrder,
  SortColumn,
  type ColorData,
  type GithubStore,
  type GithubRepo,
  DATE_SORT_COLUMNS,
  LOWERCASE_SORT_COLUMNS
} from './GithubTypes'
import { githubStore } from "../../stores/githubStore"

// TODO:
// - rename GithubStoreHandler
// - set reference to correct GithubStore
//   - githubStoreNS
//   - githubStoreRando
export default class GithubIntegration {
  setSortOrder(sortOrder: SortOrder): void {
    githubStore.update((store: GithubStore) => ({...store, sortOrder}))
    this.sortReposBy()
  }

  swapSortOrder(): void {
    githubStore.update((store: GithubStore) => {
      const sortOrder: SortOrder =  store.sortOrder === SortOrder.DESC ? SortOrder.ASC : SortOrder.DESC
      return {
        ...store,
        sortOrder
      }
    })
    this.sortReposBy()
  }

  sortReposBy(col: SortColumn|null = null): void {
    githubStore.update((store: GithubStore) => {
      let sortColumn: SortColumn = col ? col : store.sortColumn

      store.repos.sort((x, y) => this.sortFunc(x, y, sortColumn, store.sortOrder))
      return {...store, sortColumn}
    })
  }

  sortFunc(x: GithubRepo, y: GithubRepo, sortColumn: SortColumn, sortOrder: SortOrder) {
    let lessThanReturnValue: number = sortOrder === SortOrder.ASC ? 1 : -1
    if (DATE_SORT_COLUMNS.includes(sortColumn)) lessThanReturnValue *= -1
    let grtrThanReturnValue: number = -1 * lessThanReturnValue

    let xVal = x[sortColumn]
    let yVal = y[sortColumn]
    if (LOWERCASE_SORT_COLUMNS.includes(sortColumn))
    {
      xVal = xVal.toLowerCase()
      yVal = yVal.toLowerCase()
    }
    if (xVal < yVal) return lessThanReturnValue
    if (xVal > yVal) return grtrThanReturnValue
    return 0
  }

  async fetchRepos(): Promise {
    // TODO: if reposReady, do not make call
    return await fetch("api/github/repos")
      .then((resp) => resp.json())
      .then(async (resp) => {
        githubStore.update((store: GithubStore) => {
          return {
            ...store,
            repos: this.mapRepos(resp.repos),
            reposReady: true,
            userLanguageSummary: resp.user_language_summary
          }
        })

        this.sortReposBy()
      })
      .catch(() => {
        // defaut to snapshot if any trouble happens along the way
        const parsedFixture: GithubStore = reposFixture
        const repos = this.mapRepos(parsedFixture.repos)
        githubStore.update((store: GithubStore) => {
          return {
            ...store,
            repos,
            reposReady: true,
            userLanguageSummary: parsedFixture.user_language_summary
          }
        })
        this.sortReposBy()
      })
  }

  mapRepos(repos: any) {
    return repos.map(repo => {
      return {
        colorData: repo.color_data,
        commitData: repo.commit_data,
        created_at: new Date(repo.created_at),
        description: repo.description,
        html_url: repo.html_url,
        language: repo.language,
        languageData: repo.language_data,
        name: repo.name,
        pushed_at: new Date(repo.pushed_at),
        updated_at: new Date(repo.updated_at)
      }
    })
  }
}
