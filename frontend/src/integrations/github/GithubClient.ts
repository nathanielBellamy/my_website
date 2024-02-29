import reposJsonFixture from './repos.json'

export default class GithubClient {
  #baseUrl: String = "https://api.github.com"
  #userName: String = "nathanielBellamy"

  getUserName(): String {
    return this.#userName
  }

  setUserName(userName: string): void {
    this.#userName = userName
  }

  fetchRepos(): Promise {
    const repoEndpoint = [this.#baseUrl, "users", this.#userName, "repos"].join('/')
    const resolve = (res: any): Promise => res.json()
    const reject = (res: any) => {
        // if live pull fails, default to snapshot
      console.error("Trouble fetching repos.")
      return Promise.resolve(reposJsonFixture)
    }
    return this.handleFetch(repoEndpoint)
      .then((res) => {
        if (res.status !== 200)
        {
          return Promise.resolve(reposJsonFixture)
        }
        else
        {
          return res.json()
        }
      })
  }

  fetchRepoData(repoName: String, endpoint: String): Promise {
    const repoEndpointUrl = [this.#baseUrl, "repos", this.#userName , repoName, endpoint].join('/')
    return this.handleFetch(repoEndpointUrl)
  }

  handleFetch(
    url: String,
    method: String = "GET",
  ): Promise {
    return fetch(
      url,
      {
        method,
        headers: { Authorization: `token ${import.meta.env.VITE_GITHUB_KEY}`}
      }
    )
  }
}
