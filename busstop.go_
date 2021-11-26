package main

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/go-redis/redis/v8"
)

type BusArrivalTime struct {
	StopCode    string
	BusLine     string
	ArrivalTime string
}

// Encode BusArrivalTime struct with gob and save to redis
func (bat *BusArrivalTime) Save(ctx context.Context, rdb *redis.Client) error {
	// gob encoding
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(bat)
	if err != nil {
		return err
	}

	// save to redis
	err = rdb.Set(ctx, bat.StopCode, buf.Bytes(), 0).Err()

	return err
}

func Load(ctx context.Context, rdb *redis.Client, key string) (*BusArrivalTime, error) {
	// load from redis
	value, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	// gob decoding
	bytesReader := bytes.NewReader([]byte(value))
	decoder := gob.NewDecoder(bytesReader)

	var bat BusArrivalTime
	err = decoder.Decode(&bat)
	if err != nil {
		return nil, err
	}

	return &bat, nil
}
