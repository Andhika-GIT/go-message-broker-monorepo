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

type OrderFilter struct {
	Email       string
	ProductName string
	Search      string
}
