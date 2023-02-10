package repository

import (
	"context"

	"github.com/mehdieidi/storm/internal/domain"
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

func (s *userMongo) Store(ctx context.Context, u domain.User) (domain.UserID, error) {
	// TODO
	return 0, nil
}

func (s *userMongo) Find(ctx context.Context, uID domain.UserID) (domain.User, error) {
	// TODO
	return domain.User{}, nil
}

func (s *userMongo) FindAll(ctx context.Context, o offlim.Offset, l offlim.Limit) ([]domain.User, error) {
	// TODO
	return nil, nil
}

func (s *userMongo) Update(ctx context.Context, uID domain.UserID, u domain.User) error {
	// TODO
	return nil
}

func (s *userMongo) Delete(ctx context.Context, uID domain.UserID) error {
	// TODO
	return nil
}