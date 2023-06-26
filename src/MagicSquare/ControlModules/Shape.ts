export enum Shape {
  hexagon = "Hexagon",
  cube = "Cube",
  icosahedron = "Icosahedron",
  square = "Square",
  triangle = "Triangle",
  none = "None",
}

export function intoShape(s: string): Shape {
  switch (s) {
    case "Hexagon":
      return Shape.hexagon
    case "Cube":
      return Shape.cube
    case "Icosahedron":
      return Shape.icosahedron
    case "Square":
      return Shape.square
    case "Triangle":
      return Shape.triangle
    case "None":
      return Shape.none
  }
}
