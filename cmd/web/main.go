package main

import (
	"github.com/brymck/helpers/secrets"
	"github.com/brymck/helpers/servers"
	log "github.com/sirupsen/logrus"

	pb "github.com/brymck/alpha-vantage-service/genproto/brymck/alpha_vantage/v1"
	"github.com/brymck/alpha-vantage-service/pkg/alphavantage"
)

type application struct {
	api *alphavantage.AlphaVantageApi
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
