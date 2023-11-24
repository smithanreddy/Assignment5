package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const coinCapAPIURL = "https://api.coincap.io/v2/assets/"

type CryptoData struct {
	ID                string `json:"id"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
}

var cryptoPrices = map[string]CryptoData{}

func fetchCryptoData(cryptoSymbol string) error {
	resp, err := http.Get(coinCapAPIURL + cryptoSymbol)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var cryptoData struct {
		Data CryptoData `json:"data"`
	}
	err = json.Unmarshal(body, &cryptoData)
	if err != nil {
		return err
	}

	cryptoPrices[cryptoSymbol] = cryptoData.Data
	return nil
}

func getPriceHandler(w http.ResponseWriter, r *http.Request) {
	getData(w, "bitcoin")
	getData(w, "ethereum")
	getData(w, "tether")

}

func getData(w http.ResponseWriter, cryptoSymbol string) {

	err := fetchCryptoData(cryptoSymbol)
	if err != nil {
		http.Error(w, "Error fetching cryptocurrency data", http.StatusInternalServerError)
		return
	}

	cryptoData, found := cryptoPrices[cryptoSymbol]

	if !found {
		http.Error(w, "Cryptocurrency not found", http.StatusNotFound)
		return
	}

	// Convert the price to CAD
	priceUsdFloat, _ := strconv.ParseFloat(cryptoData.PriceUsd, 64)
	priceUsdFloat *= 1.33

	response := map[string]interface{}{
		"crypto":   cryptoData.Name,
		"priceCad": fmt.Sprintf("%.2f", priceUsdFloat),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/price", getPriceHandler)
	fmt.Println("Server is listening on port 8088")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
