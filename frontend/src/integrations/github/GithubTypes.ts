export type ColorData = String[];

export interface GithubRepoLangBreakdown { [key: String]: number }

export interface GithubRepo {
  colorData: ColorData,
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

export const LOWERCASE_SORT_COLUMNS = [SortColumns.NAME, SortColumns.DESCRIPTION]
