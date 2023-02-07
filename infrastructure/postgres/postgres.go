package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DSN struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

func (d DSN) String() string {
	return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		d.Host,
		d.Port,
		d.User,
		d.DB,
		d.Password,
	)
}

func Open(d DSN) (*sql.DB, error) {
	return sql.Open("postgres", d.String())
}
