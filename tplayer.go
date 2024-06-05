package main

import (
	"context"
	"errors"
)

type TpLayer struct {
	Users    map[string]*User
	Stocks   map[string]*Stock
	Trades   map[string]*Trade
	Holdings map[string]*Holding
}

func NewTpLayer() *TpLayer {
	return &TpLayer{
		Users:    make(map[string]*User),
		Stocks:   make(map[string]*Stock),
		Trades:   make(map[string]*Trade),
		Holdings: make(map[string]*Holding),
	}
}

func (t *TpLayer) ReadUser(ctx context.Context, userID string) (*User, error) {
	user, exists := t.Users[userID]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (t *TpLayer) WriteUser(ctx context.Context, user *User) error {
	t.Users[user.UserID] = user
	return nil
}

func (t *TpLayer) ReadStock(ctx context.Context, symbol string) (*Stock, error) {
	stock, exists := t.Stocks[symbol]
	if !exists {
		return nil, errors.New("stock not found")
	}
	return stock, nil
}

func (t *TpLayer) WriteStock(ctx context.Context, stock *Stock) error {
	t.Stocks[stock.Symbol] = stock
	return nil
}

func (t *TpLayer) ReadTrade(ctx context.Context, tradeID string) (*Trade, error) {
	trade, exists := t.Trades[tradeID]
	if !exists {
		return nil, errors.New("trade not found")
	}
	return trade, nil
}

func (t *TpLayer) WriteTrade(ctx context.Context, trade *Trade) error {
	t.Trades[trade.TradeID] = trade
	return nil
}

func (t *TpLayer) ReadHolding(ctx context.Context, userID string) (*Holding, error) {
	holding, exists := t.Holdings[userID]
	if !exists {
		return nil, errors.New("holding not found")
	}
	return holding, nil
}

func (t *TpLayer) WriteHolding(ctx context.Context, holding *Holding) error {
	t.Holdings[holding.UserID] = holding
	return nil
}
