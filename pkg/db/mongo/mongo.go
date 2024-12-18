package mongo

import (
	"context"
	"fmt"
	"telegram-clicker-game-be/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInterface interface{}

func NewMongoInstance(cfg *config.Config) (coll *mongo.Database, client *mongo.Client, err error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoUri := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=%s",
		cfg.ClickerGameDatabase.User,
		cfg.ClickerGameDatabase.Password,
		cfg.ClickerGameDatabase.Host,
		cfg.ClickerGameDatabase.Database)
	opts := options.
		Client().
		ApplyURI(mongoUri).
		SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		return
	}

	coll = client.Database(cfg.ClickerGameDatabase.Database)
	err = PingDB(client)
	return
}

func PingDB(client *mongo.Client) error {
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return err
	}

	return nil
}
