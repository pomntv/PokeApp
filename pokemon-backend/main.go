package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitializeService() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://pompoke:pompoke11501112@cluster0.nooipqq.mongodb.net/Data_Pokemon?retryWrites=true&w=majority")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
}

type Pokemon struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	// Add other fields as needed
}

func GetAllPokemon(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllPokemon function called")

	collection := client.Database("Data_Pokemon").Collection("stat_of_pokemon")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	var pokemons []Pokemon
	if err = cursor.All(context.TODO(), &pokemons); err != nil {
		http.Error(w, "Failed to parse data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pokemons)
}

func main() {
	InitializeService()
	defer client.Disconnect(context.Background())

	// Setting up the HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/api/pokemon/all", GetAllPokemon)

	// Handling CORS
	handler := cors.Default().Handler(mux)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", handler)
}
