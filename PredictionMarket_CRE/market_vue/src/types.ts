export interface Market {
  id: number
  description: string
  state: number // 0: Open, 1: Closed, 2: Resolved
  createTimeStamp: number
  closeTime: number
  resolveWindow: number
  resolveTimeStamp: number
  outcome: number
  resolved: boolean
  totalPool: string
  betsCount?: number // 下注数量（可选）
}

export enum MarketStatus {
  Open = 0,
  Closed = 1,
  Resolved = 2,
}
