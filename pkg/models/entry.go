package models

import (
	"time"
)

type Entry struct {
	ID    string
	Date  time.Time
	Entry string
}
