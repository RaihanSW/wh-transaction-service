package response

import "dc-transaction-service/app/model/entity"

type GoodsStatusResponse struct {
	GeneralResponse
}

type GoodsStatusDataResponse struct {
	GeneralResponse
	Data entity.GoodsStatus `json:"data,omitempty"`
}

type GoodsStatusListResponse struct {
	GeneralResponse
	Data []entity.GoodsStatus `json:"data,omitempty"`
}
