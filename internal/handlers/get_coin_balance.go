package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Mehreen-14/Goapi.git/api"
	"github.com/Mehreen-14/Goapi.git/internal/tools"
	"github.com/Mehreen-14/Goapi/api"
	"github.com/Mehreen-14/Goapi/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func getCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error 
	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: tokenDetails.Coins,
		Code: http.StatusOK,

	}

	w.Header().Set("Content-type","application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}