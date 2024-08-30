package testing

import (
	"Rentalind-Go-App/models"
	"Rentalind-Go-App/repository_mock"
	"Rentalind-Go-App/handler"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var rentalRepo = &repository_mock.RentalRepositoryMock{Mock: mock.Mock{}}
var rentalHandler = handler.RentalHandler{Repo: rentalRepo}

func TestGetRental(t *testing.T) {
	rentalExpected := models.Rental{
		ID:         1,
		BookID:     1,
		CustomerID: 1,
		RentDate:   "2022-01-01",
		ReturnDate: "2022-01-15",
	}

	rentalRepo.Mock.On("FindByID", 1).Return(rentalExpected, nil)

	res, err := rentalHandler.GetRental(1)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, rentalExpected.BookID, res.BookID, "BookID should be same")
	assert.Equal(t, rentalExpected.CustomerID, res.CustomerID, "CustomerID should be same")
	assert.Equal(t, rentalExpected.RentDate, res.RentDate, "RentDate should be same")
	assert.Equal(t, rentalExpected.ReturnDate, res.ReturnDate, "ReturnDate should be same")
}
