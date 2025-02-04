package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	client, err := mongo.Connect(ctx, cOptions)

	if err != nil {
		panic("DB error!")
	}

	Client = client
	fmt.Println("Connection success!")

}

func GetCollection(databasename, collectionname string) *mongo.Collection {
	return Client.Database(databasename).Collection(collectionname)
}
