export enum LfoDestination {
  // rotation
  pitchBase = 'PitchBase',
  pitchSpread = 'PitchSpread',
  pitchX = 'PitchX',
  pitchY = 'PitchY',
  rollBase = 'RollBase',
  rollSpread = 'RollSpread',
  rollX = 'RollX',
  rollY = 'RollY',
  yawBase = 'YawBase',
  yawSpread = 'YawSpread',
  yawX = 'YawX',
  yawY = 'YawY',

  // radius
  radiusBase = 'RadiusBase',
  radiusStep = 'RadiusStep',

  // translation
  translationXBase = 'TranslationXBase',
  translationXSpread = 'TranslationXSpread',
  translationYBase = 'TranslationYBase',
  translationYSpread = 'TranslationYSpread',
  none = 'None'
}

export function intoLfoDestination(s: string): LfoDestination {
  switch (s) {
    // rotation
    case "PitchBase":
      return LfoDestination.pitchBase
    case "PitchSpread":
      return LfoDestination.pitchSpread
    case "PitchX":
      return LfoDestination.pitchX
    case "PitchY":
      return LfoDestination.pitchY
    case "RollBase":
      return LfoDestination.rollBase
    case "RollSpread":
      return LfoDestination.rollSpread
    case "RollX":
      return LfoDestination.rollX
    case "RollY":
      return LfoDestination.rollY
    case "YawBase":
      return LfoDestination.yawBase
    case "YawSpread":
      return LfoDestination.yawSpread
    case "YawX":
      return LfoDestination.yawX
    case "YawY":
      return LfoDestination.yawY

    // geometry
    case "RadiusBase":
      return LfoDestination.radiusBase
    case "RadiusStep":
      return LfoDestination.radiusStep

    // translation
    case "TranslationXBase":  
      return LfoDestination.translationXBase
    case "TranslationXSpread":
      return LfoDestination.translationXSpread
    case "TranslationYBase": 
      return LfoDestination.translationYBase
    case "translationYSpread":
      return LfoDestination.translationYSpread
    default:
      return LfoDestination.none
  }
}
