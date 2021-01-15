package mongodb

import (
	"context"

	"github.com/jgolang/config"
	"github.com/jgolang/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewClient mongodb
func NewClient() mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetString("database.mongo_url_connect")))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return *client
}
