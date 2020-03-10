package main

import (
    "encoding/json"
    "net/http"
    "fmt"
)

type APIPriceResponse struct {
    Price float64 `json:"price,string"`
}

type APIChangeResponse struct {
    PriceChange float64 `json:"priceChange,string"`
    PriceChangePercent float64 `json:"priceChangePercent,string"`
    HighPrice float64 `json:"highPrice,string"`
    LowPrice float64 `json:"lowPrice,string"`
}

/* absolute value, 24hr change, err */
func queryHandshake() (*APIPriceResponse, *APIChangeResponse, error) {
    var price APIPriceResponse
    var change APIChangeResponse
    resp, err := http.Get("https://www.namebase.io/api/v0/ticker/price?symbol=HNSBTC")
    if err != nil {
        return &price, &change, err
    }
    decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&price)
    if err != nil {
        return &price, &change, err
    }
    resp.Body.Close()

    resp, err = http.Get("https://www.namebase.io/api/v0/ticker/day?symbol=HNSBTC")
    if err != nil {
        return &price, &change, err
    }
    decoder = json.NewDecoder(resp.Body)
    err = decoder.Decode(&change)
    if err != nil {
        return &price, &change, err
    }
    resp.Body.Close()
    return &price, &change, err
}

func main() {
    price, change, e := queryHandshake()
    if e != nil {
        fmt.Println(e)
    }
    fmt.Println(price)
    fmt.Println(change)
}
