package user

import (
	"context"
	"rest-api-test/pkg/logging"
)

type Service struct {
	storage Storage
	logger  *logging.Logger
	cache   Cache
}

func NewService(repo Storage, logger *logging.Logger, cache Cache) *Service {
	return &Service{
		storage: repo,
		logger:  logger,
		cache:   cache,
	}
}

func (s *Service) CreateOne(ctx context.Context, user CreateUserDTO) (User, error) {
	s.cache.SetJson(ctx)
	s.cache.GetJson(ctx)
	return User{}, nil
}
