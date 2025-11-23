package order

import (
	"context"
	"fmt"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

type OrderUseCase struct {
	Repository *OrderRepository
	rmq        *shared.RabbitMqProducer
}

func NewOrderUseCase(Repository *OrderRepository, rmq *shared.RabbitMqProducer) *OrderUseCase {
	return &OrderUseCase{
		Repository: Repository,
		rmq:        rmq,
	}
}

func (u *OrderUseCase) FindAllOrders(c context.Context, paginationReq *shared.PaginationRequest, filter *OrderFilter) (*shared.Paginated[OrderResponse], error) {
	paginated, err := u.Repository.FindAll(c, paginationReq, filter)

	if err != nil {
		return nil, shared.WriteError(500, fmt.Sprintf("failed to find all users %s", err.Error()))
	}

	formatedOrders := ConvertToOrdersResponse(paginated.Data)

	// return new paginated response with different type (OrderResponse)
	return &shared.Paginated[OrderResponse]{
		Data:       formatedOrders,
		Total:      paginated.Total,
		TotalPages: paginated.TotalPages,
	}, nil

}
