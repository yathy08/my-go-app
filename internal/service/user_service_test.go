package service

import (
    "errors"
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/yourusername/my-go-app/internal/domain"
    "github.com/yourusername/my-go-app/internal/repository"
)

// Mock repository
type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) GetUserByUsername(username string) (*domain.User, error) {
    args := m.Called(username)
    if user, ok := args.Get(0).(*domain.User); ok {
        return user, args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *MockUserRepository) SaveUser(user *domain.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func TestUserService_Login(t *testing.T) {
    mockRepo := new(MockUserRepository)
    userService := NewUserService(mockRepo)

    testUser := &domain.User{ID: "1", Username: "test", Password: "hashedpassword"}

    mockRepo.On("GetUserByUsername", "test").Return(testUser, nil)
    token, err := userService.Login("test", "hashedpassword")

    assert.Nil(t, err)
    assert.NotEmpty(t, token)
    mockRepo.AssertExpectations(t)
}
