package models

import (
	"time"
)

type Metadata struct {
	ID        string
	Author    string
	Timestamp time.Time
}
