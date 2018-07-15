package main

import (
	"strings"
	"net/http"
	"io/ioutil"
	"fmt"
	"time"
	"encoding/json"
)

type Arbitration struct {
	From Currency
	To Currency
	Profit float64
}

type Currency struct {
	Exchange string `json:"Exchange"`
	MarketName string `json:"MarketName"`
	Bid float64 `json:"Bid"`
	Ask float64 `json:"Ask"`
}

type MarketProvider interface {
	getCurrencies(c *MarketClient) []Currency
}

// Compare name market BTC/ETH = ETH/BTC
func equalsPairs(pair1 string, pair2 string) bool  {
	p1 := strings.Split(pair1,"-")
	p2 := strings.Split(pair2,"-")

	var result bool

	if ( (p1[0]==p2[0]) || (p1[0]==p2[1]) ) && ( (p1[1]==p2[0]) || (p1[1]==p2[1]) ) {
		result = true
	} else {
		result = false
	}

	return result
}




func comparePairs(pair1, pair2 *[]Currency) []Arbitration  {

	var arbit []Arbitration

	for _,v1 := range *pair1 {
		for _,v2 := range *pair2 {
			if equalsPairs(v1.MarketName,v2.MarketName) {

				if v1.Ask < v2.Bid {

					r := v2.Bid - v1.Ask
					p := r*100/v2.Bid

					//TODO
					if (p > 0.0) && (p < 10 ){
						arbit = append(arbit, Arbitration{From:v1, To: v2, Profit: p})
					}

				}

				if v2.Ask < v1.Bid {
					r := v1.Bid - v2.Ask
					p := r*100/v1.Bid

					//TODO
					if (p > 0.0) && (p < 10 ){
						arbit = append(arbit, Arbitration{From:v2, To: v1, Profit: p})
					}
				}

			}
		}
	}

	return arbit
}



func index(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		fmt.Println(err)
	}
	w.Write(body)

}

func get(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Write(*store)
}


var store *[]byte

func main()  {

	client := MarketClient{httpClient: &http.Client{},}

	bittrex := Bittrex{
		Name: "Bittrex",
		BaseURL: "https://bittrex.com/api/v1.1/public/",
	}
	exmo :=	&Exmo{
		Name: "Exmo",
		BaseURL: "https://api.exmo.com/v1/",
	}




	go func(c *MarketClient, storage *[]byte) {
		for {
			resultCompare := comparePairs(
				bittrex.getCurrencies(c),
				exmo.getCurrencies(c),
			)

			b, err := json.Marshal(resultCompare)
			if err != nil {
				fmt.Println(err)
				return
			}

			store = &b

			time.Sleep(5 * time.Second)
		}

	}(&client, store)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/get", get)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}



}