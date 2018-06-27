package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

func main() {
	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("baz").Collection("qux")
	res, err := collection.InsertOne(context.Background(), map[string]string{"hello": "world"})
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID

	fmt.Println(id)
}
