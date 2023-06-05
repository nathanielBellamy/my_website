import { localeEn } from '../locales/en'
import { localeEs } from '../locales/es'
import { localeFr } from '../locales/fr'

export enum Lang {
  en = "en",
  es = "es",
  fr = "fr"
}

export class I18n {
  
  public prefix: string[]

  private locales: any = {
    en: localeEn,
    es: localeEs,
    fr: localeFr
  }

  constructor(prefix:string ) {
    this.prefix = prefix.split('/')
  }

  t(path: string, lang: Lang): string {
    const steps: string[] = this.prefix.concat(path.split('/'))
    const locales: any = !!this.locales[lang] ? this.locales[lang] : this.locales[Lang.en]
    var result: string = locales[steps.shift()]
    if (steps.length) {
      steps.forEach((level: string) => {
        result = (result || {})[level]
      })
    }

    if (typeof result == 'string') {
      return result
    } else {
      switch (lang) {
        case Lang.en:
          return "Translation not found"
        case Lang.es:
          return "Traducción no encontrada"
        case Lang.fr:
          return "Traduction non trouvée"
      }
    }
  }
}


