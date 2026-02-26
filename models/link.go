package models

import (
	"time"
)

// Modelo de url acortado

type Link struct {
	ID          int       `bson:"_id,omitempty" json:"id"`
	OriginalURL string    `bson:"original_url" json:"original_url"`
	ShortCode   string    `bson:"short_code" json:"short_code"`
	Clicks      int       `bson:"clicks" json:"clicks"`
	UserID      int       `bson:"user_id" json:"user_id"`
	UserEmail   string    `bson:"user_email" json:"user_email"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
}
