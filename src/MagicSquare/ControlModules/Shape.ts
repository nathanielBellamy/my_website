export enum Shape {
  triangle = "Triangle",
  hexagon = "Hexagon",
  cube = "Cube",
  icosahedron = "Icosahedron",
  none = "None",
}

export function intoShape(s: string): Shape {
  switch (s) {
    case "Triangle":
      return Shape.triangle
    case "Hexagon":
      return Shape.hexagon
    case "Cube":
      return Shape.cube
    case "Icosahedron":
      return Shape.icosahedron
    case "None":
      return Shape.none
  }
}
