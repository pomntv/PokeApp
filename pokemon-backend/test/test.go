package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set up MongoDB connection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://pompoke:pompoke11501112@cluster0.nooipqq.mongodb.net/Data_Pokemon?retryWrites=true&w=majority")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Fetch a Pokémon by its name
	var result bson.M
	collection := client.Database("Data_Pokemon").Collection("stat_of_pokemon")
	// err = collection.FindOne(ctx, bson.M{"Name": "electrode"}).Decode(&result)
	err = collection.FindOne(ctx, bson.M{"ID": 900}).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	// Print the result
	fmt.Println(result)
}
