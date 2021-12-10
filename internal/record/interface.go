package record

import (
	"context"
	"time"
)

type ArrivalTimeRecord struct {
	StopCode    string
	BusLine     string
	ArrivalTime time.Time
}

type ArrivalTimeHandler interface {
	CreateRecord(ctx context.Context, record ArrivalTimeRecord) error
	SearchRecord(ctx context.Context, stopCode string) ([]ArrivalTimeRecord, error)
	DeleteRecord(ctx context.Context, id int) error
}
