package service

import "github.com/nguyenhoai890/wager_service/pkg/model"

type IWager interface {
	Create(request CreateWagerRequest) (createdModel *model.Wager,err error)
	Buy(wagerId int64, request BuyWagerRequest) (*model.TransactionWager, error)
	List(page int, limit int) ([]model.Wager, error)
}
