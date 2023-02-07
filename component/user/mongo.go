package user

import (
	"context"

	"github.com/mehdieidi/storm/model"
	"github.com/mehdieidi/storm/pkg/type/offlim"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoStorage struct {
	c *mongo.Client
}

func NewMongoStorage(c *mongo.Client) model.UserStorage {
	return &mongoStorage{
		c: c,
	}
}

func (s *mongoStorage) Store(context.Context, model.User) (model.UserID, error) {
	return 0, nil
}

func (s *mongoStorage) Find(context.Context, model.UserID) (model.User, error) {
	return model.User{}, nil
}

func (s *mongoStorage) FindAll(context.Context, offlim.Offset, offlim.Limit) ([]model.User, error) {
	return nil, nil
}

func (s *mongoStorage) Update(context.Context, model.UserID, model.User) error {
	return nil
}

func (s *mongoStorage) Delete(context.Context, model.UserID) error {
	return nil
}
