package alphavantage

import "github.com/brymck/alpha-vantage-service/pkg/models"

// A Quote contains the latest price and volume information for a security.
type Quote struct {
	// Open price
	Open float64
	// High price
	High float64
	// Low price
	Low float64
	// Current price
	Price float64
	// Daily volume
	Volume float64
	// Latest trading day
	LatestTradingDay *models.Date
	// Previous close
	PreviousClose float64
}

type globalQuoteResponse struct {
	GlobalQuote struct {
		Symbol           string `json:"01. symbol"`
		Open             string `json:"02. open"`
		High             string `json:"03. high"`
		Low              string `json:"04. low"`
		Price            string `json:"05. price"`
		Volume           string `json:"06. volume"`
		LatestTradingDay string `json:"07. latest trading day"`
		PreviousClose    string `json:"08. previous close"`
		Change           string `json:"09. change"`
		ChangePercent    string `json:"10. change percent"`
	} `json:"Global Quote"`
}

// GetQuote retrieves the latest price and volume information for a security of your choice.
//
// See here for more information:
//
// https://www.alphavantage.co/documentation/
func (api *AlphaVantageApi) GetQuote(symbol string) (*Quote, error) {
	var resp globalQuoteResponse
	err := api.call("GLOBAL_QUOTE", map[string]string{"symbol": symbol}, &resp)
	if err != nil {
		return nil, err
	}
	gq := resp.GlobalQuote
	open, err := asFloat64(gq.Open)
	if err != nil {
		return nil, err
	}
	high, err := asFloat64(gq.High)
	if err != nil {
		return nil, err
	}
	low, err := asFloat64(gq.Low)
	if err != nil {
		return nil, err
	}
	price, err := asFloat64(gq.Price)
	if err != nil {
		return nil, err
	}
	volume, err := asFloat64(gq.Volume)
	if err != nil {
		return nil, err
	}
	latestTradingDay, err := asDate(gq.LatestTradingDay)
	if err != nil {
		return nil, err
	}
	previousClose, err := asFloat64(gq.PreviousClose)
	if err != nil {
		return nil, err
	}

	return &Quote{
		Open:             open,
		High:             high,
		Low:              low,
		Price:            price,
		Volume:           volume,
		LatestTradingDay: latestTradingDay,
		PreviousClose:    previousClose,
	}, nil
}
