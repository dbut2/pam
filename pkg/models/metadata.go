package models

import (
	"time"
)

type Metadata struct {
	id string
	owner *Account
	createdAt time.Time
	lastModifiedAt time.Time
}
