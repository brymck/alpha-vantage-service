package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1beta1"
	"github.com/brymck/helpers/env"
	"github.com/sirupsen/logrus"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/brymck/alpha-vantage-service/genproto"
)

func init() {
	formatter := &logrus.TextFormatter{FullTimestamp: true}
	logrus.SetFormatter(formatter)
}

type GlobalQuoteResponse struct {
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

type TimeSeriesDailyResponse struct {
	TimeSeriesDaily map[string]struct {
		Open string `json:"1. open"`
	} `json:"Time Series (Daily)"`
}

type server struct {
	apiKey string
}

func trackTime(start time.Time, log *logrus.Entry) {
	elapsed := time.Since(start)
	log.Infof("completed in %d ms", elapsed.Milliseconds())
}

func (s *server) GetQuote(_ context.Context, in *pb.GetQuoteRequest) (*pb.GetQuoteResponse, error) {
	log := logrus.WithFields(logrus.Fields{"method": "GetQuote", "symbol": in.Symbol})
	log.Info("received request")
	defer trackTime(time.Now(), log)

	var resp GlobalQuoteResponse
	err := s.callApi("GLOBAL_QUOTE", map[string]string{"symbol": in.Symbol}, &resp)
	if err != nil {
		return nil, err
	}
	gq := resp.GlobalQuote
	open, err := strconv.ParseFloat(gq.Open, 64)
	if err != nil {
		return nil, err
	}
	high, err := strconv.ParseFloat(gq.High, 64)
	if err != nil {
		return nil, err
	}
	low, err := strconv.ParseFloat(gq.Low, 64)
	if err != nil {
		return nil, err
	}
	price, err := strconv.ParseFloat(gq.Price, 64)
	if err != nil {
		return nil, err
	}
	volume, err := strconv.ParseFloat(gq.Volume, 64)
	if err != nil {
		return nil, err
	}
	previousClose, err := strconv.ParseFloat(gq.PreviousClose, 64)
	if err != nil {
		return nil, err
	}

	return &pb.GetQuoteResponse{
		Open:          open,
		High:          high,
		Low:           low,
		Price:         price,
		Volume:        volume,
		PreviousClose: previousClose,
	}, nil
}

func (s *server) GetTimeSeries(_ context.Context, in *pb.GetTimeSeriesRequest) (*pb.GetTimeSeriesResponse, error) {
	log := logrus.WithFields(logrus.Fields{"method": "GetTimeSeries", "symbol": in.Symbol})
	log.Info("received request")
	defer trackTime(time.Now(), log)

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

func getSecret(secret string) (string, error) {
	log := logrus.WithField("secret", secret)
	envKey := strings.ReplaceAll(strings.ToUpper(secret), "-", "_")
	apiKeyFromEnv := os.Getenv(envKey)
	if apiKeyFromEnv != "" {
		log.Infof("using value of environment variable %s for secret %s", envKey, secret)
		return apiKeyFromEnv, nil
	}

	project := env.MustGetEnv("PROJECT_ID")
	log = log.WithField("project", project)

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to set up client: %w", err)
	}

	name := fmt.Sprintf("%s/%s", project, secret)
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{Name: name}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return "", fmt.Errorf("failed to access secret: %v", err)
	}
	log.Infof("used Secrets Manager to retrieve secret %s", envKey)
	return string(result.Payload.Data), nil
}

func main() {
	port := env.GetPort("8080")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	logrus.Infof("listening for gRPC on port %s", port)

	s := grpc.NewServer()
	project := os.Getenv("PROJECT_ID")
	apiKey, err := getSecret("alpha-vantage-api-key")
	if err != nil {
		logrus.Fatalf("failed to retrieve API key: %v", err)
	}
	svc := &server{
		apiKey: apiKey,
	}
	pb.RegisterAlphaVantageAPIServer(s, svc)
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("failed to server: %v", err)
	}
}
