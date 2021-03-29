package repository

import (
	"fmt"
	"github.com/nguyenhoai890/wager_service/pkg/configuration"
	"github.com/nguyenhoai890/wager_service/pkg/model"
	mysqlUtils "github.com/nguyenhoai890/wager_service/pkg/mysql"
	migrationMysql "github.com/nguyenhoai890/wager_service/tools/migration/mysql"
	"github.com/stretchr/testify/suite"
	"math"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var uri = mysqlUtils.Uri{
	ServerUri: configuration.MysqlServerUri,
	Params: configuration.MysqlServerUriParams,
	DatabaseName: configuration.MysqlTestDB,
}

type WagerTestSuite struct {
	wagerRepo IWager
	suite.Suite
}

func TestWagerTestSuite(t *testing.T) {
	suite.Run(t, new(WagerTestSuite))
}

func (suite *WagerTestSuite) SetupSuite() {
	var err error
	err = migrationMysql.Up(uri, configuration.MigratePath)
	if err != nil {
		panic(err)
	}

	suite.wagerRepo, err = Init(uri)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Testing with DB Name: %s\n", uri.DatabaseName)
}

func (suite *WagerTestSuite) TearDownSuite() {
	defer func() {
		suite.wagerRepo.Close()
	}()
	if err := migrationMysql.Drop(uri, configuration.MigratePath); err != nil {
		suite.Fail(err.Error())
	}
}

func (suite *WagerTestSuite) TestWager_InsertWager_InsertSuccessfully() {
	wager := initValue()
	result, err := suite.wagerRepo.Insert(wager)
	suite.Nil(err, "insert wager was fail")
	suite.True(result.Id > 0, "returned wager id should be greater than 0")
}

func (suite *WagerTestSuite) TestWager_InsertWager_FailWithWrongConstraints() {
	var err error
	var wager model.Wager
	wager = initValue()
	wager.TotalWagerValue = 0


	_, err = suite.wagerRepo.Insert(wager)

	suite.EqualError(err,
		"Error 3819: Check constraint 'total_wager_value_greater_0' is violated.",
		"total wager value should be greater than 0",
	)

	wager = initValue()
	wager.Odds = 0
	_, err = suite.wagerRepo.Insert(wager)
	suite.EqualError(err,
		"Error 3819: Check constraint 'odds_value_greater_0' is violated.",
		"odds should be greater than 0")

	wager = initValue()
	wager.SellingPercentage = 0
	_, err = suite.wagerRepo.Insert(wager)
	suite.EqualError(err,
		"Error 3819: Check constraint 'selling_percentage_between_1_100' is violated.",
		"selling percentage should be between 1 and 100",
	)

	wager = initValue()
	wager.SellingPercentage = 101
	_, err = suite.wagerRepo.Insert(wager)
	suite.EqualError(err, "Error 3819: Check constraint 'selling_percentage_between_1_100' is violated.")

	wager = initValue()
	wager.SellingPercentage = 101
	_, err = suite.wagerRepo.Insert(wager)
	suite.EqualError(err, "Error 3819: Check constraint 'selling_percentage_between_1_100' is violated.")

	wager = initValue()
	wager.SellingPrice = float32(wager.TotalWagerValue) * float32(wager.SellingPercentage / 100)
	_, err = suite.wagerRepo.Insert(wager)
	suite.EqualError(err, "Error 3819: Check constraint 'selling_price_value' is violated.")
}

func (suite *WagerTestSuite) TestWager_BuyWagerSuccessfully(){
	wager := initValue()
	result, err := suite.wagerRepo.Insert(wager)
	if err != nil {
		suite.Fail(err.Error())
	}
	transaction, err := suite.wagerRepo.Buy(result.Id, 3)
	suite.Nil(err, "insert wager was fail")
	suite.True(transaction.Id > 0, "returned transaction id should be greater than 0")
}

func (suite *WagerTestSuite) TestWager_BuyWagerWithMultiRequests() {
	v := initValue()
	wager, err := suite.wagerRepo.Insert(v)
	if err != nil {
		suite.Fail(err.Error())
	}
	maxTotalTest := float32(200)
	maxRealTotal := float32(0)
	listBuyPrices := make([]float32, 0)
	min, max := float32(1), float32(30)
	for maxRealTotal < maxTotalTest {
		rand.Seed(time.Now().UnixNano())
		price :=min + rand.Float32() * (max - min)
		listBuyPrices = append(listBuyPrices, price)
		maxRealTotal += price
	}
	listBuySuccess := make([]model.TransactionWager, 0, len(listBuyPrices))
	cBuySuccess := make(chan model.TransactionWager)
	var waitGroup sync.WaitGroup
	go func() {
		waitGroup.Wait()
		close(cBuySuccess)
	}()
	for _, v := range listBuyPrices {
		price := v
		waitGroup.Add(1)
		go func() {
			defer func() {
				waitGroup.Done()
			}()
			transaction, err := suite.wagerRepo.Buy(wager.Id, price)
			if err == nil && transaction != nil {
				cBuySuccess <- *transaction
			}
		}()
	}

	for transaction := range cBuySuccess {
		listBuySuccess =append(listBuySuccess, transaction)
	}
	totalBought := float32(0)
	for _, t := range listBuySuccess {
		totalBought += t.BuyingPrice
	}
	totalBought = float32(math.Round(float64(totalBought)*100)/100)
	wagerDB, err := suite.wagerRepo.Get(wager.Id)
	if err != nil {
		suite.Fail(err.Error())
	}
	suite.NotNil(wagerDB.AmountSold, "the amount sold should have value after having transaction")
	suite.Equal(*wagerDB.AmountSold, totalBought, fmt.Sprintf("the total of transaction should be equal to wager selling price"))
}

func (suite *WagerTestSuite) TestWager_List() {
	offset, limit := 0, 100
	wagers, err := suite.wagerRepo.List(offset, limit)
	suite.Nil(err, fmt.Sprintf("suite.wagerRepo List has error with offset %v limit %v", offset, limit))
	suite.NotEmpty(wagers)
}

func initValue() model.Wager {
	return model.Wager{
		TotalWagerValue:     100,
		Odds:                10,
		SellingPercentage:   80,
		SellingPrice:        100,
		CurrentSellingPrice: 100,
		PlaceAt:             time.Time{},
	}
}