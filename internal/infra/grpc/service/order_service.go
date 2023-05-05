package service

import (
	"context"

	"github.com/jorgemarinho/go-expert-clean-architecture/internal/infra/grpc/pb"
	"github.com/jorgemarinho/go-expert-clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(
	createOrderUsease usecase.CreateOrderUseCase,
	listOrdersUseCase usecase.ListOrdersUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUsease,
		ListOrdersUseCase:  listOrdersUseCase,
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
		Order: &pb.Order{
			Id:         output.ID,
			Price:      output.Price,
			Tax:        output.Tax,
			FinalPrice: output.FinalPrice,
		},
	}, nil
}

func (s *OrderService) ListOrders(c context.Context, b *pb.Blank) (*pb.OrderList, error) {
	output, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var orders []*pb.Order
	for _, orderOutput := range output {
		orders = append(orders, &pb.Order{
			Id:         orderOutput.ID,
			FinalPrice: orderOutput.FinalPrice,
		})
	}
	return &pb.OrderList{Orders: orders}, nil
}
