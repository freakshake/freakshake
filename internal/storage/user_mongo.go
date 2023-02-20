package storage

import (
	"context"

	"github.com/mehdieidi/storm/internal/domain"
	"github.com/mehdieidi/storm/pkg/type/email"
	"github.com/mehdieidi/storm/pkg/type/id"
	"github.com/mehdieidi/storm/pkg/type/offlim"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMongo struct {
	c *mongo.Client
}

func NewUserMongoStorage(c *mongo.Client) domain.UserStorage {
	return &userMongo{
		c: c,
	}
}

func (s *userMongo) Store(ctx context.Context, u domain.User) (id.ID[domain.User], error) {
	// TODO
	panic("todo")
}

func (s *userMongo) Find(ctx context.Context, uID id.ID[domain.User]) (domain.User, error) {
	// TODO
	panic("todo")
}

func (s *userMongo) FindByEmail(ctx context.Context, e email.Email) (domain.User, error) {
	// TODO
	panic("todo")
}

func (s *userMongo) FindAll(ctx context.Context, o offlim.Offset, l offlim.Limit) ([]domain.User, error) {
	// TODO
	panic("todo")
}

func (s *userMongo) Update(ctx context.Context, uID id.ID[domain.User], u domain.User) error {
	// TODO
	panic("todo")
}

func (s *userMongo) Delete(ctx context.Context, uID id.ID[domain.User]) error {
	// TODO
	panic("todo")
}
