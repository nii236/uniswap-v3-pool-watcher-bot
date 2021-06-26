package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Params: Name of token whose price needs to be fetched. Ex: uniswap, ethereum
func FetchTokenPrice(slug string) (float64, error) {
	API_URL := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", slug)
	client := &http.Client{}
	req, err := http.NewRequest("GET", API_URL, nil)
	if err != nil {
		log.Printf("Error %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer resp.Body.Close()
	output := make(map[string]map[string]float32)
	err = json.Unmarshal(body, &output)
	if err != nil {
		log.Printf("Error %v", err)
		return 0, nil
	}
	priceInUSD := output[slug]["usd"]
	return float64(priceInUSD), nil
}
