package tools
import (
	"time"
)

type mockDB struct{
	
}


var mockLoginDetails = map[string]LoginDetails{
	"alex":{
		
		Username:  "alex",
		AuthToken: "123abc",
	},
	"john": {
		Username:  "john",
		AuthToken: "xyz789",
	},
	"mary": {
		Username:  "mary",
		AuthToken: "abc123",
	},
	"emma": {
		Username:  "emma",
		AuthToken: "def456",
	},
	"lisa": {
		Username:  "lisa",
		AuthToken: "ghi789",
	},
	"charlie": {
		Username:  "charlie",
		AuthToken: "jkl012",
	},
	
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:   100,
		Username: "alex",
	},
	"john": {
		Coins:   200,
		Username: "john",
	},
	"mary": {
		Coins:   150,
		Username: "mary",
	},
	"emma": {
		Coins:   300,
		Username: "emma",
	},
	"lisa": {
		Coins:   250,
		Username: "lisa",
	},
	"charlie": {
		Coins:   400,
		Username: "charlie",
	},
	
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData LoginDetails
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	coinData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &coinData
}

func (d *mockDB) SetupDatabase() error {

	return nil
}
