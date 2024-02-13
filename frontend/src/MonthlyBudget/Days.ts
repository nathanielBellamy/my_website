export interface Day {
  id: number;
  name: string;
  abbreviation_1: string;
  abbreviation_3: string;
}

export const Days: { [key: number]: Day; } = {
  0: {
    id: 0,
    name: 'Sunday',
    abbreviation_1: 'Su',
    abbreviation_3: 'Sun'
  },
  1: {
    id: 1,
    name: 'Monday',
    abbreviation_1: 'M',
    abbreviation_3: 'Mon'
  },
  2: {
    id: 2,
    name: 'Tuesday',
    abbreviation_1: 'Tu',
    abbreviation_3: 'Tue'
  },
  3: {
    id: 3,
    name: 'Wednesday',
    abbreviation_1: 'W',
    abbreviation_3: 'Wed'
  },
  4: {
    id: 4,
    name: 'Thursday',
    abbreviation_1: 'Th',
    abbreviation_3: 'Thu'
  },
  5: {
    id: 4,
    name: 'Friday',
    abbreviation_1: 'F',
    abbreviation_3: 'Fri'
  },
  6: {
    id: 4,
    name: 'Saturday',
    abbreviation_1: 'Sa',
    abbreviation_3: 'Sat'
  },
}
