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

func HandleStatusCmd(geth_url string) (string) {
	final_result := ""
	for k, v := range RegisteredPools {
		token0_name := strings.Split(k, "/")[0]
		token1_name := strings.Split(k, "/")[1]
		errorOutputStr := fmt.Sprintf("%s/%s: Unable to fetch\n", token0_name, token1_name)
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
		API_URL := geth_url
		client := &http.Client{}
		req, err := http.NewRequest("POST", API_URL, requestBody)
		if err != nil {
			log.Printf("Error %v", err)
			final_result += errorOutputStr
			continue;
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("origin", "https://app.uniswap.org/")

		resp, err := client.Do(req)
		// Handle Error
		if err != nil {
			log.Printf("Error %v", err)
			final_result += errorOutputStr
			continue;
		}
		defer resp.Body.Close()
		//Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error %v", err)
			final_result += errorOutputStr
			continue;
		}
		var output EthCallResponse
		err = json.Unmarshal(body, &output)
		if err != nil {
			log.Printf("Error %v", err)
			final_result += errorOutputStr
			continue;
		}

		//get token0 unclaimed fees
		token0_unclaimed_fees, err := CalcUnclaimedFees(0, output.Result[2:66])
		if err != nil {
			log.Printf("Error %v", err)
			final_result += errorOutputStr
			continue;
		}

		// //get token1 unclaimed fees
		token1_unclaimed_fees, err := CalcUnclaimedFees(1, output.Result[66:])
		if err != nil {
			log.Printf("Error %v", err)
			final_result += errorOutputStr
			continue;
		}

		final_result += fmt.Sprintf("%s: %s %s: %s \n", token0_name, token0_unclaimed_fees, token1_name, token1_unclaimed_fees)
	}
	return final_result
}
