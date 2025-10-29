package order

type OrderImport struct {
	Email       string `json:"email"`
	ProductName string `json:"product_name"`
	Quantity    int64  `json:"quantity"`
}

type OrderResponse struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	ProductName string `json:"product_name"`
	Quantity    int64  `json:"quantity"`
}

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
