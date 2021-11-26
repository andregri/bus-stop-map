package record

import (
	"context"
)

type ArrivalTimeRecord struct {
	StopCode    string
	BusLine     string
	ArrivalTime string
}

type ArrivalTimeHandler interface {
	CreateRecord(ctx context.Context, record ArrivalTimeRecord) error
	DeleteRecord(ctx context.Context, record ArrivalTimeRecord) error
}