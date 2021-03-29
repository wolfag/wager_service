package model

import (
	errService "github.com/nguyenhoai890/wager_service/pkg/error"
	"time"
)

type Wager struct {
	Id int64 `json:"id"`
	TotalWagerValue int `json:"total_wager_value" db:"total_wager_value"`
	Odds int `json:"odds" db:"odds"`
	SellingPercentage int `json:"selling_percentage" db:"selling_percentage"`
	SellingPrice float32 `json:"selling_price" db:"selling_price"`
	CurrentSellingPrice float32 `json:"current_selling_price" db:"current_selling_price"`
	PercentageSold *float32 `json:"percentage_sold" db:"percentage_sold"`
	AmountSold *float32 `json:"amount_sold" db:"amount_sold"`
	PlaceAt time.Time `json:"placed_at" db:"placed_at"`
}

func (wager *Wager) IsValid() error {
	if wager.TotalWagerValue <= 0 {
		return  errService.ErrTotalWagerValue
	}
	if wager.Odds <= 0 {
		return errService.ErrOddsValue
	}
	if wager.SellingPercentage < 1 || wager.SellingPercentage > 100 {
		return errService.ErrSellingPercentage
	}
	if wager.SellingPrice <= float32(wager.TotalWagerValue) * (float32(wager.SellingPercentage) / 100) {
		return errService.ErrSellingPrice
	}
	return nil
}