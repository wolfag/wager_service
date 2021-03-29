package repository

import "github.com/nguyenhoai890/wager_service/pkg/model"

type IWager interface {
	Insert(wager model.Wager) (model.Wager, error)
	Buy(wagerId int64, buyingPrice float32) (transaction *model.TransactionWager,err error)
	List(offset int, limit int) (wagers []model.Wager, err error)
	Get(wagerId int64) (*model.Wager, error)
	Ping() error
	Close() error
}
