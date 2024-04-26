package middleware

import (
	"errors"
	"net/http"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/rishiselvakumaran98/goapi/internal/tools"
	"github.com/rishiselvakumaran98/goapi/api"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")		

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
		}

		var database *tools.DatabaseInterface 
		var err error
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)
		var errorStr string
		fmt.Println(errorStr)
		if (loginDetails == nil || token != (*loginDetails).AuthToken){
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w,r)
	})
}