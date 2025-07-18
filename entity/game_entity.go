package entity

import "github.com/google/uuid"

type Game struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	Name string `json:"name"`

	Timestamp
}
