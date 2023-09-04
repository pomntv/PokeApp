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

type MoveDetail struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

func main() {
	// Create a new CSV file
	file, err := os.Create("moves.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header to CSV
	if err := writer.Write([]string{"ID", "Name", "Type"}); err != nil {
		panic(err)
	}

	// Loop through moves IDs;
	for id := 1; id <= 918; id++ {
		resp, err := http.Get("https://pokeapi.co/api/v2/move/" + strconv.Itoa(id))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var moveData MoveDetail
		json.Unmarshal(body, &moveData)

		// Remove hyphens and capitalize to make it more readable
		moveName := strings.Title(strings.ReplaceAll(moveData.Name, "-", " "))

		// Write to CSV
		if err := writer.Write([]string{strconv.Itoa(moveData.ID), moveName, moveData.Type.Name}); err != nil {
			panic(err)
		}

		// // Sleep to avoid API rate-limiting
		// time.Sleep(100 * time.Millisecond)
		fmt.Println("Processed move ID", id)
	}
	fmt.Println("CSV file generated successfully!")

}
