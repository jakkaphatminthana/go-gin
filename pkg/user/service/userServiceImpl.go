package service

import (
	"github.com/jakkaphatminthana/go-gin/entities"
	_providerRepository "github.com/jakkaphatminthana/go-gin/pkg/provider/repository"
	_userRepository "github.com/jakkaphatminthana/go-gin/pkg/user/repository"
)

type userServiceImpl struct {
	userRepository     _userRepository.UserRepository
	providerRepository _providerRepository.ProviderRepository
}

func NewUserServiceImpl(
	userRepository _userRepository.UserRepository,
	providerRepository _providerRepository.ProviderRepository,
) UserService {
	return &userServiceImpl{userRepository, providerRepository}
}

// implement
func (s *userServiceImpl) Create(user *entities.User, provider *entities.Provider) (*entities.User, error) {
	tx := s.userRepository.TransactionBegin()

	//TODO : add check providerID, provider
	//...

	// create user
	createdUser, err := s.userRepository.Create(tx, user)
	if err != nil {
		s.userRepository.TransactionRollback(tx)
		return nil, err
	}

	// login by social
	if provider != nil {
		provider.UserID = createdUser.ID
		_, err := s.providerRepository.Create(tx, provider)
		if err != nil {
			s.userRepository.TransactionRollback(tx)
			return nil, err
		}
	}

	if err := s.userRepository.TransactionCommit(tx); err != nil {
		return nil, err
	}

	return createdUser, nil
}

// implement
func (s *userServiceImpl) FindByEmail(email string) (*entities.User, error) {
	return s.userRepository.FindByEmail(email)
}
