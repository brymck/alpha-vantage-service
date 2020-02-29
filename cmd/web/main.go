package main

import (
	"github.com/brymck/helpers/secrets"
	"github.com/brymck/helpers/servers"
	log "github.com/sirupsen/logrus"

	pb "github.com/brymck/alpha-vantage-service/genproto"
	"github.com/brymck/alpha-vantage-service/pkg/alphavantage"
)

func init() {
	formatter := &log.TextFormatter{FullTimestamp: true}
	log.SetFormatter(formatter)
}

type application struct {
	api *alphavantage.Api
}

func main() {
	// Get Alpha Vantage API key
	apiKey, err := secrets.AccessSecret("alpha-vantage-api-key")
	if err != nil {
		log.Fatalf("failed to retrieve API key: %v", err)
	}
	api := alphavantage.NewApi(apiKey)

	app := &application{api: api}

	// Start server
	s := servers.NewGrpcServer()
	pb.RegisterAlphaVantageAPIServer(s.Server, app)
	s.Serve()
}
