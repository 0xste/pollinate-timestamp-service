package model

import "time"

type Timestamp struct {
	EventTimestamp time.Time
	CommandId      string
}
