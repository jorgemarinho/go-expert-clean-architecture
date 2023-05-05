package usecase

import "github.com/jorgemarinho/go-expert-clean-architecture/internal/entity"

type ListOrdersOrderOutputDTO struct {
	ID         string  `json:"id"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: orderRepository}
}

func (u ListOrdersUseCase) Execute() ([]ListOrdersOrderOutputDTO, error) {
	orders, err := u.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var dtos []ListOrdersOrderOutputDTO
	for _, order := range orders {
		dtos = append(dtos, ListOrdersOrderOutputDTO{
			ID:         order.ID,
			FinalPrice: order.FinalPrice,
		})
	}
	return dtos, nil
}
