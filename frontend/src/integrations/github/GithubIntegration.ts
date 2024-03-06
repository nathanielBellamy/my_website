import reposFixture from './reposFixture.json'
import {
  SortOrder,
  SortColumn,
  type ColorData,
  type GithubRepos,
  DATE_SORT_COLUMNS,
  LOWERCASE_SORT_COLUMNS
} from './GithubTypes'

export default class GithubIntegration {
  sortOrder: SortOrder = SortOrder.DESC
  sortColumn: SortColumn = SortColumn.PUSHED_AT
  repos: Writeable<GithubRepos>
  reposReady: boolean = false
  reposVal: GithubRepos
  updateReposReady: (val: boolean) => void

  constructor(
    repos: Writeable<GithubRepos>,
    updateReposReady: (val: boolean) => void,
    updateRepos: (val: GithubRepos) => void
  ) {
    this.repos = repos
    this.repos.subscribe(val => this.reposVal = val)
    this.updateReposReady = (val: boolean) => updateReposReady(val)
    this.updateRepos = (val: boolean) => updateRepos(val)
  }

  sortReposBy(col: SortColumn|null = null): void {
    if (col)
    {
      this.sortColumn = col
    }
    let lessThanReturnValue: number = this.sortOrder === SortOrder.ASC ? 1 : -1
    if (DATE_SORT_COLUMNS.includes(this.sortColumn)) lessThanReturnValue *= -1

    let grtrThanReturnValue: number = -1 * lessThanReturnValue
    this.repos.update(() => [...this.reposVal.sort((x: any, y: any) => {
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
      })]
    )
  }

  async fetchRepos() {
    return await fetch("api/github/repos")
      .then((resp) => resp.json())
      .then(async (resp) => {
        const repos = resp.repos
        this.repos.update(() => {
          this.reposVal = this.mapRepos(repos)
          return this.reposVal
        })

        this.sortReposBy()
        this.updateReposReady(true)
      })
      .catch(() => {
        // defaut to snapshot if any trouble happens along the way
        const parsedFixture: GithubRepos = this.mapRepos(reposFixture)
        this.reposVal = parsedFixture
        this.repos.update(() => parsedFixture)
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
