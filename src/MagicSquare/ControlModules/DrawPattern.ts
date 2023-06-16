export enum DrawPatternType {
    fix = "Fix",
    out = "Out",
    in = "In",
}

export function intoDrawPatternType(s: string): DrawPatternType {
  switch (s) {
    case "Fix":
      return DrawPatternType.fix
    case "Out":
      return DrawPatternType.out
    case "In":
      return DrawPatternType.in
    default:
      DrawPatternType.fix
  }
}
