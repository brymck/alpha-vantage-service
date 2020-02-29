package main

import (
	"fmt"
	"net"

	"github.com/brymck/helpers/env"
	"github.com/brymck/helpers/secrets"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

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
	port := env.GetPort("8080")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Infof("listening for gRPC on port %s", port)

	s := grpc.NewServer()
	apiKey, err := secrets.AccessSecret("alpha-vantage-api-key")
	if err != nil {
		log.Fatalf("failed to retrieve API key: %v", err)
	}
	api := alphavantage.NewApi(apiKey)
	app := &application{api: api}
	pb.RegisterAlphaVantageAPIServer(s, app)
	reflection.Register(s)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
