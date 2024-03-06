package gateways

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/HenCor2019/task-go/api/models"
)

func Fetch(quantityEstimations int, wg *sync.WaitGroup) ([]models.Shipengine, error) {
	estimations := []models.Shipengine{}
	for i := 0; i <= quantityEstimations; i++ {
		wg.Add(1)
		go func() {
			init := time.Now()
			data := []byte(`{
      "carrier_ids": [ "se-2221957", "se-2221960", "se-2221962" ],
  "from_country_code": "US",
  "from_postal_code": "78756",
  "to_country_code": "US",
  "to_postal_code": "95128",
  "to_city_locality": "San Jose",
  "to_state_province": "CA",
  "weight": {
    "value": 16.0,
    "unit": "ounce"
  },
  "dimensions": {
    "unit": "inch",
    "length": 5.0,
    "width": 5.0,
    "height": 5.0
  },
  "confirmation": "none",
  "address_residential_indicator": "no",
  "ship_date": "2022-10-28T00:00:00.000Z"}`)

			req, _ := http.NewRequest("POST", os.Getenv("SHIPENGINE_BASE_URL"), bytes.NewBuffer(data))
			req.Header.Set("API-Key", os.Getenv("SHIPENGINE_API_KEY"))
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Println("Cannot fetch the request with ")
				wg.Done()
				return
			}

			if resp.Status != "200 OK" {
				wg.Done()
				log.Println("Cannot fetch the request with")
				return
			}

			defer resp.Body.Close()
			bodyBytes, _ := io.ReadAll(resp.Body)
			shipengineStruct := []models.Shipengine{}
			err = json.Unmarshal(bodyBytes, &shipengineStruct)
			if err != nil {
				wg.Done()
				log.Println("Cannot fetch the request with ")
				return
			}

			estimations = append(estimations, shipengineStruct...)
			end := time.Since(init).Milliseconds()
			log.Println("Goroutine finished in: ", end, " ms")
			wg.Done()
		}()
	}
	wg.Wait()
	return estimations, nil
}

func (g *Gateway) FetchRatesEstimation(quantityEstimations int) ([]models.Shipengine, error) {
	wg := &sync.WaitGroup{}
	carriersEstimations, err := Fetch(quantityEstimations, wg)

	return carriersEstimations, err
}
