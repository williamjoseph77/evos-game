package db

import (
	"log"

	"github.com/williamjoseph77/evos/domains"

	"github.com/go-pg/pg/v10"
)

// CreateCharacter create new character
func CreateCharacter(dbi *pg.DB, character domains.Character) (*domains.Character, error) {
	tx, err := dbi.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Model(&character).Insert()
	if err != nil {
		log.Println("error inserting new character ", err)
		return nil, err
	}

	//nolint:errcheck
	defer tx.Rollback()

	err = tx.Commit()

	if err != nil {
		log.Println("error committing new character")
		return nil, err
	}

	return &character, nil
}

// GetCharacters get all characters
func GetCharacters(dbi *pg.DB) ([]domains.Character, error) {
	var characters []domains.Character

	query := dbi.Model(&characters).Relation("Role")

	err := query.Select()
	if err != nil {
		log.Println("error getting characters")
		return nil, err
	}

	return characters, nil
}
