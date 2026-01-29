
export interface FeedMessage {
  clientId: number,
  body: string,
  system: boolean
}

export enum SystemMessage {
  init = "__init__connected__",
  sqConnected = "__sq__connected__",
  sqDisconnected = "__sq__disconnected__",
  none = "none"
}
