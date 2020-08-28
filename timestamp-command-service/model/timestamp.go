package model

import "time"

type Timestamp struct {
	EventTimestamp   time.Time
	CommandTimestamp time.Time
	CommandId        string
}
