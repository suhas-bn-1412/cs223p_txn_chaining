package main

import (
	"context"

	pb "./se_coordinator/"
)

type SeCoordinatorServerImpl struct {
	pb.UnimplementedSeCoordinatorServer
	DatacenterID string
	Clients      map[string]pb.SeCoordinatorClient
	Coordinator  *SeCoordinator
}

func (s *SeCoordinatorServerImpl) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	err := s.Coordinator.CreateUser(ctx, req.UserId, req.Name, req.AccountBalance)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Status: "User created successfully"}, nil
}

func (s *SeCoordinatorServerImpl) AddSymbol(ctx context.Context, req *pb.AddSymbolRequest) (*pb.AddSymbolResponse, error) {
	err := s.Coordinator.AddSymbol(ctx, req.Symbol, req.Name, req.CurrentPrice)
	if err != nil {
		return nil, err
	}
	return &pb.AddSymbolResponse{Status: "Symbol added successfully"}, nil
}

func (s *SeCoordinatorServerImpl) UpdatePrice(ctx context.Context, req *pb.UpdatePriceRequest) (*pb.UpdatePriceResponse, error) {
	err := s.Coordinator.UpdatePrice(ctx, req.Symbol, req.NewPrice)
	if err != nil {
		return nil, err
	}
	return &pb.UpdatePriceResponse{Status: "Price updated successfully"}, nil
}

func (s *SeCoordinatorServerImpl) BuyLocal(ctx context.Context, req *pb.BuyLocalRequest) (*pb.BuyLocalResponse, error) {
	err := s.Coordinator.BuyLocal(ctx, req.UserId, req.Symbol, int(req.Quantity), req.Price)
	if err != nil {
		return nil, err
	}
	return &pb.BuyLocalResponse{Status: "Buy local successful"}, nil
}

func (s *SeCoordinatorServerImpl) BuyInternational(ctx context.Context, req *pb.BuyInternationalRequest) (*pb.BuyInternationalResponse, error) {
	err := s.Coordinator.BuyInternational(ctx, req.UserId, req.Symbol, int(req.Quantity), req.Price)
	if err != nil {
		return nil, err
	}
	return &pb.BuyInternationalResponse{Status: "Buy international successful"}, nil
}

func (s *SeCoordinatorServerImpl) SellLocal(ctx context.Context, req *pb.SellLocalRequest) (*pb.SellLocalResponse, error) {
	err := s.Coordinator.SellLocal(ctx, req.UserId, req.Symbol, int(req.Quantity), req.Price)
	if err != nil {
		return nil, err
	}
	return &pb.SellLocalResponse{Status: "Sell local successful"}, nil
}

func (s *SeCoordinatorServerImpl) SellInternational(ctx context.Context, req *pb.SellInternationalRequest) (*pb.SellInternationalResponse, error) {
	err := s.Coordinator.SellInternational(ctx, req.UserId, req.Symbol, int(req.Quantity), req.Price)
	if err != nil {
		return nil, err
	}
	return &pb.SellInternationalResponse{Status: "Sell international successful"}, nil
}

func (s *SeCoordinatorServerImpl) CheckBalance(ctx context.Context, req *pb.CheckBalanceRequest) (*pb.CheckBalanceResponse, error) {
	balance, err := s.Coordinator.CheckBalance(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.CheckBalanceResponse{AccountBalance: balance}, nil
}
