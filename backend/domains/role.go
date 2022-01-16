package domains

import "github.com/google/uuid"

const WizardID = 1
const ElfID = 2

type Role struct {
	ID   int       `pg:",pk" json:"id"`
	GUID uuid.UUID `json:"guid"`
	Name string    `json:"name"`

	// nolint:staticcheck
	tableName struct{} `pg:"roles"`
}
