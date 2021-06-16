package models

import "time"

type Activity struct {
	metadata Metadata
	Date time.Time
}
