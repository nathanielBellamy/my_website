import { type Day } from '../Days'
import { type FixedLengthArray } from '../FixedLengthArray'

export const CALENDAR_STATE_LENGTH = 42
export type CalendarState = FixedLengthArray<Date, CALENDAR_STATE_LENGTH>

const dummyDate = new Date()
export const initialCalendarState: CalendarState = [
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,

  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,

  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,

  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,
  dummyDate,

  dummyDate,
  dummyDate,
]
