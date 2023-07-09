export enum ColorDirection {
  in = "In",
  fix = "Fix",
  out = "Out"
}

export function intoColorDirection(cd: string): ColorDirection {
  switch (cd) {
    case "In":
      return ColorDirection.in
    case "Fix":
      return ColorDirection.fix
    case "Out":
      return ColorDirection.out
  }
}

export interface ColorGradient {
  idx_a: number,
  idx_b: number,
}
