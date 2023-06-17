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

export function intoColorGradient(cg: any): ColorGradient {
  console.log("hai - ho! let's go!")
  console.log(cg)
  return { color_a: [ 255, 0, 0, 1 ], idx_a: 0, color_b: [ 0, 0, 255, 1 ], idx_b: 9 }
}
