package main

type User struct {
	UserID         string  `json:"user_id"`
	Name           string  `json:"name"`
	AccountBalance float32 `json:"account_balance"`
}

type Stock struct {
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	CurrentPrice float32 `json:"current_price"`
}

type Trade struct {
	TradeID  string  `json:"trade_id"`
	UserID   string  `json:"user_id"`
	Symbol   string  `json:"symbol"`
	Quantity int     `json:"quantity"`
	Type     string  `json:"type"` // buy/sell
	Price    float32 `json:"price"`
}

type Holding struct {
	UserID   string `json:"user_id"`
	Symbol   string `json:"symbol"`
	Quantity int    `json:"quantity"`
}
