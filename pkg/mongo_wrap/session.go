package mongo_wrap

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	client *mongo.Client
}

func NewClient(url string) (*Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	return &Client{client}, err
}

func (c *Client) GetCollection(db string, col string) *mongo.Collection {
	return c.client.Database(db).Collection(col)
}

func (c *Client) Close() error {
	if c.client != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := c.client.Disconnect(ctx)
		return err
	}
	return nil
}

func (c *Client) DropDatabase(db string) error {
	if c.client != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		return c.client.Database(db).Drop(ctx)
	}
	return nil
}
