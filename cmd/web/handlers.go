package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	pb "github.com/brymck/alpha-vantage-service/genproto"
)

func trackTime(start time.Time, log *logrus.Entry) {
	elapsed := time.Since(start)
	log.Infof("completed in %d ms", elapsed.Milliseconds())
}

func (app *application) GetQuote(_ context.Context, in *pb.GetQuoteRequest) (*pb.GetQuoteResponse, error) {
	log := logrus.WithFields(logrus.Fields{"method": "GetQuote", "symbol": in.Symbol})
	log.Info("received request")
	defer trackTime(time.Now(), log)

	q, err := app.api.GetQuote(in.Symbol)
	if err != nil {
		return nil, err
	}

	return &pb.GetQuoteResponse{
		Open:             q.Open,
		High:             q.High,
		Low:              q.Low,
		Price:            q.Price,
		Volume:           q.Volume,
		LatestTradingDay: asProtobufDate(q.LatestTradingDay),
		PreviousClose:    q.PreviousClose,
	}, nil
}

func (app *application) GetTimeSeries(_ context.Context, in *pb.GetTimeSeriesRequest) (*pb.GetTimeSeriesResponse, error) {
	log := logrus.WithFields(logrus.Fields{"method": "GetTimeSeries", "symbol": in.Symbol})
	log.Info("received request")
	defer trackTime(time.Now(), log)

	timeSeriesDaily, err := app.api.GetTimeSeriesDaily(in.Symbol)
	if err != nil {
		return nil, err
	}

	ts := make([]*pb.TimeSeriesDataPoint, len(timeSeriesDaily))
	for i, item := range timeSeriesDaily {
		ts[i] = &pb.TimeSeriesDataPoint{
			Date:   asProtobufDate(item.Date),
			Open:   item.Open,
			High:   item.High,
			Low:    item.Low,
			Close:  item.Close,
			Volume: item.Volume,
		}
	}
	return &pb.GetTimeSeriesResponse{TimeSeries: ts}, nil
}
