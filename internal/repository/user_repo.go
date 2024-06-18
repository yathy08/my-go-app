package repository

import "github.com/yathy08/my-go-app/internal/domain"

type UserRepository interface {
    GetUserByUsername(username string) (*domain.User, error)
    SaveUser(user *domain.User) error
}

type InMemoryUserRepository struct {
    users map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
    return &InMemoryUserRepository{
        users: make(map[string]*domain.User),
    }
}

func (repo *InMemoryUserRepository) GetUserByUsername(username string) (*domain.User, error) {
    user, exists := repo.users[username]
    if !exists {
        return nil, nil
    }
    return user, nil
}

func (repo *InMemoryUserRepository) SaveUser(user *domain.User) error {
    repo.users[user.Username] = user
    return nil
}
