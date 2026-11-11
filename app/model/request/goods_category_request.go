package request

type GoodsCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}
