package btce

import (
	"errors"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"strconv"
	"testing"
	"time"
)

const SLEEP = 5

func TestAccountInfo(t *testing.T) {
	time.Sleep(SLEEP * time.Second)
	Convey("Account information data", t, func() {
		info, err := GetAccountInfoAuth(os.Getenv("API_KEY_TEST"), os.Getenv("API_SECRET_TEST"))
		Convey("No error should occur", func() {
			So(err, ShouldBeNil)
		})
		Convey("Account information fields should be returned", func() {
			So(info, ShouldHaveSameTypeAs, AccountInfo{})
			So(info.TransactionCount, ShouldEqual, 0)
			So(info.ServerTime, ShouldBeGreaterThan, 0)
			So(info.Funds["btc"], ShouldBeGreaterThanOrEqualTo, 0)
		})
	})

}

func TestActiveOrders(t *testing.T) {
	time.Sleep(SLEEP * time.Second)
	Convey("Active orders data", t, func() {
		orders, err := GetActiveOrdersAuth(os.Getenv("API_KEY_TEST"), os.Getenv("API_SECRET_TEST"), "btc_usd")

		if err != nil {
			Convey("If error is returned, it should be 'no orders'", func() {
				So(err, ShouldResemble, errors.New("no orders"))
			})
		} else {
			fmt.Println(orders)
			Convey("If no error is returned, 'order' should have length", func() {
				So(len(orders), ShouldBeGreaterThanOrEqualTo, 1)
			})
		}
	})
}

func TestOrderTrade(t *testing.T) {
	time.Sleep(SLEEP * time.Second)
	Convey("Trade new order", t, func() {
		orderResponse, err := OrderTradeAuth(os.Getenv("API_KEY_TEST"), os.Getenv("API_SECRET_TEST"), "btc_usd", "buy", 900, 1)

		if err != nil {
			Convey("If error is returned, it should be 'not enough USD'", func() {
				So(err, ShouldResemble, errors.New("It is not enough USD for purchase"))
			})
		} else {
			Convey("If no error is returned, 'btc_usd' amount should appear", func() {
				So(orderResponse.OrderID, ShouldBeGreaterThan, 0)
			})
		}
	})
}

func TestOrderInfo(t *testing.T) {
	time.Sleep(SLEEP * time.Second)
	orderID := "1140186935"
	Convey("Get order info", t, func() {
		orderResponse, err := GetOrderInfoAuth(os.Getenv("API_KEY_TEST"), os.Getenv("API_SECRET_TEST"), orderID)

		if err != nil {
			Convey("If error is returned, it should be 'invalid order'", func() {
				So(err, ShouldResemble, errors.New("invalid order"))
			})
		} else {
			Convey("If no error is returned, order information should be returned", func() {
				So(orderResponse[orderID], ShouldNotBeNil)
				So(orderResponse[orderID].Amount, ShouldBeGreaterThanOrEqualTo, 0)
			})
		}
	})
}

func TestCancelOrder(t *testing.T) {
	time.Sleep(SLEEP * time.Second)
	orderID := "1140186997"
	Convey("Cancel order", t, func() {
		orderResponse, err := CancelOrderAuth(os.Getenv("API_KEY_TEST"), os.Getenv("API_SECRET_TEST"), orderID)

		if err != nil {
			Convey("If error is returned, it should be 'bad status'", func() {
				So(err, ShouldResemble, errors.New("bad status"))
			})
		} else {
			Convey("If no error is returned, same order id should be returned", func() {
				So(strconv.Itoa(orderResponse.OrderID), ShouldEqual, orderID)
			})
		}
	})
}

func TestTradeHistory(t *testing.T) {
	time.Sleep(SLEEP * time.Second)
	Convey("Trade history data", t, func() {

		filter := HistoryFilter{}
		tradeHistory, err := GetTradeHistoryAuth(os.Getenv("API_KEY_TEST"), os.Getenv("API_SECRET_TEST"), filter, "btc_usd")

		Convey("No error should occur", func() {
			So(err, ShouldBeNil)

		})
		Convey("Trade history should be retrieved", func() {
			So(tradeHistory, ShouldNotBeNil)
			So(len(tradeHistory), ShouldBeGreaterThanOrEqualTo, 0)
		})
	})
}

func TestTransactionHistory(t *testing.T) {
	time.Sleep(SLEEP * time.Second)
	Convey("Transaction history data", t, func() {

		filter := HistoryFilter{}
		transactionHistory, err := GetTransactionHistoryAuth(os.Getenv("API_KEY_TEST"), os.Getenv("API_SECRET_TEST"), filter)

		Convey("No error should occur", func() {
			So(err, ShouldBeNil)

		})
		Convey("Transaction history should be retrieved", func() {
			So(transactionHistory, ShouldNotBeNil)
			So(len(transactionHistory), ShouldBeGreaterThanOrEqualTo, 0)
		})
	})
}

func TestWithdrawCoin(t *testing.T) {
	time.Sleep(SLEEP * time.Second)
	Convey("Withdraw coin", t, func() {

		response, err := WithdrawCoinAuth(os.Getenv("API_KEY_TEST"), os.Getenv("API_SECRET_TEST"), "BTC", 0.001, "address")

		if err != nil {
			Convey("If error is returned, it should be 'api permission'", func() {
				So(err, ShouldResemble, errors.New("api key dont have withdraw permission"))
			})
		} else {
			Convey("If no error is returned, withdraw reponse should be returned", func() {
				So(response.TransactionID, ShouldBeGreaterThan, 0)
				So(response.AmountSent, ShouldBeGreaterThanOrEqualTo, 0)
			})
		}
	})
}

func TestCreateCoupon(t *testing.T) {
	time.Sleep(SLEEP * time.Second)
	Convey("Create coupon", t, func() {

		response, err := CreateCouponAuth(os.Getenv("API_KEY_TEST"), os.Getenv("API_SECRET_TEST"), "BTC", 0.001)

		if err != nil {
			Convey("If error is returned, it should be 'api permission'", func() {
				So(err, ShouldResemble, errors.New("api key dont have coupon permission"))
			})
		} else {
			Convey("If no error is returned, withdraw reponse should be returned", func() {
				So(response.Coupon, ShouldNotBeBlank)
				So(response.TransactionID, ShouldBeGreaterThan, 0)
			})
		}
	})
}

func TestRedeemCoupon(t *testing.T) {
	time.Sleep(SLEEP * time.Second)
	Convey("Redeem coupon", t, func() {

		response, err := RedeemCouponAuth(os.Getenv("API_KEY_TEST"), os.Getenv("API_SECRET_TEST"), "BTC-USD-XYZ")

		if err != nil {
			Convey("If error is returned, it should be 'api permission'", func() {
				So(err, ShouldResemble, errors.New("api key dont have coupon permission"))
			})
		} else {
			Convey("If no error is returned, withdraw reponse should be returned", func() {
				So(response.CouponCurrency, ShouldNotBeBlank)
				So(response.TransactionID, ShouldBeGreaterThan, 0)
			})
		}
	})
}