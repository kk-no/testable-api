package store

import (
	"context"
	"errors"

	"github.com/kk-no/testable-api/models"
)

type MockUser struct{}

func NewMockUser() *MockUser {
	return &MockUser{}
}

var _ UserHandler = (*MockUser)(nil)

func (m *MockUser) FindByID(ctx context.Context, id string) (*models.User, error) {
	return &models.User{ID: id, Name: "tester"}, nil
}

func (m *MockUser) Create(ctx context.Context, user *models.User) error {
	if user == nil {
		return errors.New("argument gets error")
	}
	return nil
}

func (m *MockUser) Update(ctx context.Context, user *models.User) error {
	if user == nil {
		return errors.New("argument gets error")
	}
	return nil
}
