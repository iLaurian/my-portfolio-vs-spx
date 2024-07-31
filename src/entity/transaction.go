package entity

type Transaction struct {
	Ticker string  `json:"ticker"`
	Volume float32 `json:"volume"`
	Price  float32 `json:"price"`
	Date   string  `json:"date"`
}
