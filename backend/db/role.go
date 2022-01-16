package db

import (
	"log"

	"github.com/williamjoseph77/evos/domains"

	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

// GetRoles get all roles
func GetRoles(dbi *pg.DB) ([]domains.Role, error) {
	var roles []domains.Role

	query := dbi.Model(&roles)

	err := query.Select()
	if err != nil {
		log.Println("error getting roles")
		return nil, err
	}

	return roles, nil
}

//GetRoleByGUID get role by guid
func GetRoleByGUID(dbi *pg.DB, guid uuid.UUID) (*domains.Role, error) {
	var role domains.Role

	query := dbi.Model(&role).Where("guid = ?", guid)
	err := query.Select()
	if err != nil {
		log.Println("error getting role by guid")
		return nil, err
	}

	return &role, nil
}
