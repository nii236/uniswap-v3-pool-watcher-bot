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

// Calls the Uniswap V3 API and returns 130 hex chars length output string
// This string is later processed to get the individual tokens unclaimed fees percentage
func UniswapAPICall(from string, to string, data string, geth_url string) (string, error) {
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
	`, from, to, data,
	)
	requestBody := bytes.NewBuffer([]byte(jsonStr))
	API_URL := geth_url
	client := &http.Client{}
	req, err := http.NewRequest("POST", API_URL, requestBody)
	if err != nil {
		log.Printf("Error %v", err)
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("origin", "https://app.uniswap.org/")

	resp, err := client.Do(req)
	// Handle Error
	if err != nil {
		log.Printf("Error %v", err)
		return "", err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var output EthCallResponse
	err = json.Unmarshal(body, &output)
	if err != nil {
		return "", err
	}
	return output.Result, nil
}

// Returns the final message the bot sends when it receives /status command
func HandleStatusCmd(geth_url string) (string, error) {
	final_result := ""

	for k, v := range RegisteredPools {
		token0_name := strings.Split(k, "/")[0]
		token1_name := strings.Split(k, "/")[1]
		uniswapResp, err := UniswapAPICall(v.from, v.to, v.data, geth_url)
		if err != nil {
			return "", err
		}
		//get token0 unclaimed fees
		token0_unclaimed_fees, err := CalcUnclaimedFees(0, uniswapResp[2:66])
		if err != nil {
			return "", err
		}

		//get token1 unclaimed fees
		token1_unclaimed_fees, err := CalcUnclaimedFees(1, uniswapResp[66:])
		if err != nil {

			return "", err
		}

		// construct & append final result to send to telegram bot
		final_result += fmt.Sprintf(
			"%s: %s %s: %s\n",
			token0_name,
			token0_unclaimed_fees,
			token1_name,
			token1_unclaimed_fees,
		)
	}

	return final_result, nil
}
