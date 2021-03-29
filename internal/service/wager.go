package service

import (
	"fmt"
	"github.com/nguyenhoai890/wager_service/internal/repository"
	errService "github.com/nguyenhoai890/wager_service/pkg/error"
	"github.com/nguyenhoai890/wager_service/pkg/model"
	"github.com/nguyenhoai890/wager_service/pkg/service"
)

type Wager struct {
	repo repository.IWager
}

func Init(repo repository.IWager) service.IWager {
	return &Wager{repo: repo}
}

func (service Wager) Create(request service.CreateWagerRequest) (createdModel *model.Wager,err error) {
	wager := ToModel(request)
	if err = wager.IsValid(); err != nil {
		return
	}
	wager.CurrentSellingPrice = wager.SellingPrice
	createdWager, err := service.repo.Insert(wager)
	if err != nil {
		return
	}
	createdModel = &createdWager
	return
}

func (service Wager) Buy(wagerId int64, request service.BuyWagerRequest) (*model.TransactionWager, error) {
	if wagerId == 0 || request.BuyingPrice == 0 {
		return nil, errService.ErrInvalidData
	}
	return service.repo.Buy(wagerId, request.BuyingPrice)
}

func (service Wager) List(page int, limit int) ([]model.Wager, error) {
	fmt.Println(page, limit)
	if limit <= 0 || page < 0 {
		return nil, errService.ErrInvalidData
	}
	return service.repo.List(page, limit)
}