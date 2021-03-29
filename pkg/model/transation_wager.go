package model

import (
	"time"
)

type TransactionWager struct {
	Id int64 `json:"id" db:"id"`
	BuyingPrice float32 `json:"buying_price" db:"buying_price"`
	WagerId int64 `json:"wager_id" db:"wager_id"`
	BoughtAt time.Time `json:"bought_at" db:"bought_at"`
}