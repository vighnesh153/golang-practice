package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pokemon struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	HitPoints float64  `json:"hit_points"`
	Types     []string `json:"types"`
}

var pokemons = []pokemon{
	{
		ID:        "pikachu",
		Name:      "Pikachu",
		HitPoints: 60.0,
		Types:     []string{"Electric"},
	},
	{
		ID:        "infernape",
		Name:      "Infernape",
		HitPoints: 80.0,
		Types:     []string{"Fire", "Fighting"},
	},
	{
		ID:        "charizard",
		Name:      "Charizard",
		HitPoints: 95.0,
		Types:     []string{"Fire", "Flying"},
	},
	{
		ID:        "sceptile",
		Name:      "Sceptile",
		HitPoints: 90.0,
		Types:     []string{"Grass"},
	},
	{
		ID:        "lucario",
		Name:      "Lucario",
		HitPoints: 95.0,
		Types:     []string{"Fighting", "Steel"},
	},
	{
		ID:        "greninja",
		Name:      "Greninja",
		HitPoints: 100.0,
		Types:     []string{"Water", "Dark"},
	},
}

func main() {
	router := gin.Default()

	router.GET("/pokemons", getPokemons)
	router.GET("/pokemons/:id", getPokemonById)
	router.POST("/pokemons", postPokemons)

	router.Run("localhost:4242")
}

func getPokemons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pokemons)
}

func postPokemons(c *gin.Context) {
	var newPokemon pokemon

	if err := c.BindJSON(&newPokemon); err != nil {
		return
	}

	pokemons = append(pokemons, newPokemon)
	c.IndentedJSON(http.StatusCreated, newPokemon)
}

func getPokemonById(c *gin.Context) {
	id := c.Param("id")
	for _, pokemon := range pokemons {
		if pokemon.ID == id {
			c.IndentedJSON(http.StatusOK, pokemon)
			return
		}
	}
}
