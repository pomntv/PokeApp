package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type AbilityDetail struct {
	IsHidden bool `json:"is_hidden"`
	Ability  struct {
		Name string `json:"name"`
	} `json:"ability"`
}

type Pokemon struct {
	Name      string          `json:"name"`
	Abilities []AbilityDetail `json:"abilities"`
}

func main() {
	// Create a new CSV file
	file, err := os.Create("pokemon_abilities.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	if err := writer.Write([]string{"ID", "Name", "Ability 1", "Ability 2", "Hidden Ability"}); err != nil {
		panic(err)
	}

	// Loop through Pokémon IDs
	for id := 1; id <= 1010; id++ {
		resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(id))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var pokeData Pokemon
		json.Unmarshal(body, &pokeData)

		ability1 := ""
		ability2 := ""
		hiddenAbility := ""

		for _, ability := range pokeData.Abilities {
			if ability.IsHidden {
				hiddenAbility = ability.Ability.Name
			} else {
				if ability1 == "" {
					ability1 = ability.Ability.Name
				} else {
					ability2 = ability.Ability.Name
				}
			}
		}

		// Write to CSV
		if err := writer.Write([]string{strconv.Itoa(id), pokeData.Name, ability1, ability2, hiddenAbility}); err != nil {
			panic(err)
		}
		fmt.Println("Processed Pokémon ID", id)

	}
	fmt.Println("CSV file generated successfully!")

}
