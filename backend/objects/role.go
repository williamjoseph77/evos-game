package objects

import "github.com/google/uuid"

type GetRoleListResponse struct {
	GUID uuid.UUID `json:"guid"`
	Name string    `json:"name"`
}
