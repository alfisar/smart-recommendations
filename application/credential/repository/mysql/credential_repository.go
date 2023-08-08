package repository

import (
	"smart-recommendation/domain"
	"smart-recommendation/internal/errorhandler"

	"gorm.io/gorm"
)

type CredentialRepoesitory struct {
	conn *gorm.DB
}

// init new repository credential
func NewCredentialRepository(conn *gorm.DB) *CredentialRepoesitory {
	return &CredentialRepoesitory{
		conn: conn,
	}
}

func (obj *CredentialRepoesitory) GetByAplication(application string) (result domain.Credential, err errorhandler.ErrorData) {
	errData := obj.conn.Debug().Table("credential").Where("application = ?", application).First(&result).Error
	if errData != nil {
		err = errorhandler.ErrorRepo(errData)
	}

	return

}
