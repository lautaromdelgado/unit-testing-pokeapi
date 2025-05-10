package controller

import (
	"catching-pokemons/models"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

// func TestGetPokemonFromPokeApiSuccess(t *testing.T) {
// 	c := require.New(t)

// 	pokemon, err := GetPokemonFromPokeApi("bulbasaur")
// 	c.NoError(err, "se esperaba que no hubiera un error")

// 	body, err := ioutil.ReadFile("samples/poke_api_read.json")
// 	c.NoError(err, "se esperaba que no hubiera un error al leer el archivo: %s", err.Error())

// 	var expected models.PokeApiPokemonResponse
// 	err = json.Unmarshal([]byte(body), &expected)
// 	c.NoError(err, "se esperaba que no hubiera un error al unmarshalar el archivo: %s", err.Error())

// 	c.Equal(expected, pokemon, "se esperaba que el pokemon fuera igual al esperado")
// }

func TestGetPokemonFromPokeApiSuccessWithMocks(t *testing.T) {
	c := require.New(t)

	httpmock.Activate()                 // Activar el mock
	defer httpmock.DeactivateAndReset() // Cerrar el mock al final de la prueba

	id := "bulbasaur"

	body, err := ioutil.ReadFile("samples/poke_api_response.json")
	if body == nil {
		c.Fail("se esperaba que el body no fuera nulo")
	}
	c.NoError(err, "se esperaba que no hubiera un error al leer el archivo: %s", err.Error())

	httpmock.RegisterResponder("GET", "https://pokeapi.co/api/v2/pokemon/"+id,
		httpmock.NewStringResponder(200, string(body)))

	pokemon, err := GetPokemonFromPokeApi(id)
	c.NoError(err, "se esperaba que no hubiera un error")

	var expected models.PokeApiPokemonResponse
	err = json.Unmarshal([]byte(body), &expected)
	c.NoError(err, "se esperaba que no hubiera un error al unmarshalar el archivo: %s", err.Error())

	c.Equal(expected, pokemon, "se esperaba que el pokemon fuera igual al esperado")
}
