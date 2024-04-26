package tools

import (
	"time"
	"fmt"
)

type mockDB struct{}

var AuthDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "randomUser",
		Username: "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username: "jason",
	},
	"ben": {
		AuthToken: "hello-world",
		Username: "ben",
	},

}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"jason": {
		Coins: 300,
		Username: "jason",
	},
	"ben": {
		Coins: 500,
		Username: "ben",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails{
	// mock Database call
	time.Sleep(1*time.Second)
	var clientData = LoginDetails{}
	clientData, ok := AuthDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

// These function get implemented under the mockDB type
func (d *mockDB) GetUserCoins(username string) *CoinDetails{
	// mock Database call
	time.Sleep(1*time.Second)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	} 
	return &clientData
}

func (d *mockDB) SetupDatabase() error{
	return nil
}