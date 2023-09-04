package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type PokemonStat struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
}

func fetchPokemonStatData(id int) (*PokemonStat, error) {
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pokemonStat PokemonStat
	if err := json.NewDecoder(resp.Body).Decode(&pokemonStat); err != nil {
		return nil, err
	}
	return &pokemonStat, nil
}

func main() {
	file, err := os.Create("pokemon_stats.csv")
	if err != nil {
		fmt.Println("Could not create file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Name", "HP", "Attack", "Defense", "Special Attack", "Special Defense", "Speed"}
	writer.Write(header)

	for id := 1; id <= 1010; id++ {
		pokemonStat, err := fetchPokemonStatData(id)
		if err != nil {
			fmt.Println("Error fetching data for Pokémon ID", id, ":", err)
			continue
		}

		row := []string{strconv.Itoa(pokemonStat.ID), pokemonStat.Name}
		stats := make(map[string]int)

		for _, stat := range pokemonStat.Stats {
			stats[stat.Stat.Name] = stat.BaseStat
		}

		row = append(row,
			strconv.Itoa(stats["hp"]),
			strconv.Itoa(stats["attack"]),
			strconv.Itoa(stats["defense"]),
			strconv.Itoa(stats["special-attack"]),
			strconv.Itoa(stats["special-defense"]),
			strconv.Itoa(stats["speed"]),
		)

		writer.Write(row)
		fmt.Println("Processed Pokémon ID", id)
	}

	fmt.Println("CSV file generated successfully!")
}
