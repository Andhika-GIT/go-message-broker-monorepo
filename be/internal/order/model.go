package order

type OrderImport struct {
	Email       string `json:"email"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}
