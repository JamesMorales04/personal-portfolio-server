package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConexion = ConnectBD()

func ConnectBD() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://" + os.Getenv("USER") + ":" + os.Getenv("PASSWORD") + "@" + os.Getenv("CLUSTER") + os.Getenv("DATABASE") + "?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func CheckConnection() bool {

	error := MongoConexion.Ping(context.TODO(), nil)

	return error == nil
}
