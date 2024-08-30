package testing

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"Rentalind-Go-App/handlers"
	"Rentalind-Go-App/models"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserFromContext(c echo.Context) (*models.User, error) {
	res := m.Mock.Called(c)

	if res.Get(1) != nil {
		return nil, res.Get(1).(error)
	}

	user := res.Get(0).(models.User)
	return &user, nil
}

func TestRentProducts(t *testing.T) {
	e := echo.New()

	mockUserRepo := new(MockUserRepository)
	handlers.SetUserRepository(mockUserRepo)

	user := models.User{Email: "user@example.com"}
	mockUserRepo.On("GetUserFromContext", mock.Anything).Return(&user, nil)

	req := httptest.NewRequest(http.MethodPost, "/rent", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/rent")

	err := handlers.RentProducts(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Rental successful")
}