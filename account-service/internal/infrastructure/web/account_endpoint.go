package account_service

import (
	"net/http"

	service "github.com/ernesto2108/helpnwn-service/account-service/internal/domain/ports"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
}

func SignUp(w http.ResponseWriter, r *http.Request) {
}

type UserEndPoint struct {
	service.AccountService
}