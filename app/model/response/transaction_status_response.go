package response

import "wh-transaction-service/app/model/entity"

type TransactionStatusResponse struct {
	GeneralResponse
}

type TransactionStatusDataResponse struct {
	GeneralResponse
	Data entity.TransactionStatus `json:"data,omitempty"`
}

type TransactionStatusListResponse struct {
	GeneralResponse
	Data []entity.TransactionStatus `json:"data,omitempty"`
}
