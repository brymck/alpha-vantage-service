package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (s *server) callApi(fn string, params map[string]string, v interface{}) error {
	query := url.Values{}
	query.Set("function", fn)
	for k, v := range params {
		query.Set(k, v)
	}
	query.Set("apikey", s.apiKey)
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
