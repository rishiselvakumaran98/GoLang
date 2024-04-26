package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/rishiselvakumaran98/goapi/internal/tools"
	"github.com/rishiselvakumaran98/goapi/api"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request){
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()

	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil{
		log.Error((err))
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil{
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}