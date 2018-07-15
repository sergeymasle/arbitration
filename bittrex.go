package main

import (
	"encoding/json"
	"fmt"
)

type BittrexMarketSummaries struct {
	Success bool `json:"bool"`
	Message int  `json:"string,omitempty"`
	Result []struct {
		MarketName     string                            //`json:",string"`//"MarketName": "BTC-2GIVE",
		High           float64 `json:"High"`             //"High": 0.00000084,
		Low            float64 `json:"Low"`              //"Low": 0.00000081,
		Volume         float64 `json:"Volume"`           //"Volume": 2501632.91033042,
		Last           float64 `json:"Last"`             //"Last": 0.00000083,
		BaseVolume     float64 `json:"BaseVolume"`       //"BaseVolume": 2.06221629,
		TimeStamp      string  `json:"string,omitempty"` //"TimeStamp": "2018-07-09T13:32:07.887",
		Bid            float64 `json:"Bid"`              //"Bid": 0.00000082,
		Ask            float64 `json:"Ask"`              //"Ask": 0.00000083,
		OpenBuyOrders  int     `json:"OpenBuyOrders"`    //"OpenBuyOrders": 68,
		OpenSellOrders int     `json:"OpenSellOrders"`   //"OpenSellOrders": 778,
		PrevDay        float64 `json:"PrevDay"`          //"PrevDay": 0.00000083,
		Created        string  `json:"string,omitempty"` //"Created": "2016-05-16T06:44:15.287"
	}
}

type Bittrex struct {
	Name string

	BaseURL string

	MarketProvider

	BittrexMarketSummaries
}

func (b *Bittrex)getCurrencies(c *MarketClient) *[]Currency {

	var currencies []Currency

	resultBody := c.performRequest("https://bittrex.com/api/v1.1/public/","getmarketsummaries")

	errUnmarsh := json.Unmarshal(resultBody, &b.BittrexMarketSummaries)
	if errUnmarsh != nil {
		fmt.Errorf("JSON Bittrex unmarshal error: %v", errUnmarsh)
	}

	for _,v := range b.BittrexMarketSummaries.Result {
		 c := Currency {
		 	 Exchange: b.Name,
			 MarketName: v.MarketName,
			 Bid: v.Bid,
			 Ask: v.Ask,
		 }
		currencies = append(currencies, c)
	}

	return &currencies
}