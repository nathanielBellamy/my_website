import { type Day } from '../Days'
import { type FixedLengthArray } from '../FixedLengthArray'

export interface DatedDay { // TODO:
                            // remove this type
                            // i18n can handle abbrevs
  day: Day;
  date: Date;
}

export type CalendarState = FixedLengthArray<DatedDay, 42>
