package alphavantage

import (
	"sort"

	"github.com/brymck/alpha-vantage-service/pkg/models"
)

// A TimeSeriesDailyItem contains the date, price and volume data for a single date.
type TimeSeriesDailyItem struct {
	Date   *models.Date
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

type timeSeriesDailyResponse struct {
	TimeSeriesDaily map[string]struct {
		Open   string `json:"1. open"`
		High   string `json:"2. high"`
		Low    string `json:"3. low"`
		Close  string `json:"4. close"`
		Volume string `json:"5. volume"`
	} `json:"Time Series (Daily)"`
}

// GetTimeSeriesDaily returns the daily time series (date, open, high, low, close, volume) of the global equity
// specified, covering 20+ years of historical data.
//
// Refer to the Alpha Vantage API documentation for more information:
//
// https://www.alphavantage.co/documentation/#daily
func (api *AlphaVantageApi) GetTimeSeriesDaily(symbol string, full bool) ([]*TimeSeriesDailyItem, error) {
	var resp timeSeriesDailyResponse
	outputsize := "compact"
	if full {
		outputsize = "full"
	}
	err := api.call("TIME_SERIES_DAILY", map[string]string{"symbol": symbol, "outputsize": outputsize}, &resp)
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
		high, err := asFloat64(v.High)
		if err != nil {
			return nil, err
		}
		low, err := asFloat64(v.Low)
		if err != nil {
			return nil, err
		}
		closePrice, err := asFloat64(v.Close)
		if err != nil {
			return nil, err
		}
		volume, err := asFloat64(v.Volume)
		if err != nil {
			return nil, err
		}
		ts[i] = &TimeSeriesDailyItem{
			Date:   date,
			Open:   open,
			High:   high,
			Low:    low,
			Close:  closePrice,
			Volume: volume,
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
