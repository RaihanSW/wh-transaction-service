package response

import "dc-transaction-service/app/model/entity"

type GoodsResponse struct {
	GeneralResponse
}

type GoodsDataResponse struct {
	GeneralResponse
	Data entity.Goods `json:"data,omitempty"`
}

type GoodsListResponse struct {
	GeneralResponse
	Data []entity.Goods `json:"data,omitempty"`
}
