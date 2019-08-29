package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ceeyahya/crypto-bot/types"
)

func floatToString(n float64) string {
	return strconv.FormatFloat(n, 'f', 6, 64)
}

// GetCurrency handles the sent message and makes API calls
func GetCurrency(symbol string) string {
	req := "http://api.coinlayer.com/api/live?access_key=26bf59d11a750f73ee6f72c88f0ae043&symbols="
	resp, err := http.Get(req + symbol)
	if err != nil {
		fmt.Printf("Request Failed: %s", err)
	}
	rd, err := ioutil.ReadAll(resp.Body)

	var ro types.Currency
	json.Unmarshal(rd, &ro)

	p := floatToString(ro.Rates[symbol])

	return p
}
