package repository

import (
	"github.com/jakkaphatminthana/go-gin/database"
	"github.com/jakkaphatminthana/go-gin/entities"
	"github.com/jakkaphatminthana/go-gin/pkg/custom"
	_providerException "github.com/jakkaphatminthana/go-gin/pkg/provider/exception"
	"github.com/jakkaphatminthana/go-gin/utils"
	"gorm.io/gorm"
)

type providerRepositoryImpl struct {
	db     database.Database
	logger utils.Logger
}

func NewProviderRepositoryImpl(db database.Database, logger utils.Logger) ProviderRepository {
	return &providerRepositoryImpl{db, logger}
}

func (r *providerRepositoryImpl) Create(tx *gorm.DB, entity *entities.Provider) (*entities.Provider, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	provider := new(entities.Provider)

	if err := conn.Create(entity).Scan(provider).Error; err != nil {
		r.logger.Errorf("Creating provider failed: %s", err.Error())
		return nil, custom.ErrorInternalServerError("Creating provider failed")
	}
	return provider, nil
}

func (r *providerRepositoryImpl) FindByProviderIDAndProvider(providerID, provider string) (*entities.Provider, error) {
	result := new(entities.Provider)

	if err := r.db.Connect().Where("provider_id = ? AND provider = ?", providerID, provider).First(result).Error; err != nil {
		r.logger.Errorf("Find provider by ID failed: %s", err.Error())
		return nil, &_providerException.ProviderNotFound{ProviderID: providerID, Provider: provider}
	}
	return result, nil
}
