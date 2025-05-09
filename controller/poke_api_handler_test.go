package controller

import (
	"catching-pokemons/models"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPokemonFromPokeApiSuccess(t *testing.T) {
	c := require.New(t)

	pokemon, err := GetPokemonFromPokeApi("bulbasaur")
	c.NoError(err, "se esperaba que no hubiera un error")

	body, err := ioutil.ReadFile("samples/poke_api_read.json")
	c.NoError(err, "se esperaba que no hubiera un error al leer el archivo: %s", err.Error())

	var expected models.PokeApiPokemonResponse
	err = json.Unmarshal([]byte(body), &expected)
	c.NoError(err, "se esperaba que no hubiera un error al unmarshalar el archivo: %s", err.Error())

	c.Equal(expected, pokemon, "se esperaba que el pokemon fuera igual al esperado")
}
