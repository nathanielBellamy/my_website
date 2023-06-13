export enum DrawPattern {
    fix1 = "Fix1",
    fix2 = "Fix2",
    fix3 = "Fix3",
    fix4 = "Fix4",
    fix5 = "Fix5",
    fix6 = "Fix6",
    fix7 = "Fix7",
    fix8 = "Fix8",
    out1 = "Out1",
    out2 = "Out2",
    out3 = "Out3",
    out4 = "Out4",
    out5 = "Out5",
    out6 = "Out6",
    out7 = "Out7",
    out8 = "Out8",
    in1 = "In1",
    in2 = "In2",
    in3 = "In3",
    in4 = "In4",
    in5 = "In5",
    in6 = "In6",
    in7 = "In7",
    in8 = "In8",
}

export function intoDrawPattern(s: string): DrawPattern {
  switch (s) {
    case "Fix1":
      return DrawPattern.fix1
    case "Fix2":
      return DrawPattern.fix2
    case "Fix3":
      return DrawPattern.fix3
    case "Fix4":
      return DrawPattern.fix4
    case "Fix5":
      return DrawPattern.fix5
    case "Fix6":
      return DrawPattern.fix6
    case "Fix7":
      return DrawPattern.fix7
    case "Fix8":  
      return DrawPattern.fix8
    case "Out1":
      return DrawPattern.out1
    case "Out2":
      return DrawPattern.out2
    case "Out3":
      return DrawPattern.out3
    case "Out4":
      return DrawPattern.out4
    case "Out5":
      return DrawPattern.out5
    case "Out6":
      return DrawPattern.out6
    case "Out7":
      return DrawPattern.out7
    case "Out8":
      return DrawPattern.out8
    case "In1":
      return DrawPattern.in1
    case "In2":
      return DrawPattern.in2
    case "In3":
      return DrawPattern.in3
    case "In4":
      return DrawPattern.in4
    case "In5":
      return DrawPattern.in5
    case "In6":
      return DrawPattern.in6
    case "In7":
      return DrawPattern.in7
    case "In8":
      return DrawPattern.in8
  }
}
