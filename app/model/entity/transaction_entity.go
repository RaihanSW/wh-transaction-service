package entity

import "time"

type Transaction struct {
	UID         string     `json:"uid"`
	Code        string     `json:"code"`
	DateProcess *time.Time `json:"dateprocess,omitempty"`
	TtlWeight   int        `json:"ttl_weight"`
	Status      int        `json:"status"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	TtlVolume   int        `json:"ttl_volume"`
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	UpdatedAt   time.Time  `json:"updated_at"`
	UpdatedBy   string     `json:"updated_by"`
}
