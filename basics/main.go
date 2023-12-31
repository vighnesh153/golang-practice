package main

import (
	"fmt"
)

func main() {
	// var pokemon1 string = "Pikachu"
	// pokemon2 := "Charizard"
	// pokemon3 := newPokemon()

	// fmt.Println(pokemon1)
	// fmt.Println(pokemon2)
	// fmt.Println(pokemon3)

	///--------------------------------------------

	// pokemons := []string{"Lucario", "Sceptile"}
	// pokemons2 := append(pokemons, "Infernape")
	// fmt.Println(pokemons)
	// fmt.Println(pokemons2)

	// for _, pokemon := range pokemons2 {
	// 	fmt.Println(pokemon)
	// }
	// for i, pokemon := range pokemons2 {
	// 	fmt.Println(i, pokemon)
	// }

	///--------------------------------------------

	// pokemons := PokemonNamesSlice{"Lucario", "Sceptile", "Infernape"}
	// pokemons.print()

	///--------------------------------------------

	// pokemons := PokemonNamesSlice{"Lucario", "Sceptile", "Infernape"}
	// firstPart, secondPart := pokemons.divide(1)
	// firstPart.print()
	// secondPart.print()

	///--------------------------------------------

	// pokemons := PokemonNamesSlice{"Lucario", "Sceptile", "Infernape"}
	// pokemons.saveToFile("pokemons.txt")

	// ///--------------------------------------------

	// pokemons, err := readFromFile("pokemons.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	// pokemons.print()

	///--------------------------------------------

	// pokemons, err := readFromFile("pokemons.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	// pokemons.shuffle()
	// pokemons.print()

	///--------------------------------------------

	// pikachu := Pokemon{
	// 	name: "Pikachu",
	// 	moves: []PokemonMoves{
	// 		{
	// 			pokemonType: "electric",
	// 			moves: []string{
	// 				"electro volt",
	// 				"volt tackle",
	// 			},
	// 		},
	// 		{
	// 			pokemonType: "normal",
	// 			moves: []string{
	// 				"tackle",
	// 				"quick attack",
	// 			},
	// 		},
	// 	},
	// }
	// pikachu.printName()
	// pikachu.updateName("Pikachu ⚡️")
	// pikachu.printName()

	// pikachu2 := &pikachu
	// pikachu2.printName()
	// pikachu2.updateName("Pikachu")
	// pikachu2.printName()
	// pikachu.printName()

	// pikachu.print()
	// pikachu2.print()
	// pikachu.moves[0].pokemonType = "thunder"
	// pikachu.print()
	// pikachu2.print()

	// moves1 := pikachu.moves
	// moves2 := moves1

	// fmt.Println(moves1)
	// fmt.Println(moves2)
	// moves1[0].pokemonType = "thunder"
	// fmt.Println(moves1)
	// fmt.Println(moves2)

	///--------------------------------------------

	// a := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// }

	// fmt.Printf("%v\n", a)

	// var b map[string]int = make(map[string]int)
	// b["a"] = 1

	// delete(a, "a")

	// fmt.Printf("%v\n", a)

	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("i: %v\n", i)
	// }

	///--------------------------------------------

	// eB := EnglishBot{}
	// sB := SpanishBot{}

	// eB.getGreeting()
	// sB.getGreeting()
	// printGreeting(&eB)
	// printGreeting(&sB)

	///--------------------------------------------

	// res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	// defer res.Body.Close()
	// body, _ := io.ReadAll(res.Body)
	// fmt.Println(string(body))

	a := make([]int, 0, 20)

	fmt.Println(len(a))
	fmt.Printf("a[0]=%v.\n", a[0])

}

type User struct {
	id string
}

type Bot interface {
	getGreeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func (*EnglishBot) getGreeting() string {
	return "Hi there"
}

func (*SpanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}
