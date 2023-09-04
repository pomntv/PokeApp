package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type PokemonMove struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
		} `json:"move"`
	} `json:"moves"`
}

func fetchPokemonMoveData(id int) (*PokemonMove, error) {
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pokemonMove PokemonMove
	if err := json.NewDecoder(resp.Body).Decode(&pokemonMove); err != nil {
		return nil, err
	}
	return &pokemonMove, nil
}

func main() {
	file, err := os.Create("pokemon_moves.csv")
	if err != nil {
		fmt.Println("Could not create file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Name", "Moves"}
	writer.Write(header)

	for id := 1; id <= 1010; id++ {
		pokemonMove, err := fetchPokemonMoveData(id)
		if err != nil {
			fmt.Println("Error fetching data for Pokémon ID", id, ":", err)
			continue
		}

		moves := []string{}
		for _, move := range pokemonMove.Moves {
			moves = append(moves, move.Move.Name)
		}
		moveList := fmt.Sprintf("\"%s\"", strings.Join(moves, ", "))

		row := []string{strconv.Itoa(pokemonMove.ID), pokemonMove.Name, moveList}
		writer.Write(row)
		fmt.Println("Processed Pokémon ID", id)
	}

	fmt.Println("CSV file generated successfully!")
}
