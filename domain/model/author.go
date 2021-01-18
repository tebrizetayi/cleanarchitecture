package model

import "github.com/google/uuid"

type Author struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}
