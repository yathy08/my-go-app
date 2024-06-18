package service

import (
    "errors"
    "github.com/yourusername/my-go-app/internal/domain"
    "github.com/yourusername/my-go-app/internal/repository"
    "github.com/yourusername/my-go-app/pkg/utils"
)

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) Login(username, password string) (string, error) {
    user, err := s.repo.GetUserByUsername(username)
    if err != nil || user == nil {
        return "", errors.New("user not found")
    }

    if !utils.CheckPasswordHash(password, user.Password) {
        return "", errors.New("invalid password")
    }

    token, err := utils.GenerateJWT(user.ID, user.Username, user.Role)
    if err != nil {
        return "", err
    }

    return token, nil
}
