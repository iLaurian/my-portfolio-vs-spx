package entity

type Transaction struct {
	ID     int     `json:"ID"`
	Type   string  `json:"type" binding:"required,gte=2,lte=5"`
	Ticker string  `json:"ticker" binding:"required,gte=1,lte=30"`
	Volume float32 `json:"volume" binding:"required"`
	Price  float32 `json:"price" binding:"required"`
	Date   string  `json:"date" binding:"required"`
}
