package order

import (
	"net/http"

	"gorm.io/gorm"
)

func ConvertToOrdersResponse(orders []Order) []OrderResponse {
	var ordersResp []OrderResponse

	for _, order := range orders {
		resp := OrderResponse{
			ID:          order.ID,
			Email:       order.User.Email,
			ProductName: order.ProductName,
			Quantity:    order.Quantity,
		}

		ordersResp = append(ordersResp, resp)
	}

	return ordersResp
}

func BindOrderFilterFromRequest(r *http.Request) *OrderFilter {
	return &OrderFilter{
		Email:       r.URL.Query().Get("email"),
		ProductName: r.URL.Query().Get("product_name"),
		Search:      r.URL.Query().Get("search"),
	}
}

func FilterOrderQuery(filter *OrderFilter, query *gorm.DB) *gorm.DB {

	if filter.Search != "" {
		query = query.Joins("LEFT JOIN users ON users.id = orders.user_id").Where(
			"LOWER(email) ILIKE LOWER(?) OR LOWER(product_name) ILIKE LOWER(?)",
			"%"+filter.Search+"%",
			"%"+filter.Search+"%",
		)
	}

	if filter.Email != "" {
		query = query.Where("users.email = ?", filter.Email)
	}
	if filter.ProductName != "" {
		query = query.Where("orders.phone_number = ?", filter.ProductName)
	}

	return query
}
