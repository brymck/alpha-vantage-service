package alphavantage

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brymck/alpha-vantage-service/pkg/models"
)

func asFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func asInt32(s string) (int32, error) {
	n, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(n), nil
}

func asDate(s string) (*models.Date, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 3 {
		return nil, fmt.Errorf("expected three parts: %s", s)
	}
	year, err := asInt32(parts[0])
	if err != nil {
		return nil, err
	}
	month, err := asInt32(parts[1])
	if err != nil {
		return nil, err
	}
	day, err := asInt32(parts[2])
	if err != nil {
		return nil, err
	}
	return &models.Date{Year: year, Month: month, Day: day}, nil
}
