package pokeApi

import (
	"net/http"

	httpClient "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/http-client"
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/domain"
)

type PokeApiPokemonTypesRepository struct{}

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

func (pokeApiPokemonTypesRepository PokeApiPokemonTypesRepository) FindByPokemonName(pokemonName domain.PokemonName) (domain.PokemonTypes, error) {
	urlPath := pokeApiUrl + pokemonName.GetValue()
	response, statusCode, errorOnResponse := httpClient.Get(urlPath)

	if statusCode == http.StatusServiceUnavailable {
		serviceUnavailableException := domain.CreateRepositoryUnavailableException()
		return domain.PokemonTypes{}, serviceUnavailableException.GetError()
	}

	if errorOnResponse != nil {
		return domain.PokemonTypes{}, errorOnResponse
	}

	if statusCode == http.StatusNotFound {
		pokemonNotFoundException := domain.CreatePokemonNotFoundException(pokemonName)
		return domain.PokemonTypes{}, pokemonNotFoundException.GetError()
	}

	return mapResponseToPokemonTypes(pokemonName, response)
}
