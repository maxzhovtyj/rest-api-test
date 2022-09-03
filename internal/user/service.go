package user

import (
	"context"
	"rest-api-test/pkg/logging"
)

type Service struct {
	storage *Storage
	logger  *logging.Logger
}

func NewService(repo *Storage, logger *logging.Logger) *Service {
	return &Service{
		storage: repo,
		logger:  logger,
	}
}

func (s *Service) CreateOne(ctx context.Context, user CreateUserDTO) (User, error) {
	// TODO generatePasswordHash
	// passwordHash := generatePasswordHash(password)
	//userInfo := User{
	//	Username:     user.Username,
	//	PasswordHash: "",
	//	Email:        user.Email,
	//}
	//createdUserId, err := s.storage.Create(ctx, userInfo)
	//if err != nil {
	//}
	//
	//return createdUserId, nil
	return User{}, nil
}
