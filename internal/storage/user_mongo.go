package storage

import (
	"context"

	"github.com/freakshake/logger"
	"github.com/freakshake/type/email"
	"github.com/freakshake/type/id"
	"github.com/freakshake/type/offlim"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/freakshake/internal/domain"
)

type userMongo struct {
	c      *mongo.Client
	logger logger.Logger
}

func NewUserMongoStorage(c *mongo.Client, logger logger.Logger) domain.UserStorage {
	return &userMongo{
		c:      c,
		logger: logger,
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
