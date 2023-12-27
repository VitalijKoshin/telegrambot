package mongodb

import (
	"context"
	"time"

	"telegrambot/internal/config"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	Client *mongo.Client
}

func NewClient(cfg *config.Config) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credentials := options.Credential{
		Username: cfg.MongoDbUser,
		Password: cfg.MongoDbPass,
	}
	clientOptions := options.Client().ApplyURI(cfg.MongoDbUri).SetAuth(credentials)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &MongoClient{
		Client: client,
	}, nil
}

func (m *MongoClient) Disconnect() {
	err := m.Client.Disconnect(context.Background())
	if err != nil {
		logrus.Fatal(err)
	}
}

func (m *MongoClient) GetClient() *mongo.Client {
	return m.Client
}
