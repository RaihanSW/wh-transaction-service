package request

import "time"

type TransactionRequest struct {
	Code        string     `json:"code" validate:"required"`
	DateProcess *time.Time `json:"dateprocess"`
	TtlWeight   int        `json:"ttl_weight"`
	Status      int        `json:"status" validate:"required"`
	EndDate     *time.Time `json:"end_date"`
	TtlVolume   int        `json:"ttl_volume"`
}
