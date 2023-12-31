package main

import (
	"os"
	"testing"
)

func TestNewPokemons(t *testing.T) {
	pokemons := newPokemons()

	if len(pokemons) != 6 {
		t.Fatalf("Expected the length of pokemons to be 6, but got %v", len(pokemons))
	}

	if pokemons[0] != "Pikachu" {
		t.Fatalf("Expected first pokemon to be Pikachu, but got %v", pokemons[0])
	}

	if pokemons[len(pokemons)-1] != "Lucario" {
		t.Fatalf("Expected first pokemon to be Lucario, but got %v", pokemons[0])
	}
}

func TestPokemonsSaveToFileAndReadFromFile(t *testing.T) {
	filename := "_pokemon-testing"
	os.Remove(filename)

	pokemons := newPokemons()

	pokemons.saveToFile(filename)

	loadedPokemons, _ := readFromFile(filename)

	if len(loadedPokemons) != len(pokemons)+1 {
		t.Fatalf("Expected the length of pokemons to be %v, but got %v", len(pokemons), len(loadedPokemons))
	}

	// cleanup
	os.Remove(filename)
}
