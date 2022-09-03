package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGO_URL  string = fmt.Sprintf("mongodb://mongo:%s/admin?authSource=admin&retryWrites=true&w=majority", os.Getenv("MONGO_PORT"))
	MONGO_USER string = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	MONGO_PASS string = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URL).SetAuth(options.Credential{
		Username: MONGO_USER,
		Password: MONGO_PASS,
	}))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongoDB")
	return client
}
