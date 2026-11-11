package request

type GoodsStatusRequest struct {
	Name string `json:"name" validate:"required"`
}
