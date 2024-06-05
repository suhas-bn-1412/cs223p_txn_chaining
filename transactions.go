package main

import (
	"context"
)

type SeCoordinator struct {
	tpLayer *TpLayer
}

func NewSeCoordinator(tpLayer *TpLayer) *SeCoordinator {
	return &SeCoordinator{
		tpLayer: tpLayer,
	}
}

func (s *SeCoordinator) CreateUser(ctx context.Context, userID, name string, accountBalance float32) error {
	user := &User{UserID: userID, Name: name, AccountBalance: accountBalance}
	return s.tpLayer.WriteUser(ctx, user)
}

func (s *SeCoordinator) AddSymbol(ctx context.Context, symbol, name string, currentPrice float32) error {
	stock := &Stock{Symbol: symbol, Name: name, CurrentPrice: currentPrice}
	return s.tpLayer.WriteStock(ctx, stock)
}

func (s *SeCoordinator) UpdatePrice(ctx context.Context, symbol string, newPrice float32) error {
	stock, err := s.tpLayer.ReadStock(ctx, symbol)
	if err != nil {
		return err
	}
	stock.CurrentPrice = newPrice
	return s.tpLayer.WriteStock(ctx, stock)
}

func (s *SeCoordinator) BuyLocal(ctx context.Context, userID, symbol string, quantity int, price float32) error {
	user, err := s.tpLayer.ReadUser(ctx, userID)
	if err != nil {
		return err
	}

	holding, err := s.tpLayer.ReadHolding(ctx, userID)
	if err != nil {
		return err
	}

	trade := &Trade{TradeID: "new_trade_id", UserID: userID, Symbol: symbol, Quantity: quantity, Type: "buy", Price: price}
	if err := s.tpLayer.WriteTrade(ctx, trade); err != nil {
		return err
	}

	user.AccountBalance -= float32(quantity) * price
	if err := s.tpLayer.WriteUser(ctx, user); err != nil {
		return err
	}

	holding.Quantity += quantity
	return s.tpLayer.WriteHolding(ctx, holding)
}

func (s *SeCoordinator) BuyInternational(ctx context.Context, userID, symbol string, quantity int, price float32) error {
	// Implement buy international logic using tpLayer
	return nil
}

func (s *SeCoordinator) SellLocal(ctx context.Context, userID, symbol string, quantity int, price float32) error {
	user, err := s.tpLayer.ReadUser(ctx, userID)
	if err != nil {
		return err
	}

	holding, err := s.tpLayer.ReadHolding(ctx, userID)
	if err != nil {
		return err
	}

	trade := &Trade{TradeID: "new_trade_id", UserID: userID, Symbol: symbol, Quantity: quantity, Type: "sell", Price: price}
	if err := s.tpLayer.WriteTrade(ctx, trade); err != nil {
		return err
	}

	user.AccountBalance += float32(quantity) * price
	if err := s.tpLayer.WriteUser(ctx, user); err != nil {
		return err
	}

	holding.Quantity -= quantity
	return s.tpLayer.WriteHolding(ctx, holding)
}

func (s *SeCoordinator) SellInternational(ctx context.Context, userID, symbol string, quantity int, price float32) error {
	// Implement sell international logic using tpLayer
	return nil
}

func (s *SeCoordinator) CheckBalance(ctx context.Context, userID string) (float32, error) {
	user, err := s.tpLayer.ReadUser(ctx, userID)
	if err != nil {
		return 0, err
	}
	return user.AccountBalance, nil
}
