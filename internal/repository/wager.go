package repository

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	serviceErr "github.com/nguyenhoai890/wager_service/pkg/error"
	"github.com/nguyenhoai890/wager_service/pkg/model"
	mysqlUtils "github.com/nguyenhoai890/wager_service/pkg/mysql"
	"math"
	"time"
)

type Wager struct {
	db *sqlx.DB
}

func Init(uri mysqlUtils.Uri) (IWager, error) {
	driverName := "mysql"
	db, err := sqlx.Open(driverName, uri.GetFullUri())
	if err != nil {
		return nil, err
	}
	return &Wager{db: db}, nil
}

func (repo Wager) Insert(wager model.Wager) (model.Wager, error) {
	if repo.db == nil {
		return wager, errors.New("database was not init")
	}
	wager.PlaceAt = time.Now()
	insertResult, err := repo.db.NamedExec(`INSERT INTO wagers (
						total_wager_value,
						odds,
						selling_percentage,
						selling_price,
						current_selling_price,
						percentage_sold,
						amount_sold,
						placed_at) VALUES (
							:total_wager_value,
							:odds,
							:selling_percentage,
							:selling_price,
							:current_selling_price,
							:percentage_sold,
							:amount_sold,
							:placed_at
						);`, &wager)
	if err != nil {
		return wager, err
	}
	wager.Id, err = insertResult.LastInsertId()
	return wager, err
}

func (repo Wager) Buy(wagerId int64, buyingPrice float32) (transaction *model.TransactionWager,err error) {
	if repo.db == nil {
		err = errors.New("database was not init")
		return
	}
	if buyingPrice <= 0 || wagerId <= 0 {
		err = serviceErr.ErrInvalidData
		return
	}

	buyingPrice = float32(math.Round(float64(buyingPrice)*100)/100)

	tx, err := repo.db.Beginx()
	if err != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		}
	}()

	var wager model.Wager
	err = tx.Get(&wager, "SELECT * FROM wagers WHERE id=? AND (current_selling_price - ?) >= 0 FOR UPDATE", wagerId, buyingPrice)
	if err != nil  {
		if errors.Is(err, sql.ErrNoRows) {
			err =  serviceErr.ErrWagerNotFoundOrCantNotBuy
		}
	}

	if wager.CurrentSellingPrice = wager.CurrentSellingPrice - buyingPrice; wager.CurrentSellingPrice < 0 {
		err = serviceErr.ErrBuyingGreaterThanCurrentSell
		return
	}

	if wager.AmountSold == nil {
		sold := float32(0)
		wager.AmountSold = &sold
	}
	*wager.AmountSold = *wager.AmountSold + buyingPrice

	p :=  (*wager.AmountSold / wager.SellingPrice) * 100
	pRoundUp := float32(math.Round(float64(p)*100)/100)
	wager.PercentageSold = &pRoundUp

	_, err = tx.NamedExec(`Update wagers SET 
		current_selling_price = :current_selling_price,
		amount_sold = :amount_sold,
		percentage_sold = :percentage_sold
		WHERE id = :id
	`, &wager)

	if err != nil {
		return
	}

	tResult := model.TransactionWager{
		BuyingPrice: buyingPrice,
		WagerId:     wagerId,
		BoughtAt:    time.Now(),
	}
	result, err := tx.NamedExec(`INSERT INTO transaction_wagers (
		buying_price, wager_id, bought_at
	) VALUES(:buying_price, :wager_id, :bought_at)`, &tResult)
	if err != nil {
		return
	}
	tResult.Id, err = result.LastInsertId()
	if err != nil {
		return
	}
	err = tx.Commit()
	transaction = &tResult
	return
}

func (repo Wager) List(offset int, limit int) (wagers []model.Wager, err error) {
	if repo.db == nil {
		err = errors.New("database was not init")
		return
	}
	err = repo.db.Select(&wagers, `Select * from wagers LIMIT ? OFFSET ?`, limit, offset)
	return
}

func (repo Wager) Get(wagerId int64) (*model.Wager, error) {
	if repo.db == nil {
		return nil, errors.New("database was not init")
	}
	var wager model.Wager
	err := repo.db.Get(&wager, `Select * from wagers Where id = ?`, wagerId)
	if err != nil && errors.Is(err, sql.ErrNoRows){
		return nil, serviceErr.ErrWagerNotFound
	}
	return &wager, nil
}

func (repo Wager) Ping() error {
	return repo.db.Ping()
}

func (repo Wager) Close() error {
	return repo.db.Close()
}