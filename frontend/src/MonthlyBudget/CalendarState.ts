import { type Day } from './Days'
import { type FixedLengthArray } from './FixedLengthArray'

export interface DatedDay {
  day: Day;
  date: Date;
}

export type CalendarState = FixedLengthArray<DatedDay, 35>
