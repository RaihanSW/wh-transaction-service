package response

import "dc-transaction-service/app/model/entity"

type TransactionResponse struct {
	GeneralResponse
}

type TransactionDataResponse struct {
	GeneralResponse
	Data entity.Transaction `json:"data,omitempty"`
}

type TransactionListResponse struct {
	GeneralResponse
	Data []entity.Transaction `json:"data,omitempty"`
}
