package storage

import (
	"github.com/kaynetik/modular-monolith-example/internal/pkg/models"
	"github.com/stretchr/testify/assert"
)

func (s *RepositorySuite) TestCreateUser() {
	newUser, err := getTestUser()
	assert.NoError(s.T(), err)
	assert.NoError(s.T(), s.repo.CreateUser(newUser))

	result, err := s.repo.GetUserByID(newUser.ID)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), result)
	assert.Equal(s.T(), newUser.ID, result.ID)
	assert.Equal(s.T(), newUser.FirstName, result.FirstName)
	assert.Equal(s.T(), result.CreatedAt, result.UpdatedAt)
}

func (s *RepositorySuite) TestGetUserByID() {
	testUser, err := getTestUser()
	assert.NoError(s.T(), err)
	assert.NoError(s.T(), s.repo.CreateUser(testUser))

	result, err := s.repo.GetUserByID(testUser.ID)
	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), result)
	assert.Equal(s.T(), testUser.ID, (*result).ID)
}

func getTestUser() (*models.User, error) {
	// Just a placeholder for a test user. In a real app, you would probably use a factory pattern.
	return &models.User{}, nil
}
