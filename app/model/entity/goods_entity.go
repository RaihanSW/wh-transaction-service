package entity

import "time"

type Goods struct {
	UID            string     `json:"uid"`
	DateProcess    *time.Time `json:"dateprocess,omitempty"`
	TransactionUID string     `json:"transaction_uid"`
	CategoryID     int        `json:"category_id"`
	Status         int        `json:"status"`
	ExpiredDate    *time.Time `json:"expired_date,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	CreatedBy      string     `json:"created_by"`
	UpdatedAt      time.Time  `json:"updated_at"`
	UpdatedBy      string     `json:"updated_by"`
}
