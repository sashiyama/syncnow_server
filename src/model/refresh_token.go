package model

import "time"

type RefreshToken struct {
	Id        string    `json:"id" validate:"required"`
	Token     string    `json:"token" validate:"required"`
	ExpiresAt time.Time `json:"expires_at" validate:"required"`
}
