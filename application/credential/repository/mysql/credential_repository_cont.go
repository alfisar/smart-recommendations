package repository

import (
	"smart-recommendation/domain"
	"smart-recommendation/internal/errorhandler"
)

// struct to store the connection that will be used
type CredentialRepositoryCont interface {
	GetByAplication(application string) (result domain.Credential, err errorhandler.ErrorData)
}
