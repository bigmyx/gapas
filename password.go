package main

import (
	"github.com/google/uuid"
	"time"
)

type Pass struct {
	Id   uuid.UUID     `json:"id"`
	Pass string        `json:"pass_value"`
	TTL  time.Duration `json:"ttl"`
}
