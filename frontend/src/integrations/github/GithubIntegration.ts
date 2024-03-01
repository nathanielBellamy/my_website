import colors from './colors.json'
import {
  SortOrder,
  SortColumns,
  type ColorData,
  type GithubRepos,
  type GithubLangBreakdown,
  type LanguageData,
  LOWERCASE_SORT_COLUMNS
} from './GithubTypes'
import GithubClient from './GithubClient'

export default class GithubIntegration {
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

  languageBreakdownToColorData(breakdown: GithubRepoLangBreakdown): ColorData {
    var data: ColorData
    data = Object.keys(breakdown).sort((langX, langY) => {
      if (breakdown[langX] < breakdown[langY]) return 1
      if (breakdown[langX] > breakdown[langY]) return -1
      return 0
    }).map(lang => colors[lang].color)
    return data
  }

  languageBreakdownToData(breakdown: GithubRepoLangBreakdown): LanguageData {
    var data: LanguageData
    data = Object.keys(breakdown).map(k => ({name: k, value: breakdown[k]}) )
    return data
  }

  sortReposBy(col: SortColumns|null = null): void {
    if (col)
    {
      this.sortColumn = col
    }
    let lessThanReturnValue: number = this.sortOrder === SortOrder.ASC ? 1 : -1
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

  fetchRepos() {
    const client: GithubClient = new GithubClient()
    client.fetchRepos()
      .then(async (repos) => {
        const repoLangDict: { [key: String]: GithubRepoLangBreakdown[] } = {}
        const repoCommitDict: { [key: String]: any[] } = {}
        await Promise.all([
          this.fetchRepoLanguageBreakdowns(repos, repoLangDict),
          this.fetchRepoCommitHistory(repos, repoCommitDict)
        ])
        this.repos.update(() => {
          this.reposVal = repos.map(repo => {
            return {
              colorData: this.languageBreakdownToColorData(repoLangDict[repo.name]),
              commitData: repoCommitDict[repo.name],
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
          })
          return this.reposVal
        })
      })
      .then(() => this.sortReposBy())
      .then(() => this.updateReposReady(true))
      .catch((e) => console.error(e))
  }

  async fetchRepoCommitHistory(repos: any[], repoCommitDict: any): void {
    const client: GithubClient = new GithubClient();
    const commitPromises: Promise[] = repos.map(repo => {
      const repoCommitsUrl = `https://api.github.com/repos/nathanielBellamy/${repo.name}/commits`
      return client
         .fetchRepoData(repo.name, 'commits')
         .then(res => res.json())
         .then(commitsData => repoCommitDict[repo.name] = commitsData)
    })
    return await Promise.all(commitPromises)
  }

  async fetchRepoLanguageBreakdowns(repos: any[], repoLangDict: any): void {
    const client: GithubClient = new GithubClient();
    const languagesPromises: Promise[] = repos.map(repo => {
      const repoLanguagesUrl = `https://api.github.com/repos/nathanielBellamy/${repo.name}/languages`
      return client
        .fetchRepoData(repo.name, "languages")
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
