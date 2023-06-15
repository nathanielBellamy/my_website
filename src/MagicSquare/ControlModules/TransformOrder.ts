export enum TransformOrder {
  rotateThenTranslate = "RotateThenTranslate",
  translateThenRotate = "TranslateThenRotate"
}

export function intoTransformOrder(order: string): TransformOrder {
  console.log(order)
  switch (order) {
    case "RotateThenTranslate":
      return TransformOrder.rotateThenTranslate
    case "TranslateThenRotate":
      return TransformOrder.translateThenRotate
  }
}
