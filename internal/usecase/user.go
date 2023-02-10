package usecase

import (
	"context"

	"github.com/mehdieidi/storm/internal/domain"
	"github.com/mehdieidi/storm/pkg/cache"
	"github.com/mehdieidi/storm/pkg/type/offlim"
)

type user struct {
	userPostgres domain.UserStorage
	userMongo    domain.UserStorage
	cache        cache.Cache
}

func NewUserService(
	userPostgres domain.UserStorage,
	userMongo domain.UserStorage,
	cache cache.Cache,
) domain.UserService {
	return &user{
		userPostgres: userPostgres,
		userMongo:    userMongo,
		cache:        cache,
	}
}

func (s *user) Create(ctx context.Context, u domain.User) (domain.User, error) {
	// TODO
	return domain.User{}, nil
}

func (s *user) Get(ctx context.Context, uID domain.UserID) (domain.User, error) {
	// TODO
	return domain.User{}, nil
}

func (s *user) List(ctx context.Context, o offlim.Offset, l offlim.Limit) ([]domain.User, error) {
	// TODO
	return nil, nil
}

func (s *user) Update(ctx context.Context, uID domain.UserID, u domain.User) error {
	// TODO
	return nil
}

func (s *user) Delete(ctx context.Context, uID domain.UserID) error {
	// TODO
	return nil
}
