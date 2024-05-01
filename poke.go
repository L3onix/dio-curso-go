package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Name     string    `json: "name"`
	Id       int       `json: "id"`
	Pokemons []Pokemon `json: "pokemon_entries"`
}

type Pokemon struct {
	EntryNumber int           `json: "entry_number"`
	Specie      PokemonSpecie `json: "pokemon_species`
}

type PokemonSpecie struct {
	Name string `json: "name"`
	URL  string `json: "url"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto")

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		pokemons := Response{}
		err = json.Unmarshal(data, &pokemons)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(string(data))
		fmt.Println(pokemons)

		for _, pokemon := range pokemons.Pokemons {
			fmt.Println(pokemon.Specie.Name)
		}
	}
	// fmt.Println(reflect.TypeOf(data))

}
