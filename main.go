package main

import (
	"fmt"
	"strconv"
	"os"
)

func main(){
	service := internal.NewPokemonService(internal.NewPokemonRepository())

	args := os.Args
id, _ := strconv.Atoi(args[2])
name := args[3]
type_ := args[4]

pokemon := internal.Pokemon{
	ID: id,
	Name: name,
	Type: pokemonType,
}

err := service.AddPokemon(pokemon)
if err != nil {
	fmt.Printf("Error adding Pokemon %v\n", err)
} else {
	fmt.Println("Pokemon added successfully")
}

err = internal.SavePokedexToFile("database.json", repo)
if err != nil{
	fmt.Printf("Error saving Pokedex %v\n", err)
	os.Exit(1)
}
}