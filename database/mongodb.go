package database

import (
	"context"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	Username string
	Password string
	Hostname string
}

var (
	endpoint string
)

func NewMongoDbConnection(ctx context.Context, connection *MongoDBConfig) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	endpoint = fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=Cluster0", connection.Username, connection.Password, connection.Hostname)
	opts := options.Client().ApplyURI(endpoint).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		slog.Error("failed to make a connection to mongodb", "error", err)
		return nil, err
	}
	return client, nil

}

func PingDatabase(ctx context.Context, client *mongo.Client) error {
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(err)
	}
	return nil
}
