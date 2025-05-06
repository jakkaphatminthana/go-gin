package repository

import (
	"github.com/jakkaphatminthana/go-gin/database"
	"github.com/jakkaphatminthana/go-gin/entities"
	"github.com/jakkaphatminthana/go-gin/pkg/custom"
	_userException "github.com/jakkaphatminthana/go-gin/pkg/user/exception"
	"github.com/jakkaphatminthana/go-gin/utils"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db     database.Database
	logger utils.Logger
}

func NewUserRepositoryImpl(db database.Database, logger utils.Logger) UserRepository {
	return &userRepositoryImpl{db, logger}
}

// implement
func (r *userRepositoryImpl) TransactionBegin() *gorm.DB {
	tx := r.db.Connect()
	return tx.Begin()
}

// implement
func (r *userRepositoryImpl) TransactionRollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

// implement
func (r *userRepositoryImpl) TransactionCommit(tx *gorm.DB) error {
	return tx.Commit().Error
}

// implement
func (r *userRepositoryImpl) Create(tx *gorm.DB, entity *entities.User) (*entities.User, error) {
	conn := r.db.Connect()
	if tx != nil {
		conn = tx
	}

	user := new(entities.User)

	if err := conn.Create(entity).Scan(user).Error; err != nil {
		r.logger.Errorf("Creating user failed: %s", err.Error())
		return nil, custom.ErrorInternalServerError("Creating user failed")
	}
	return user, nil
}

// implement
func (r *userRepositoryImpl) FindById(userId uint64) (*entities.User, error) {
	user := new(entities.User)

	if err := r.db.Connect().Where("id = ?", userId).First(user).Error; err != nil {
		r.logger.Errorf("Find user by ID failed: %s", err.Error())
		return nil, &_userException.UserNotFound{ID: userId}
	}

	return user, nil
}

// implement
func (r *userRepositoryImpl) FindByEmail(email string) (*entities.User, error) {
	user := new(entities.User)

	if err := r.db.Connect().Where("email = ?", email).First(&user).Error; err != nil {
		r.logger.Errorf("Find user by email failed: %s", err.Error())
		return nil, err
	}

	return user, nil
}
