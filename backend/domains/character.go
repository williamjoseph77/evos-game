package domains

import "github.com/google/uuid"

type Character struct {
	ID     int       `pg:",pk" json:"id"`
	GUID   uuid.UUID `json:"guid"`
	Name   string    `json:"name"`
	RoleID int       `json:"roleID"`
	Power  int       `json:"power"`
	Wealth float64   `json:"wealth"`

	Role Role `pg:"fk:role_id"`

	// nolint:staticcheck
	tableName struct{} `pg:"characters"`
}
