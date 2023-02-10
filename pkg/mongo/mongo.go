package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectionString struct {
	Host     string
	Port     string
	User     string
	DB       string
	Password string
}

func (d ConnectionString) String() string {
	return fmt.Sprintf("mongodb://%s:%s",
		d.Host,
		d.Port,
	)
}

func NewClient(d ConnectionString) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := mongo.Connect(ctx, options.Client().ApplyURI(d.String()))
	if err != nil {
		panic(err)
	}

	return c
}

func Ping(c *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := c.Ping(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
