package repository

import (
	"context"
	"database/sql"

	"github.com/mehdieidi/storm/internal/domain"
	"github.com/mehdieidi/storm/pkg/type/offlim"
)

type userPostgres struct {
	db *sql.DB
}

func NewUserPostgresStorage(db *sql.DB) domain.UserStorage {
	return &userPostgres{
		db: db,
	}
}

func (s *userPostgres) Store(ctx context.Context, u domain.User) (domain.UserID, error) {
	// TODO
	return 0, nil
}

func (s *userPostgres) Find(ctx context.Context, uID domain.UserID) (domain.User, error) {
	// TODO
	return domain.User{}, nil
}

func (s *userPostgres) FindAll(ctx context.Context, o offlim.Offset, l offlim.Limit) ([]domain.User, error) {
	// TODO
	return nil, nil
}

func (s *userPostgres) Update(ctx context.Context, uID domain.UserID, u domain.User) error {
	// TODO
	return nil
}

func (s *userPostgres) Delete(ctx context.Context, uID domain.UserID) error {
	// TODO
	return nil
}
