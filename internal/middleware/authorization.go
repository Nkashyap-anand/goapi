package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"goapi/api"
	"goapi/internal/tools"
)

var UnAuthorizedError = errors.New("invalid username or token")

func Authoization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrrorhandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabse()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrrorhandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
