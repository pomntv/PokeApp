package main

import (
	"context"
	"encoding/json"
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
	// ID   string `json:"_id,omitempty" bson:"_id,omitempty"`
	ID   int    `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	// Add other fields as needed
}

type Pokemondef struct {
	ID       int     `json:"id,omitempty" bson:"id,omitempty"`
	Name     string  `json:"name,omitempty" bson:"name,omitempty"`
	Normal   float64 `json:"normal,omitempty" bson:"normal,omitempty"`
	Fire     float64 `json:"fire,omitempty" bson:"fire,omitempty"`
	Water    float64 `json:"water,omitempty" bson:"water,omitempty"`
	Grass    float64 `json:"grass,omitempty" bson:"grass,omitempty"`
	Electric float64 `json:"electric,omitempty" bson:"electric,omitempty"`
	Ice      float64 `json:"ice,omitempty" bson:"ice,omitempty"`
	Fighting float64 `json:"fighting,omitempty" bson:"fighting,omitempty"`
	Poison   float64 `json:"poison,omitempty" bson:"poison,omitempty"`
	Ground   float64 `json:"ground,omitempty" bson:"ground,omitempty"`
	Flying   float64 `json:"flying,omitempty" bson:"flying,omitempty"`
	Psychic  float64 `json:"psychic,omitempty" bson:"psychic,omitempty"`
	Bug      float64 `json:"bug,omitempty" bson:"bug,omitempty"`
	Rock     float64 `json:"rock,omitempty" bson:"rock,omitempty"`
	Ghost    float64 `json:"ghost,omitempty" bson:"ghost,omitempty"`
	Dragon   float64 `json:"dragon,omitempty" bson:"dragon,omitempty"`
	Dark     float64 `json:"dark,omitempty" bson:"dark,omitempty"`
	Steel    float64 `json:"steel,omitempty" bson:"steel,omitempty"`
	Fairy    float64 `json:"fairy,omitempty" bson:"fairy,omitempty"`
}

func GetAllPokemon(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("Data_Pokemon").Collection("stat_of_pokemon")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Error fetching data from MongoDB:", err) // Log the detailed error
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

func GetAllPokemonDefensive(w http.ResponseWriter, r *http.Request) {
	collection := client.Database("Data_Pokemon").Collection("defensive_of_pokemon")

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Error fetching data from MongoDB:", err) // Log the detailed error
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	var pokemons []Pokemondef
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
	mux.HandleFunc("/api/pokemon", GetAllPokemon)
	mux.HandleFunc("/api/pokemondefensive", GetAllPokemonDefensive)
	// Handling CORS
	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":8080", handler)
}
