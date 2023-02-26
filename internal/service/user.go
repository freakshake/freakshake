package service

import (
	"context"

	"github.com/mehdieidi/freakshake/internal/domain"
	"github.com/mehdieidi/freakshake/pkg/cache"
	"github.com/mehdieidi/freakshake/pkg/pick"
	"github.com/mehdieidi/freakshake/pkg/type/email"
	"github.com/mehdieidi/freakshake/pkg/type/id"
	"github.com/mehdieidi/freakshake/pkg/type/offlim"
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

func (s *user) GetByEmail(ctx context.Context, e email.Email) (domain.User, error) {
	return s.userPostgres.FindByEmail(ctx, e)
}

func (s *user) List(ctx context.Context, o offlim.Offset, l offlim.Limit) ([]domain.User, error) {
	return s.userPostgres.FindAll(ctx, o, l)
}

func (s *user) Update(ctx context.Context, uID id.ID[domain.User], newUser domain.User) error {
	oldUser, err := s.userPostgres.Find(ctx, uID)
	if err != nil {
		return err
	}

	oldUser.FirstName = pick.IfChanged(oldUser.FirstName, newUser.FirstName)
	oldUser.LastName = pick.IfChanged(oldUser.LastName, newUser.LastName)
	oldUser.Email = pick.IfChanged(oldUser.Email, newUser.Email)
	oldUser.MobileNumber = pick.IfChanged(oldUser.MobileNumber, newUser.MobileNumber)
	oldUser.HashedPassword = pick.IfChanged(oldUser.HashedPassword, newUser.HashedPassword)
	oldUser.Avatar = pick.IfSome(oldUser.Avatar, newUser.Avatar)

	if err := s.userPostgres.Update(ctx, uID, oldUser); err != nil {
		return err
	}

	return nil
}

func (s *user) Delete(ctx context.Context, uID id.ID[domain.User]) error {
	return s.userPostgres.Delete(ctx, uID)
}
