export interface Month {
  id: number;
  name: string;
  abbreviation_1: string;
  abbreviation_3: string;
}

export const Months: { [key: number]: Month; } = {
  0: {
    id: 0,
    name: 'January',
    abbreviation_1: 'Ja',
    abbreviation_3: 'Jan'
  },
  1: {
    id: 1,
    name: 'February',
    abbreviation_1: 'F',
    abbreviation_3: 'Feb'
  },
  2: {
    id: 2,
    name: 'March',
    abbreviation_1: 'Mr',
    abbreviation_3: 'Mar'
  },
  3: {
    id: 3,
    name: 'April',
    abbreviation_1: 'Ap',
    abbreviation_3: 'Apr'
  },
  4: {
    id: 4,
    name: 'May',
    abbreviation_1: 'My',
    abbreviation_3: 'May'
  },
  5: {
    id: 5,
    name: 'June',
    abbreviation_1: 'Jn',
    abbreviation_3: 'Jun'
  },
  6: {
    id: 6,
    name: 'July',
    abbreviation_1: 'Jl',
    abbreviation_3: 'Jul'
  },
  7: {
    id: 7,
    name: 'August',
    abbreviation_1: 'Au',
    abbreviation_3: 'Aug'
  },
  8: {
    id: 8,
    name: 'September',
    abbreviation_1: 'S',
    abbreviation_3: 'Sep'
  },
  9: {
    id: 9,
    name: 'October',
    abbreviation_1: 'O',
    abbreviation_3: 'Oct'
  },
  10: {
    id: 10,
    name: 'November',
    abbreviation_1: 'N',
    abbreviation_3: 'Nov'
  },
  11: {
    id: 11,
    name: 'December',
    abbreviation_1: 'D',
    abbreviation_3: 'Dec'
  },
}
