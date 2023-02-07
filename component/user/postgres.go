package user

import (
	"context"
	"database/sql"

	"github.com/mehdieidi/storm/model"
	"github.com/mehdieidi/storm/pkg/type/offlim"
)

type postgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) model.UserStorage {
	return &postgresStorage{
		db: db,
	}
}

func (s *postgresStorage) Store(context.Context, model.User) (model.UserID, error) {
	return 0, nil
}

func (s *postgresStorage) Find(context.Context, model.UserID) (model.User, error) {
	return model.User{}, nil
}

func (s *postgresStorage) FindAll(context.Context, offlim.Offset, offlim.Limit) ([]model.User, error) {
	return nil, nil
}

func (s *postgresStorage) Update(context.Context, model.UserID, model.User) error {
	return nil
}

func (s *postgresStorage) Delete(context.Context, model.UserID) error {
	return nil
}
