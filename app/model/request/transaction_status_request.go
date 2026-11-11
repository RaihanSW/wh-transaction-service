package request

type TransactionStatusRequest struct {
	Name string `json:"name" validate:"required"`
}
