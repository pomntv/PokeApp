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
	ID      int    `json:"id,omitempty" bson:"id,omitempty"`
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
	HP      int    `json:"hp,omitempty" bson:"hp,omitempty"`
	Attack  int    `json:"attack,omitempty" bson:"attack,omitempty"`
	Defense int    `json:"defense,omitempty" bson:"defense,omitempty"`
	SpAtk   int    `json:"special attack,omitempty" bson:"special attack,omitempty"`
	SpDef   int    `json:"special defense,omitempty" bson:"special defense,omitempty"`
	Speed   int    `json:"speed,omitempty" bson:"speed,omitempty"`
	// Add other fields as needed
}

type Pokemondef struct {
	ID       int     `json:"id" bson:"id"`
	Name     string  `json:"name" bson:"name"`
	Type1    string  `json:"type 1" bson:"type 1"`
	Type2    string  `json:"type 2" bson:"type 2"`
	Normal   float64 `json:"normal" bson:"normal"`
	Fire     float64 `json:"fire" bson:"fire"`
	Water    float64 `json:"water" bson:"water"`
	Grass    float64 `json:"grass" bson:"grass"`
	Electric float64 `json:"electric" bson:"electric"`
	Ice      float64 `json:"ice" bson:"ice"`
	Fighting float64 `json:"fighting" bson:"fighting"`
	Poison   float64 `json:"poison" bson:"poison"`
	Ground   float64 `json:"ground" bson:"ground"`
	Flying   float64 `json:"flying" bson:"flying"`
	Psychic  float64 `json:"psychic" bson:"psychic"`
	Bug      float64 `json:"bug" bson:"bug"`
	Rock     float64 `json:"rock" bson:"rock"`
	Ghost    float64 `json:"ghost" bson:"ghost"`
	Dragon   float64 `json:"dragon" bson:"dragon"`
	Dark     float64 `json:"dark" bson:"dark"`
	Steel    float64 `json:"steel" bson:"steel"`
	Fairy    float64 `json:"fairy" bson:"fairy"`
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
