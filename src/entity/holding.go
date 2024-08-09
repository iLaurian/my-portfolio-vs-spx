package entity

type Holding struct {
	Symbol      string  `json:"symbol"`
	Shares      float32 `json:"shares"`
	MarketValue float32 `json:"marketvalue"`
	OpenPrice   float32 `json:"openprice"`
	MarketPrice float32 `json:"marketprice"`
	GrossProfit float32 `json:"grossprofit"`
	WinOrLoss   float32 `json:"winorloss"`
}
