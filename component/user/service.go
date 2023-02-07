package user

import (
	"context"

	"github.com/mehdieidi/storm/model"
	"github.com/mehdieidi/storm/pkg/type/offlim"
)

type service struct {
	userPostgresStorage model.UserStorage
}

func NewService(userPostgresStorage model.UserStorage) model.UserService {
	return &service{
		userPostgresStorage: userPostgresStorage,
	}
}

func (s *service) Create(ctx context.Context, u model.User) (model.User, error) {
	return model.User{}, nil
}
func (s *service) Get(ctx context.Context, id model.UserID) (model.User, error) {
	return model.User{}, nil
}
func (s *service) List(ctx context.Context, o offlim.Offset, l offlim.Limit) ([]model.User, error) {
	return nil, nil
}
func (s *service) Update(ctx context.Context, id model.UserID, u model.User) error {
	return nil
}
func (s *service) Delete(ctx context.Context, id model.UserID) error {
	return nil
}
