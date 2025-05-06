package service

import "github.com/jakkaphatminthana/go-gin/entities"

type UserService interface {
	Create(user *entities.User, provider *entities.Provider) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
}
