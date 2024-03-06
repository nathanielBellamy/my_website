import reposFixture from './reposFixture.json'
import {
  SortOrder,
  SortColumn,
  type ColorData,
  type GithubStore,
  DATE_SORT_COLUMNS,
  LOWERCASE_SORT_COLUMNS
} from './GithubTypes'

export default class GithubIntegration {
  sortOrder: SortOrder = SortOrder.DESC
  sortColumn: SortColumn = SortColumn.PUSHED_AT
  store: Writeable<GithubStore>
  reposReady: boolean = false
  reposVal: GithubRepos
  updateReposReady: (val: boolean) => void

  constructor(
    store: Writeable<GithubStore>,
    updateReposReady: (val: boolean) => void,
  ) {
    this.store = store
    this.store.subscribe((store: GithubStore) => this.reposVal = [...store.repos])
    this.updateReposReady = updateReposReady
  }

  sortReposBy(col: SortColumn|null = null): void {
    if (col)
    {
      this.sortColumn = col
    }
    let lessThanReturnValue: number = this.sortOrder === SortOrder.ASC ? 1 : -1
    if (DATE_SORT_COLUMNS.includes(this.sortColumn)) lessThanReturnValue *= -1

    let grtrThanReturnValue: number = -1 * lessThanReturnValue
    this.store.update((store: GithubStore) => {
      this.reposVal.sort((x: any, y: any) => {
        let xVal = x[this.sortColumn]
        let yVal = y[this.sortColumn]
        if (LOWERCASE_SORT_COLUMNS.includes(this.sortColumn))
        {
          xVal = xVal.toLowerCase()
          yVal = yVal.toLowerCase()
        }
        if (xVal < yVal) return lessThanReturnValue
        if (xVal > yVal) return grtrThanReturnValue
        return 0
      })
      return {...store, repos: this.reposVal}
      }
    )
  }

  async fetchRepos(): Promise {
    return await fetch("api/github/repos")
      .then((resp) => resp.json())
      .then(async (resp) => {
        const repos = resp.repos
        this.store.update((store: GithubStore) => {
          this.reposVal = this.mapRepos(store.repos)
          return {...store, repos: this.reposVal}
        })

        this.sortReposBy()
        this.updateReposReady(true)
      })
      .catch(() => {
        // defaut to snapshot if any trouble happens along the way
        const parsedFixture: GithubRepos = this.mapRepos(reposFixture)
        this.reposVal = parsedFixture
        this.store.update((store: GithubStore) => ({...store, repos: parsedFixture}))
        this.sortReposBy()
        this.updateReposReady(true)
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
