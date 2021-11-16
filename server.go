package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoTest() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Errorf("client establish failed. err: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err == nil {
		fmt.Println("connect to db success.")
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	db := client.Database("Mongo")
	collectionNames, err := db.ListCollectionNames(ctx, bson.M{})
	fmt.Println("collectionNames:", collectionNames)
}

func main() {
	// e := echo.New()
	// router.RouterInit(e)
	// e.Logger.Fatal(e.Start(":2222"))
	mongoTest()
}
