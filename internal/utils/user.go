package utils

import (
	"encoding/json"
	"os"

	"github.com/seatedro/v3/internal/models"
)

func GetUserData() (models.User, error) {
	userData, err := os.ReadFile("content/user.json")
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = json.Unmarshal(userData, &user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
