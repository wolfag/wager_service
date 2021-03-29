package error

var ErrBuyingGreaterThanCurrentSell = NewServiceError("the buying price is greater than current selling price")
var ErrInvalidData = NewServiceError("invalid data")
var ErrWagerNotFound = NewServiceError("wager is not found")
var ErrWagerNotFoundOrCantNotBuy = NewServiceError("wager is not found or can't buy")

var ErrTotalWagerValue = NewServiceError("total wager value must be specified as a positive integer above 0")
var ErrOddsValue = NewServiceError("Odds must be specified as a positive integer above 0")
var ErrSellingPercentage = NewServiceError("selling percentage must be specified as an integer between 1 and 100")
var ErrSellingPrice = NewServiceError("selling price must be greater than total_wager_value * (selling_percentage / 100)")