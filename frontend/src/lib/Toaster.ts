  export enum ToastColor {
    green = "green",
    blue = "blue",
    red = "red",
    gray = "gray",
    yellow = "yellow",
    indigo = "indigo",
    purple = "purple",
    orange = "orange",
    none = "none"
  }

  export interface ToasterProps {
    color: ToastColor,
    text: string,
  }
