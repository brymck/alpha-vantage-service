package alphavantage

import (
	"sort"

	"github.com/brymck/alpha-vantage-service/pkg/models"
)

type TimeSeriesDailyItem struct {
	Date *models.Date
	Open float64 `json:"1. open"`
}

type TimeSeriesDailyResponse struct {
	TimeSeriesDaily map[string]struct {
		Open string `json:"1. open"`
	} `json:"Time Series (Daily)"`
}

func (api *AlphaVantageApi) GetTimeSeriesDaily(symbol string) ([]*TimeSeriesDailyItem, error) {
	var resp TimeSeriesDailyResponse
	err := api.call("TIME_SERIES_DAILY", map[string]string{"symbol": symbol}, &resp)
	if err != nil {
		return nil, err
	}

	tsd := resp.TimeSeriesDaily
	ts := make([]*TimeSeriesDailyItem, len(tsd))
	i := 0
	for k, v := range tsd {
		date, err := asDate(k)
		if err != nil {
			return nil, err
		}
		open, err := asFloat64(v.Open)
		if err != nil {
			return nil, err
		}
		ts[i] = &TimeSeriesDailyItem{
			Date: date,
			Open: open,
		}
		i++
	}
	sort.Slice(ts, func(i, j int) bool {
		x := ts[i].Date
		y := ts[j].Date
		return x.Year < y.Year || (x.Year == y.Year && (x.Month < y.Month || (x.Month == y.Month && x.Day < y.Day)))
	})

	return ts, nil
}
