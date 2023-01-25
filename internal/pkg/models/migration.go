package models

import "time"

type Migrations []*Migration

type Migration struct {
	ID        uint16
	AppliedAt time.Time
}
