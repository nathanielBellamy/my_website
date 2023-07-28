export enum Module {
  color = 'color',
  drawPattern = 'drawPattern',
  feed = 'feed',
  geometry = 'geometry',
  lfo = 'lfo',
  presets = 'presets',
  rotation = 'rotation',
  translation = 'translation'
}

export function into_module(s: string): Module {
  switch (s) {
    case "color":
      return Module.color
    case "drawPattern":
      return Module.drawPattern
    case "feed":
      return Module.feed
    case "geometry":
      return Module.geometry
    case "lfo":
      return Module.lfo
    case "presets":
      return Module.presets
    case "rotation":
      return Module.rotation
    case "translation":
      return Module.translation
    default:
      return Module.presets
  }
}
