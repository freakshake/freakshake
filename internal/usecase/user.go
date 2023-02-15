package usecase

import (
	"context"

	"github.com/mehdieidi/storm/internal/domain"
	"github.com/mehdieidi/storm/pkg/cache"
	"github.com/mehdieidi/storm/pkg/type/id"
	"github.com/mehdieidi/storm/pkg/type/offlim"
	"github.com/mehdieidi/storm/pkg/update"
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
	hashedPassword, err := u.Password.Hash()
	if err != nil {
		return domain.User{}, err
	}
	u.HashedPassword = string(hashedPassword)

	u.ID, err = s.userPostgres.Store(ctx, u)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (s *user) Get(ctx context.Context, uID id.ID[domain.User]) (domain.User, error) {
	return s.userPostgres.Find(ctx, uID)
}

func (s *user) List(ctx context.Context, o offlim.Offset, l offlim.Limit) ([]domain.User, error) {
	return s.userPostgres.FindAll(ctx, o, l)
}

func (s *user) Update(ctx context.Context, uID id.ID[domain.User], newUser domain.User) error {
	oldUser, err := s.userPostgres.Find(ctx, uID)
	if err != nil {
		return err
	}

	oldUser.FirstName = update.IfChanged(oldUser.FirstName, newUser.FirstName)
	oldUser.LastName = update.IfChanged(oldUser.LastName, newUser.LastName)
	oldUser.Email = update.IfChanged(oldUser.Email, newUser.Email)
	oldUser.MobileNumber = update.IfChanged(oldUser.MobileNumber, newUser.MobileNumber)
	oldUser.HashedPassword = update.IfChanged(oldUser.HashedPassword, newUser.HashedPassword)
	oldUser.Avatar = update.IfNilChanged(oldUser.Avatar, newUser.Avatar)

	if err := s.userPostgres.Update(ctx, uID, oldUser); err != nil {
		return err
	}

	return nil
}

func (s *user) Delete(ctx context.Context, uID id.ID[domain.User]) error {
	return s.userPostgres.Delete(ctx, uID)
}
