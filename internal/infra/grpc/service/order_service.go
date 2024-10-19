package service

import (
	"context"

	"github.com/gabogomes/goexpert-challenge-3-clean-architecture/internal/infra/grpc/pb"
	"github.com/gabogomes/goexpert-challenge-3-clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase  usecase.CreateOrderUseCase
	GetAllOrdersUseCase usecase.GetAllOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, getAllOrdersUseCase usecase.GetAllOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase:  createOrderUseCase,
		GetAllOrdersUseCase: getAllOrdersUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) GetAllOrders(ctx context.Context, _ *pb.EmptyRequest) (*pb.GetAllOrdersResponse, error) {
	orders, err := s.GetAllOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrder := &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		}
		pbOrders = append(pbOrders, pbOrder)
	}

	return &pb.GetAllOrdersResponse{
		Orders: pbOrders,
	}, nil
}
