package repository

import (
	"github.com/jakkaphatminthana/go-gin/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	TransactionBegin() *gorm.DB
	TransactionRollback(tx *gorm.DB) error
	TransactionCommit(tx *gorm.DB) error

	Create(tx *gorm.DB, entity *entities.User) (*entities.User, error)
	FindById(userId uint64) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
}
