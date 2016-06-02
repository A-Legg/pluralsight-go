package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=googl")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	quote := new(Quoteresponse)
	xml.Unmarshal(body, &quote)

	fmt.Printf("%s: %.2f", quote.Name, quote.LastPrice)
}

type Quoteresponse struct {
	Status        string
	Name          string
	LastPrice     float32
	ChangePercent float32
	TimeStamp     string
	MSDate        float32
	MarketCap     int
	Volume        int
	ChangeYTD     float32
	High          float32
	Low           float32
	Open          float32
}
