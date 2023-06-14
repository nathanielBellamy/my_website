export enum TransformOrder {
  rotateThenTranslate = "RotateThenTranslate",
  translateThenRotate = "TranslateThenRotate"
}

export function intoTransformOrder(order: string): TransformOrder {
  switch (order) {
    case "RotateThenTranslate":
      return TransformOrder.rotateThenTranslate
    case "TranslateThenRotate":
      return TransformOrder.translateThenRotate
  }
}
