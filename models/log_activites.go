package models

import (
	"github.com/google/uuid"
	"time"
)

type LogActivity struct {
	Uuid      uuid.UUID `json:"id"`
	TimeStamp time.Time `json:"timestamp"`
	Action    string    `json:"action"`
	IpAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	Detail    string    `json:"detail"`
}
