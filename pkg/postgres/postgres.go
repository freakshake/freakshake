package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DSN struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

func (d DSN) String() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		d.Host,
		d.Port,
		d.User,
		d.Password,
		d.DB,
	)
}

func Open(d DSN) (*sql.DB, error) {
	return sql.Open("postgres", d.String())
}

func MigrateUp(db *sql.DB, migrationsPath string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return err
		}
	}

	return nil
}
