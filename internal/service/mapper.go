package service

import (
	"github.com/nguyenhoai890/wager_service/pkg/model"
	"github.com/nguyenhoai890/wager_service/pkg/service"
)

func ToModel(request service.CreateWagerRequest) model.Wager {
	return model.Wager{
		TotalWagerValue: request.TotalWagerValue,
		SellingPrice: request.SellingPrice,
		Odds: request.Odds,
		SellingPercentage: request.SellingPercentage,
	}
}
