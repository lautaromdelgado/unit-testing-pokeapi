package controller

import (
	"catching-pokemons/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ErrPokemonNotFound = fmt.Errorf("pokemon not found")
	ErrPokeApiFailure  = fmt.Errorf("unexpected response from pokeapi")
)

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	apiPokemon, err := GetPokemonFromPokeApi(id)
	if errors.Is(err, ErrPokemonNotFound) {
		respondwithJSON(w, http.StatusNotFound, fmt.Sprintf("pokemon with id %s not found", id))
	}
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, fmt.Sprintf("error while calling pokeapi: %s", err.Error()))
	}

	// parsedPokemon, err := util.ParsePokemon(apiPokemon)
	// if err != nil {
	// 	respondwithJSON(w, http.StatusInternalServerError, fmt.Sprintf("error found: %s", err.Error()))
	// }

	respondwithJSON(w, http.StatusOK, apiPokemon)
}

func GetPokemonFromPokeApi(id string) (models.PokeApiPokemonResponse, error) {
	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	response, err := http.Get(request)
	if err != nil {
		return models.PokeApiPokemonResponse{}, err
	}

	// Error Status Not Found
	if response.StatusCode == http.StatusNotFound {
		return models.PokeApiPokemonResponse{}, ErrPokemonNotFound
	}

	// Error Status OK
	// Si el estado no es 200, entonces no se puede parsear el body
	if response.StatusCode != http.StatusOK {
		return models.PokeApiPokemonResponse{}, ErrPokeApiFailure
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return models.PokeApiPokemonResponse{}, err
	}

	var apiPokemon models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &apiPokemon)
	if err != nil {
		log.Fatal(err)
	}

	return apiPokemon, nil
}
