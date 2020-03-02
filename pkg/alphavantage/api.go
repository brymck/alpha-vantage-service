package alphavantage

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// AlphaVantageApi wraps calls to the Alpha Vantage API.
type AlphaVantageApi struct {
	apiKey string
}

// NewApi accepts creates an AlphaVantageApi instance that will authenticate with the provided API key.
//
// You can obtain an API key here:
//
// https://www.alphavantage.co/support/#api-key
func NewApi(apiKey string) *AlphaVantageApi {
	return &AlphaVantageApi{apiKey: apiKey}
}

func (api *AlphaVantageApi) call(fn string, params map[string]string, v interface{}) error {
	query := url.Values{}
	query.Set("function", fn)
	for k, v := range params {
		query.Set(k, v)
	}
	query.Set("apikey", api.apiKey)
	u := &url.URL{
		Scheme:   "https",
		Host:     "www.alphavantage.co",
		Path:     "query",
		RawQuery: query.Encode(),
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&v)
	return err
}
