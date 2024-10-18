package usecase

import (
	"github.com/gabogomes/goexpert-challenge-3-clean-architecture/internal/entity"
)

type GetAllOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetAllOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *GetAllOrdersUseCase {
	return &GetAllOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (u *GetAllOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := u.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var outputDTOs []OrderOutputDTO
	for _, order := range orders {
		outputDTO := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		outputDTOs = append(outputDTOs, outputDTO)
	}

	return outputDTOs, nil
}
