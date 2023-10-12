package usecase

import "github.com/rodrigoaustincascao/fullcycle-golang_intensivo/internal/order/entity"

type OrderInputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPriceUseCase(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{OrderRepository: orderRepository}
}

func (c *CalculateFinalPriceUseCase) Execute(input OrderInputDTO) (*OrderInputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderInputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
