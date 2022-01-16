package services

import (
	"fmt"

	"github.com/williamjoseph77/evos/db"
	"github.com/williamjoseph77/evos/domains"
	"github.com/williamjoseph77/evos/objects"

	"github.com/go-pg/pg/v10"
)

// validateCreateCharacterSecurePayload validate request payload
func validateCreateCharacterSecurePayload(request objects.CreateCharacterSecureRequest) string {
	if request.Name == "" {
		return "Name"
	}

	if !(request.Power > 0 && request.Power < 100) {
		return "Power"
	}

	if request.Wealth < 0 {
		return "Wealth"
	}

	return ""
}

// validateCreateCharacterNonSecurePayload validate request payload
func validateCreateCharacterNonSecurePayload(request objects.CreateCharacterNonSecureRequest) string {
	if request.Name == "" {
		return "Name"
	}

	if request.RoleID != domains.ElfID && request.RoleID != domains.WizardID {
		return "RoleID"
	}

	if !(request.Power > 0 && request.Power < 100) {
		return "Power"
	}

	if request.Wealth < 0 {
		return "Wealth"
	}

	return ""
}

// CreateCharacterNonSecure ...
func CreateCharacterNonSecure(dbi *pg.DB, request objects.CreateCharacterNonSecureRequest) (*domains.Character, error) {
	fieldErr := validateCreateCharacterNonSecurePayload(request)
	if fieldErr != "" {
		return nil, fmt.Errorf("nilai input %s tidak valid", fieldErr)
	}

	// EXPLANATION:
	// This is app level error check, so we can get "human error text" instead of postgres error message
	// But I create a feature so this name column needs to be unique on DB too, so we can't accidentally create records with same name
	// You can test if we can insert same name (case insensitive) by commenting line line 62-69
	isExist, err := db.IsCharacterExistByName(dbi, request.Name)
	if err != nil {
		return nil, err
	}

	if isExist {
		return nil, fmt.Errorf("nama sudah digunakan, silakan input nama lain")
	}

	newCharacter, err := db.CreateCharacter(dbi, domains.Character{
		Name:   request.Name,
		RoleID: request.RoleID,
		Power:  request.Power,
		Wealth: request.Wealth,
	})

	if err != nil {
		return nil, err
	}

	return newCharacter, nil
}

// CreateCharacterSecure ...
func CreateCharacterSecure(dbi *pg.DB, request objects.CreateCharacterSecureRequest) (*domains.Character, error) {
	fieldErr := validateCreateCharacterSecurePayload(request)
	if fieldErr != "" {
		return nil, fmt.Errorf("nilai input %s tidak valid", fieldErr)
	}

	// EXPLANATION:
	// This is app level error check, so we can get "human error text" instead of postgres error message
	// But I create a feature so this name column needs to be unique on DB too, so we can't accidentally create records with same name
	// You can test if we can insert same name (case insensitive) by commenting line line 96-103
	isExist, err := db.IsCharacterExistByName(dbi, request.Name)
	if err != nil {
		return nil, err
	}

	if isExist {
		return nil, fmt.Errorf("nama sudah digunakan, silakan input nama lain")
	}

	role, err := db.GetRoleByGUID(dbi, request.RoleGUID)
	if err != nil {
		return nil, err
	}

	newCharacter, err := db.CreateCharacter(dbi, domains.Character{
		Name:   request.Name,
		RoleID: role.ID,
		Power:  request.Power,
		Wealth: request.Wealth,
	})
	if err != nil {
		return nil, err
	}

	return newCharacter, nil
}

// GetCharacterList ...
func GetCharacterList(dbi *pg.DB) ([]domains.Character, error) {
	characters, err := db.GetCharacters(dbi)
	if err != nil {
		return nil, err
	}

	return characters, nil
}
