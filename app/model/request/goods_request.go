package request

import "time"

type GoodsRequest struct {
	DateProcess    *time.Time `json:"dateprocess"`
	TransactionUID string     `json:"transaction_uid" validate:"required"`
	CategoryID     int        `json:"category_id" validate:"required"`
	Status         int        `json:"status" validate:"required"`
	ExpiredDate    *time.Time `json:"expired_date"`
}
