package model

import "time"

type Timestamp struct {
	EventTimestamp time.Time `json:"event_timestamp"`
	CommandId      string    `json:"command_id"`
}
