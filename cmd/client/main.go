package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/brymck/alpha-vantage-service/genproto"
)

func main() {
	address := os.Getenv("ADDR")

	creds := credentials.NewClientTLSFromCert(nil, "")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAlphaVantageAPIClient(conn)

	// Contact the server and print out its response.
	symbol := os.Args[1]
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetQuote(ctx, &pb.GetQuoteRequest{Symbol: symbol})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%v", r)
}