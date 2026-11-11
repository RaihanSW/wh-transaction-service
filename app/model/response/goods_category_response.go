package response

import "dc-transaction-service/app/model/entity"

type GoodsCategoryResponse struct {
	GeneralResponse
}

type GoodsCategoryDataResponse struct {
	GeneralResponse
	Data entity.GoodsCategory `json:"data,omitempty"`
}

type GoodsCategoryListResponse struct {
	GeneralResponse
	Data []entity.GoodsCategory `json:"data,omitempty"`
}
