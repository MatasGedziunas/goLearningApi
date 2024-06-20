package database

import (
	"errors"
	"fmt"

	models "github.com/MatasGedziunas/goLearningApi/cmd/api"
)

var database []models.Account = []models.Account{
	createMockAccount(50, "John"),
	createMockAccount(100, "Alice"),
}

func createMockAccount(balance int, username string) models.Account {
	return models.Account{
		Balance:  balance,
		Username: username,
	}
}

func GetAccount(username string) (models.Account, error) {
	for _, account := range database {
		if account.Username == username {
			return account, nil
		}
	}
	return models.Account{}, errors.New(fmt.Sprintf("Account with username %v not found", username))
}
