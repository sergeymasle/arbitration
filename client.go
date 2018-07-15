package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

type MarketClient struct {
	httpClient *http.Client
}

func (mc *MarketClient) performRequest(URL string, method string) []byte {

	req, reqErr := http.NewRequest("GET", URL+method, nil)
	if reqErr != nil {
		fmt.Errorf("Request error: %v", reqErr)
	}

	resp, respErr := mc.httpClient.Do(req) //client.c.Do(req)
	if respErr != nil {
		fmt.Errorf("Response error: %v", respErr)
	}

	body, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		fmt.Errorf("Read body error: %v", bodyErr)
	}

	defer resp.Body.Close()

	return body
}