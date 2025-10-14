package order

type OrderStatus string

const (
	StatusPending    OrderStatus = "pending"
	StatusProcessing OrderStatus = "processing"
	StatusCompleted  OrderStatus = "completed"
	StatusCancelled  OrderStatus = "cancelled"
)

type OrderImport struct {
	UserEmail   string      `json:"email"`
	ProductName string      `json:"product_name"`
	Quantity    int         `json:"quantity"`
	Status      OrderStatus `json:"status"`
}
