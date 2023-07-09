export interface Shape {
  t: ShapeTag,
  c: number
}

export enum ShapeTag {
  misc = "Misc",
  ngon = "Ngon",
  platoThree = "PlatoThree"
}

export function intoShape(from_serde: any): Shape {
  return { t: from_serde.t, c: from_serde.c }
}
