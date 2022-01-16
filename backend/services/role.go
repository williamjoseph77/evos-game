package services

import (
	"github.com/williamjoseph77/evos/db"
	"github.com/williamjoseph77/evos/domains"

	"github.com/go-pg/pg/v10"
)

// GetRoleList ...
func GetRoleList(dbi *pg.DB) ([]domains.Role, error) {
	roles, err := db.GetRoles(dbi)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
