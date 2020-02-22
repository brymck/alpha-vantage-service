package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	pb "github.com/brymck/alpha-vantage-service/genproto"
)

func mustGetenv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

func stringToDate(s string) (*pb.Date, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 3 {
		return nil, fmt.Errorf("expected three parts: %s", s)
	}
	year, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return nil, err
	}
	month, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		return nil, err
	}
	day, err := strconv.ParseInt(parts[2], 10, 32)
	if err != nil {
		return nil, err
	}
	return &pb.Date{Year: int32(year), Month: int32(month), Day: int32(day)}, nil
}
