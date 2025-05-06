package repository

import (
	"github.com/jakkaphatminthana/go-gin/entities"
	"gorm.io/gorm"
)

type ProviderRepository interface {
	Create(tx *gorm.DB, entity *entities.Provider) (*entities.Provider, error)
	FindByProviderIDAndProvider(provider, providerID string) (*entities.Provider, error)
}
