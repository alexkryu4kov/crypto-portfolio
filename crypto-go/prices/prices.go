package prices

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type PriceResponse struct {
	Bitcoin  Currency `json:"bitcoin"`
	Ethereum Currency `json:"ethereum"`
}

type Currency struct {
	USD float64 `json:"usd"`
}

func FetchCryptoPrices() (PriceResponse, error) {
	client := resty.New()
	resp, err := client.R().SetQueryParams(map[string]string{
		"ids":           "bitcoin,ethereum",
		"vs_currencies": "usd",
	}).
		Get("https://api.coingecko.com/api/v3/simple/price")

	if err != nil {
		return PriceResponse{}, err
	}
	var priceData PriceResponse
	err = json.Unmarshal(resp.Body(), &priceData)
	if err != nil {
		return PriceResponse{}, err
	}

	fmt.Printf("Bitcoin Price: %f\n", priceData.Bitcoin.USD)
	fmt.Printf("Ethereum Price: %f\n", priceData.Ethereum.USD)

	return priceData, nil
}
