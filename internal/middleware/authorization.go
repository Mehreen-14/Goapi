package middleware

import (
	"errors"
	"net/http"

	"github.com/Mehreen-14/Goapi.git/api"
	"github.com/Mehreen-14/Goapi/api"
	"github.com/Mehreen-14/Goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorized = errors.New("unauthorized access")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")

		var err error
		
		if username == "" || token == "" {
			log.Error(UnAuthorized)
			api.RequestErrorHandler(w, UnAuthorized)
			return
		}
		var database *tools.DatabaseInterface
		database,err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}
		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		if(loginDetails == nil || (token != (*loginDetails).Authotken)) {
			log.Error(UnAuthorized)
			api.RequestErrorHandler(w, UnAuthorized)
			return
		}

		next.ServeHTTP(w, r)
		
	})
}