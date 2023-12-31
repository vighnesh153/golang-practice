package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type PokemonNamesSlice []string

type PokemonMoves struct {
	pokemonType string
	moves       []string
}

type Pokemon struct {
	name  string
	moves []PokemonMoves
}

func (p Pokemon) print() {
	fmt.Printf("%+v\n", p)
}

func (p Pokemon) printName() {
	fmt.Println(p.name)
}

func (p *Pokemon) updateName(newName string) {
	p.name = newName
}

func newPokemons() PokemonNamesSlice {
	return []string{
		"Pikachu",
		"Charizard",
		"Sceptile",
		"Infernape",
		"Greninja",
		"Lucario",
	}
}

func (pokemons PokemonNamesSlice) getSlice(start int, end int) PokemonNamesSlice {
	return pokemons[start:end]
}

func (pokemons PokemonNamesSlice) divide(size int) (PokemonNamesSlice, PokemonNamesSlice) {
	return pokemons[:size], pokemons[size:]
}

func (slice PokemonNamesSlice) print() {
	for i, pokemon := range slice {
		fmt.Println(i, pokemon)
	}
}

func (slice PokemonNamesSlice) toString() string {
	return strings.Join(slice, ", ")
}

func (slice PokemonNamesSlice) saveToFile(filename string) error {
	byteRepr := []byte(slice.toString())

	return os.WriteFile(filename, byteRepr, 0666)
}

func readFromFile(filename string) (PokemonNamesSlice, error) {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	stringContent := string(byteSlice)
	return strings.Split(stringContent, ", "), nil
}

func (slice PokemonNamesSlice) shuffle() {
	// source := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(source)

	// for i := range slice {
	// 	newPosition := r.Intn(len(slice))
	// 	slice[i], slice[newPosition] = slice[newPosition], slice[i]
	// }

	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
}
