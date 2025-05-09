package util

import (
	"catching-pokemons/models"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserPokemonSuccess(t *testing.T) {
	c := require.New(t)

	// Respuesta del resultado de la API
	body, err := ioutil.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse
	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)

	parserPokemon, err := ParsePokemon(response)
	c.NoError(err)

	// Respuesta esperada
	body, err = ioutil.ReadFile("samples/api_response.json")
	c.NoError(err)

	var expected models.Pokemon
	err = json.Unmarshal([]byte(body), &expected)
	c.NoError(err)

	c.Equal(expected, parserPokemon, "el pokemon analizado no es igual al esperado")
}

func TestParserPokemonTypeNotFound(t *testing.T) {
	c := require.New(t)

	// Lectura del archivo de respuesta de la API
	body, err := ioutil.ReadFile("samples/pokeapi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse
	err = json.Unmarshal([]byte(body), &response)
	c.NoError(err)

	response.PokemonType = []models.PokemonType{}

	_, err = ParsePokemon(response)
	c.NotNil(err)
	c.EqualError(ErrNotFoundPokemonType, err.Error())
}
