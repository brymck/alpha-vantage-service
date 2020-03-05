package main

import (
	pb "github.com/brymck/alpha-vantage-service/genproto/brymck/alpha_vantage/v1"
	"github.com/brymck/alpha-vantage-service/pkg/models"
)

func asProtobufDate(d *models.Date) *pb.Date {
	return &pb.Date{
		Year:  d.Year,
		Month: d.Month,
		Day:   d.Day,
	}
}
