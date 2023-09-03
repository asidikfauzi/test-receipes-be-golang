package models

import (
	"github.com/google/uuid"
	"time"
)

type LogActivity struct {
	Uuid      uuid.UUID `gorm:"uuid;default:uuid_generate_v4();primary_key;column:id;" json:"id"`
	TimeStamp time.Time `json:"timestamp"`
	Action    string    `json:"action"`
	IpAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	Detail    string    `json:"detail"`
}
