package main

import (
	"crypto-go/prices"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		priceData, err := prices.FetchCryptoPrices()
		if err != nil {
			fmt.Println("Error fetching prices:", err)
			continue
		}
		SavePricesToFile(priceData)

		fmt.Println("Prices saved:", priceData)
	}
}

func SavePricesToFile(priceData prices.PriceResponse) {

	data := map[string]interface{}{
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		"prices":    priceData,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	file, err := os.OpenFile("data/crypto_prices.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}

	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
	}

	_, err = file.WriteString("\n")
	if err != nil {
		fmt.Println("Error writing newline to file")
	}
}
