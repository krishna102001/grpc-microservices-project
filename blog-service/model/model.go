package model

import "github.com/google/uuid"

type Blog struct {
	Id          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string    `gorm:"type:varchar(300);not null"`
	Description string    `gorm:"type:varchar(1000);"`
	AuthorId    string    `gorm:"not null"`
}
