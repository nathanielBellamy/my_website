import { localeEn } from '../locales/en'
import { localeEs } from '../locales/es'
import { localeFr } from '../locales/fr'

export enum Lang {
  en = "en",
  es = "es",
  fr = "fr"
}

export class I18n {
  private locales: any = {
    en: localeEn,
    es: localeEs,
    fr: localeFr
  }

  t(path: string, lang: Lang): string {
    const steps: string[] = path.split('/')
    const locales: any = !!this.locales[lang] ? this.locales[lang] : this.locales[Lang.en]
    var result: string | null = null
    steps.forEach((level: string, idx: number) => {
      if (!idx) { // first step
        result = locales[level]
      } else {
        result = (result || {})[level]
      }
    })

    if (typeof result == 'string') {
      return result
    } else {
      return "Translation not found."
    }
  }
}

