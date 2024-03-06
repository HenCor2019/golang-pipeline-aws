package PokemonsGateways

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/HenCor2019/task-go/api/models"
	"github.com/gofiber/fiber/v2"
)

func Fetch(ids []int, wg *sync.WaitGroup) ([]models.Pokemon, error) {
	pokemons := []models.Pokemon{}
	for _, pokemonId := range ids {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			init := time.Now()
			resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(id))
			if err != nil {
				log.Println("Cannot fetch the request with id: ", id)
				return
			}

			if resp.Status != "200 OK" {
				log.Println("Cannot fetch the request with id: ", id)
				return
			}

			defer resp.Body.Close()
			bodyBytes, _ := io.ReadAll(resp.Body)
			var pokemonStruct models.Pokemon
			err = json.Unmarshal(bodyBytes, &pokemonStruct)
			if err != nil {
				log.Println("Cannot fetch the request with id: ", id)
				return
			}

			pokemons = append(pokemons, pokemonStruct)
			end := time.Since(init).Milliseconds()
			log.Println("Goroutine ", id, " finished in: ", end, " ms")
		}(pokemonId)
	}
	wg.Wait()
	hasAllPokemonFetched := len(pokemons) == len(ids)
	if !hasAllPokemonFetched {
		return pokemons, fiber.NewError(fiber.StatusBadRequest, "Some ids were invalid")
	}

	return pokemons, nil
}

func (g *Gateway) FetchPokemonsByIds(pokemonsIds []int) ([]models.Pokemon, error) {
	wg := &sync.WaitGroup{}
	pokemons, err := Fetch(pokemonsIds, wg)

	return pokemons, err
}
