package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type StockData struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Time   string  `json:"time"`
}

func FetchStockPrice(apiKey, symbol string) (StockData, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", symbol, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return StockData{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return StockData{}, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return StockData{}, err
	}

	// Extract stock price from API response
	data := result["Global Quote"].(map[string]interface{})
	stock := StockData{
		Symbol: data["01. symbol"].(string),
		Price:  parseFloat(data["05. price"].(string)),
		Time:   data["07. latest trading day"].(string),
	}
	return stock, nil
}

func parseFloat(value string) float64 {
	parsed, _ := strconv.ParseFloat(value, 64)
	return parsed
}
