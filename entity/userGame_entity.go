package entity

import "github.com/google/uuid"

type UserGame struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID uuid.UUID `json:"user_id"`
	GameID uuid.UUID `json:"game_id"`
	Score  int       `json:"score"`

	User *User `gorm:"foreignKey:UserID"`
	Game *Game `gorm:"foreignKey:GameID"`

	Timestamp
}
