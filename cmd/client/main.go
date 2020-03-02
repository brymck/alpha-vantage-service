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

func getConnection(addr string) (*grpc.ClientConn, error) {
	creds := credentials.NewClientTLSFromCert(nil, "")

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func main() {
	// Contact the server and print out its response.
	symbol := os.Args[1]

	// Set up a connection to the server.
	conn, err := getConnection(os.Getenv("ADDR"))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAlphaVantageAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r, err := c.GetQuote(ctx, &pb.GetQuoteRequest{Symbol: symbol})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("%v", r)
}
