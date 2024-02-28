import reposJsonFixture from './repos.json'

export interface GithubRepoLangBreakdown { [key: String]: number }

export interface GithubRepo {
  created_at: Date,
  description: String,
  html_url: String,
  language: String,
  languageBreakdown: GithubRepoLangBreakdown,
  languageData: LanguageData
  name: String,
  pushed_at: Date,
  updated_at: Date,
}

export type GithubRepos = GithubRepo[]

export enum SortOrder {
  ASC = "asc",
  DESC = "desc"
}
export enum SortColumns {
  NAME = "name",
  LANGUAGE = "language",
  DESCRIPTION = "description",
  PUSHED_AT = "pushed_at",
  UPDATED_AT = "updated_at",
  CREATED_AT = "created_at"
}
const LOWERCASE_SORT_COLUMNS = [SortColumns.NAME, SortColumns.DESCRIPTION]

export class GithubIntegration {
  sortOrder: SortOrder = SortOrder.DESC
  sortColumn: SortColumns = SortColumns.PUSHED_AT
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

  languageBreakdownToData(breakdown: GithubRepoLangBreakdown|null): LanguageData|void {
    if (breakdown)
    {
      var data: LanguageData
      data = Object.keys(breakdown).map(k => ({name: k, value: breakdown[k]}) )
      return data
    }
  }

  sortReposBy(col: SortColumns|null = null): void {
    if (col)
    {
      this.sortColumn = col
    }
    const lessThanReturnValue: number = this.sortOrder === SortOrder.ASC ? -1 : 1
    const grtrThanReturnValue: number = -1 * lessThanReturnValue
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

  fetchRepos() {
    const url: String = "https://api.github.com/users/nathanielBellamy/repos"
    return fetch(
      url,
      {
        method: "GET",
        headers: { Authorization: `token ${import.meta.env.VITE_GITHUB_KEY}`}
      }
    )
    .then((res) => res.json())
    .then(async (repos) => {
      const repoLangDict: { [key: String]: GithubRepoLangBreakdown[] } = {}
      await this.fetchRepoLanguageBreakdowns(repos, repoLangDict);
      this.repos.update(() => this.reposVal = repos.map(repo => {
        return {
          created_at: new Date(repo.created_at),
          description: repo.description,
          html_url: repo.html_url,
          language: repo.language,
          languageBreakdown: repoLangDict[repo.name],
          languageData: this.languageBreakdownToData(repoLangDict[repo.name]),
          name: repo.name,
          pushed_at: new Date(repo.pushed_at),
          updated_at: new Date(repo.updated_at)
        }
      }))
    })
    .then(() => this.sortReposBy())
    .then(() => this.updateReposReady(true))
    .catch((e) => console.error(e))
  }

  async fetchRepoLanguageBreakdowns(repos: any[], repoLangDict: any): void {
    if (typeof repos !== "array") repos = reposJsonFixture
    const languagesPromises: Promise[] = repos.map(repo => {
      const repoLanguagesUrl = `https://api.github.com/repos/nathanielBellamy/${repo.name}/languages`
      return fetch(
        repoLanguagesUrl,
        {
          method: "GET",
          headers: { Authorization: `token ${import.meta.env.VITE_GITHUB_KEY}`} }
      )
      .then(res => {
        if (res.status !== 200)
        {
          throw new Error(`Trouble Fetching Language Data for ${repo.name} from Github`)
        }
        return res
      })
      .then(res => res.json())
      .then(res => repoLangDict[repo.name] = res)
      .catch((e) => console.error(e))
    })

    return await Promise.all(languagesPromises)
  }
}
