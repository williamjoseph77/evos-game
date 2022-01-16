package objects

import (
	"github.com/google/uuid"
)

// CreateCharacterNonSecureRequest ...
type CreateCharacterNonSecureRequest struct {
	Name   string
	RoleID int
	Power  int
	Wealth float64
}

// CreateCharacterSecureRequest ...
type CreateCharacterSecureRequest struct {
	Name     string
	RoleGUID uuid.UUID
	Power    int
	Wealth   float64
}

// CreateCharacterResponse ...
type CreateCharacterResponse struct {
	GUID uuid.UUID `json:"guid"`
}

type GetCharacterListResponse struct {
	GUID   uuid.UUID `json:"guid"`
	Name   string    `json:"name"`
	Power  int       `json:"power"`
	Wealth float64   `json:"wealth"`

	Role GetRoleListResponse `json:"role"`
}
