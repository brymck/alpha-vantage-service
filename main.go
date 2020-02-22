package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/brymck/alpha-vantage-service/genproto"
)

const (
	grpcPort = 50051
)

var (
	log = logrus.New()
)

type TimeSeriesDailyResponse struct {
	TimeSeriesDaily map[string]struct {
		Open string `json:"1. open"`
	} `json:"Time Series (Daily)"`
}

type server struct {
	apiKey string
}

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

func (s *server) GetTimeSeries(_ context.Context, in *pb.GetTimeSeriesRequest) (*pb.GetTimeSeriesResponse, error) {
	log.Info("[GetTime] received request")
	defer log.Info("[GetTime] completed request")

	var resp TimeSeriesDailyResponse
	err := s.callApi("TIME_SERIES_DAILY", map[string]string{"symbol": in.Symbol}, &resp)
	if err != nil {
		return nil, err
	}

	tsd := resp.TimeSeriesDaily
	ts := make([]*pb.TimeSeriesDataPoint, len(tsd))
	i := 0
	for k, v := range tsd {
		date, err := stringToDate(k)
		if err != nil {
			return nil, err
		}
		open, err := strconv.ParseFloat(v.Open, 64)
		if err != nil {
			return nil, err
		}
		ts[i] = &pb.TimeSeriesDataPoint{
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
	return &pb.GetTimeSeriesResponse{TimeSeries: ts}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("listening for gRPC on port %d\n", grpcPort)

	s := grpc.NewServer()
	svc := &server{}
	mustMapEnv(&svc.apiKey, "ALPHA_VANTAGE_API_KEY")
	pb.RegisterAlphaVantageAPIServer(s, svc)
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}
