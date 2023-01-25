package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type BaseModel struct {
	BaseModelTimes

	ID uuid.UUID `bun:"type:uuid"`
}

type BaseModelTimes struct {
	CreatedAt time.Time  `bun:"type:timestampz,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time  `bun:"type:timestampz,nullzero,notnull,default:current_timestamp"`
	DeletedAt *time.Time `bun:"type:timestampz"`
}

type BaseModelResponse struct {
	BaseModelTimesResponse

	ID uuid.UUID `json:"id"`
}

type BaseModelTimesResponse struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
