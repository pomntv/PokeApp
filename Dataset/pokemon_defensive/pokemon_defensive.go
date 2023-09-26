package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Pokemon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

var types = []string{
	"Normal", "Fire", "Water", "Grass", "Electric", "Ice", "Fighting", "Poison",
	"Ground", "Flying", "Psychic", "Bug", "Rock", "Ghost", "Dragon", "Dark", "Steel", "Fairy",
}

var typeEffectiveness = map[string]map[string]float64{
	"Normal":   {"Fighting": 2, "Ghost": 0},
	"Fire":     {"Water": 2, "Ground": 2, "Rock": 2, "Fire": 0.5, "Grass": 0.5, "Ice": 0.5, "Bug": 0.5, "Steel": 0.5, "Fairy": 0.5},
	"Water":    {"Grass": 2, "Electric": 2, "Water": 0.5, "Fire": 0.5, "Ice": 0.5, "Steel": 0.5},
	"Grass":    {"Fire": 2, "Ice": 2, "Poison": 2, "Flying": 2, "Bug": 2, "Grass": 0.5, "Water": 0.5, "Ground": 0.5, "Electric": 0.5},
	"Electric": {"Ground": 2, "Electric": 0.5, "Flying": 0.5, "Steel": 0.5},
	"Ice":      {"Fire": 2, "Fighting": 2, "Rock": 2, "Steel": 2, "Ice": 0.5},
	"Fighting": {"Flying": 2, "Psychic": 2, "Fairy": 2, "Bug": 0.5, "Rock": 0.5, "Dark": 0.5},
	"Poison":   {"Ground": 2, "Psychic": 2, "Fighting": 0.5, "Poison": 0.5, "Bug": 0.5, "Grass": 0.5, "Fairy": 0.5},
	"Ground":   {"Water": 2, "Grass": 2, "Ice": 2, "Electric": 0, "Poison": 0.5, "Rock": 0.5},
	"Flying":   {"Electric": 2, "Ice": 2, "Rock": 2, "Grass": 0.5, "Fighting": 0.5, "Bug": 0.5, "Ground": 0, "Ghost": 0},
	"Psychic":  {"Bug": 2, "Ghost": 2, "Dark": 2, "Fighting": 0.5, "Psychic": 0.5},
	"Bug":      {"Flying": 2, "Rock": 2, "Fire": 2, "Grass": 0.5, "Fighting": 0.5, "Ground": 0.5},
	"Rock":     {"Water": 2, "Grass": 2, "Fighting": 2, "Ground": 2, "Steel": 2, "Normal": 0.5, "Fire": 0.5, "Poison": 0.5, "Flying": 0.5},
	"Ghost":    {"Ghost": 2, "Dark": 2, "Normal": 0, "Fighting": 0, "Poison": 0.5, "Bug": 0.5},
	"Dragon":   {"Ice": 2, "Dragon": 2, "Fairy": 2, "Fire": 0.5, "Water": 0.5, "Grass": 0.5, "Electric": 0.5},
	"Dark":     {"Fighting": 2, "Bug": 2, "Fairy": 2, "Ghost": 0.5, "Psychic": 0},
	"Steel":    {"Fire": 2, "Fighting": 2, "Ground": 2, "Normal": 0.5, "Grass": 0.5, "Ice": 0.5, "Flying": 0.5, "Psychic": 0.5, "Bug": 0.5, "Rock": 0.5, "Dragon": 0.5, "Steel": 0.5},
	"Fairy":    {"Poison": 2, "Steel": 2, "Fighting": 2, "Dark": 0.5, "Dragon": 0},
}

func main() {
	file, err := os.Create("pokemon_defensive_data.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header row
	header := append([]string{"ID", "Name", "Type 1", "Type 2"}, types...)
	writer.Write(header)

	// Loop through each Pokémon ID
	for id := 1; id <= 1010; id++ {
		pokemon, err := fetchPokemonData(id)
		if err != nil {
			fmt.Println("Error fetching data for Pokémon ID", id, ":", err)
			continue
		}

		row := []string{strconv.Itoa(pokemon.ID), pokemon.Name, "", ""}
		if len(pokemon.Types) > 0 {
			row[2] = pokemon.Types[0].Type.Name
		}
		if len(pokemon.Types) > 1 {
			row[3] = pokemon.Types[1].Type.Name
		}
		for _, attackingType := range types {
			effectiveness := 1.0
			for _, pokeType := range pokemon.Types {
				typeName := titleCase(pokeType.Type.Name) // Convert to title case
				effect, ok := typeEffectiveness[typeName][attackingType]
				if ok {
					effectiveness *= effect
				}
			}
			row = append(row, fmt.Sprintf("%.2f", effectiveness)) // Changed to "%.2f" to ensure two decimal places
		}

		writer.Write(row)
		fmt.Println("Processed Pokémon ID", id)
	}

	fmt.Println("CSV file generated successfully!")
}
func titleCase(str string) string {
	return strings.Title(strings.ToLower(str))
}
func fetchPokemonData(pokemonID int) (*Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", pokemonID)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}
