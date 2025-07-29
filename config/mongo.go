package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

var DB *mongo.Database

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Configurações com monitoramento OpenTelemetry
	clientOpts := options.Client().
		ApplyURI("mongodb://mongo:27017").
		SetMonitor(otelmongo.NewMonitor())

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		panic(err)
	}

	// Testa a conexão
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	DB = client.Database("devopsdb")
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
