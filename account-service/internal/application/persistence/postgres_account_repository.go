package account_service

import "database/sql"

type PostgresAccountRepository interface {
	signUp()
	singIn()
}

type PsqlService struct {
	*sql.DB
}
