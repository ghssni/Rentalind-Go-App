package repository_mock

import (
	"github.com/stretchr/testify/mock"
)

type RentalRepositoryMock struct {
	mock.Mock
}

func (m *RentalRepositoryMock) FindByID(id int) (models.Rental, error) {
	args := m.Called(id)
	return args.Get(0).(models.Rental), args.Error(1)
}

func (m *RentalRepositoryMock) FindAll() ([]models.Rental, error) {
	args := m.Called()
	return args.Get(0).([]models.Rental), args.Error(1)
}

func (m *RentalRepositoryMock) Create(rental models.Rental) error {
	args := m.Called(rental)
	return args.Error(0)
}

func (m *RentalRepositoryMock) Update(rental models.Rental) error {
	args := m.Called(rental)
	return args.Error(0)
}

func (m *RentalRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}