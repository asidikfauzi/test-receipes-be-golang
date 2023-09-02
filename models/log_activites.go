package models

import (
	"github.com/google/uuid"
	"time"
)

type LogActivity struct {
	Uuid      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;column:id;" json:"id"`
	TimeStamp time.Time `gorm:"not null;column:timestamp;" json:"timestamp"`
	Action    string    `gorm:"type:varchar(50);column:action;" json:"action"`
	IpAddress string    `gorm:"type:varchar(50);column:ip_address" json:"ip_address"`
	UserAgent string    `gorm:"type:varchar(120);column:user_agent" json:"user_agent"`
	Detail    string    `gorm:"type:varchar(255);column:detail" json:"detail"`
}
