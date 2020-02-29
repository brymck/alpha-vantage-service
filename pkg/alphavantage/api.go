package alphavantage

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Api struct {
	apiKey string
}

func NewApi(apiKey string) *Api {
	return &Api{apiKey: apiKey}
}

func (api *Api) call(fn string, params map[string]string, v interface{}) error {
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
