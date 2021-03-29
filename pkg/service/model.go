package service

type CreateWagerRequest struct {
	TotalWagerValue int `json:"total_wager_value" binding:"required,gt=0"`
	Odds int `json:"odds"  binding:"required,gt=0"`
	SellingPercentage int `json:"selling_percentage"  binding:"required,gte=1,lte=100"`
	SellingPrice float32 `json:"selling_price"  binding:"required"`
}

type BuyWagerRequest struct {
	BuyingPrice float32 `json:"buying_price" binding:"required,gt=0"`
}

type ListWagerRequestQuery struct {
	Page  int `form:"page" binding:"gte=0"`
	Limit int `form:"limit" binding:"gt=0"`
}