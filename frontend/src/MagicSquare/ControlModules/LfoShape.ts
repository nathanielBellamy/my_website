export enum LfoShape {
  linear = "Linear",
  sine = "Sine"
}

export function intoLfoShape(s: string): LfoShape {
  switch (s) {
    case "Linear":
      return LfoShape.linear
    default:
      return LfoShape.sine
  }
}
