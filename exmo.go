package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Ticker struct {
	Buy_price   float64 `json:",string"`
	Last_trade  float64 `json:",string"`
	Sell_price  float64 `json:",string"`
	High        float64 `json:",string"`
	Low         float64 `json:",string"`
	Avg         float64 `json:",string"`
	Vol         float64 `json:",string"`
	Vol_curr    float64 `json:",string"`
	Updated     int64   `json:",string"`
}

type Exmo struct {
	Name string

	BaseURL string

	MarketProvider

	Ticker map[string]Ticker
}

func (e *Exmo)getCurrencies(c *MarketClient) *[]Currency {//*map[string]Currency{

	//var currencies = make(map[string]Currency)
	var currencies []Currency


	resultBody := c.performRequest("https://api.exmo.com/v1/","ticker")
	errUnmarsh := json.Unmarshal(resultBody, &e.Ticker)
	if errUnmarsh != nil {
		fmt.Errorf("JSON Bittrex unmarshal error: %v", errUnmarsh)
	}

	for k,v := range e.Ticker {
		k = strings.Replace(k,"_", "-", -1)

		currencies = append(currencies, Currency{
			Exchange: e.Name,
			MarketName: k,
			Bid: v.Buy_price,
			Ask: v.Sell_price,
		})


	}

	return &currencies
}