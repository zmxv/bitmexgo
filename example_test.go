package bitmexgo

import (
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	// Get your API key/secret pair at https://www.bitmex.com/app/apiKeys
	testApiKey := os.Getenv("TEST_API_KEY")
	testApiSecret := os.Getenv("TEST_API_SECRET_KEY")

	// Create an authentication context
	auth := NewAPIKeyContext(testApiKey, testApiSecret)

	// Create a shareable API client instance
	// apiClient := bitmexgo.NewAPIClient(bitmexgo.NewConfiguration())

	// Create a testnet API client instance
	testnetClient := NewAPIClient(NewTestnetConfiguration())

	// Call APIs without parameters by passing the auth context.
	// e.g. getting exchange-wide turnover and volume statistics:
	// stats, res, err := apiClient.StatsApi.StatsGet(auth)

	// Call APIs with default parameters by passing the auth context and a nil.
	// e.g. getting all open positions:
	// pos, res, err := apiClient.PositionApi.PositionGet(auth, nil)

	// Call APIs with additional parameters by constructing a corresponding XXXOpts struct.
	// e.g. submitting a limit order to buy 100 contracts of XBTUSD at $8000.5 and to sell 50 contracts of XBTUSD at 16000:
	var params []*OrderNewOpts

	var paramBuy OrderNewOpts
	paramBuy.OrdType.Set("Limit")
	paramBuy.Side.Set("Buy")
	paramBuy.OrderQty.Set(100)
	paramBuy.Price.Set(8000.5)

	params = append(params, &paramBuy)

	var paramSell OrderNewOpts
	paramSell.OrdType.Set("Limit")
	paramSell.Side.Set("Sell")
	paramSell.OrderQty.Set(50)
	paramSell.Price.Set(16000)

	params = append(params, &paramSell)

	order, res, err := testnetClient.OrderApi.OrderNewBulk(auth, "XBTUSD", params)
	fmt.Println(order)
	fmt.Println(res)
	fmt.Println(err)
}
