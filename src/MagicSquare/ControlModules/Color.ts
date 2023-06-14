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

export enum ColorMode {
  eight = "Eight",
  gradient = "Gradient"
}

export function intoColorMode(cm: string): ColorMode {
  switch (cm) {
    case "Eight":
      return ColorMode.eight
    case "Gradient":
      return ColorMode.gradient
  }
}
