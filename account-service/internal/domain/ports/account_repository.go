package account_service

import (
	persistence "github.com/ernesto2108/helpnwn-service/account-service/internal/application/persistence"
)

type AccountRepository interface {
	SignUp()
	SingIn()
}

type AccountService struct {
	persistence.PostgresAccountRepository
}
