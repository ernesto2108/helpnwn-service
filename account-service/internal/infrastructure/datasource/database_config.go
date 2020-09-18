package account_service

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once
var db *sql.DB
var errNotInitialized = errors.New("the connection pool has not been initialized")

const (
	// dns database connection
	dnsPsql      = "postgres://%s:%s@%s:%d/%s?sslmode=disable"

	// name database
	postgres  = "postgres"
)

type Model struct {
	Engine   string
	Server   string
	port     int
	User     string
	Password string
	Name     string
}

func (m *Model) NewConn() (*sql.DB, error) {

	var err error
	var	dns string

	if m.Engine == "" {
		return nil, errors.New("databases mandatory")
	}

	once.Do(func() {

		switch m.Engine {
		case postgres:
			dns = dnsPsql
		}

		db, err = m.initConn(dns)

	})

	return db, err
}

func (m *Model) initConn(dns string) (*sql.DB, error) {

	var err error

	d := fmt.Sprintf(dns, m.User, m.Password, m.Server, m.port, m.Name)
	db, err := sql.Open(m.Engine, d)

	if err != nil {
		return db, err
	}

	return db, nil
}

func InitConn() (*sql.DB, error) {

	if db == nil {
		return db, errNotInitialized
	}

	if db.Ping() != nil {
		return db, errNotInitialized
	}

	return db, nil

}

// close connection
func CloseConn() error {

	err := db.Close()

	if err != nil {
		return err
	}

	return nil
}
