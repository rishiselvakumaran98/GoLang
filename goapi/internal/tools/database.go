package tools

import (
	log "github.com/sirupsen/logrus"
)


// We can define a database interface so that it is very easy to swap our database endpoint when
// we need to change the database endpoint

type LoginDetails struct {
	AuthToken string
	Username string
}

type CoinDetails struct {
	Coins int64
	Username string
}

type DatabaseInterface interface{
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error){
	var database DatabaseInterface = &mockDB{}

	var err = database.SetupDatabase()
	if err != nil{
		log.Error(err)
		return nil, err
	}

	return &database, nil
}