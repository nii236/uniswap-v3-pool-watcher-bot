package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type EthCallResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
}

// updatedMsg := fmt.Sprintf("Token 0: %s Token 1: %s", unclaimed_fees_token_0, unclaimed_fees_token_1)
func HandleStatusCmd() string {
	final_result := ""
	for k, v := range RegisteredPools {
		var jsonStr = fmt.Sprintf(`
			{
				"jsonrpc": "2.0",
				"id": 1,
				"method": "eth_call",
				"params": [
					{
						"from": "%s",
						"to": "%s",
						"data": "%s"
					},
					"latest"
				]
			}
		`, v.from, v.to, v.data)
		requestBody := bytes.NewBuffer([]byte(jsonStr))
		API_URL := "https://mainnet.infura.io/v3/099fc58e0de9451d80b18d7c74caa7c1"
		client := &http.Client{}
		req, err := http.NewRequest("POST", API_URL, requestBody)
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("origin", "https://app.uniswap.org/")

		resp, err := client.Do(req)
		// Handle Error
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
		//Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var output EthCallResponse
		json.Unmarshal(body, &output)

		//get token0 unclaimed fees
		token0_unclaimed_fees := CalcUnclaimedFees(0, output.Result[2:66])

		// //get token1 unclaimed fees
		token1_unclaimed_fees := CalcUnclaimedFees(1, output.Result[66:])

		token0_name := strings.Split(k, "/")[0]
		token1_name := strings.Split(k, "/")[1]
		final_result += fmt.Sprintf("%s: %s %s: %s", token0_name, token0_unclaimed_fees, token1_name, token1_unclaimed_fees)
	}
	return final_result
}