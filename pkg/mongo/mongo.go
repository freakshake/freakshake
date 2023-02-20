package mongo

import (
	"context"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectionString struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

func (d ConnectionString) String() string {
	return "mongodb://" + net.JoinHostPort(d.Host, d.Port)
}

func NewClient(s ConnectionString) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := mongo.Connect(ctx, options.Client().ApplyURI(s.String()))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func Ping(c *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := c.Ping(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
