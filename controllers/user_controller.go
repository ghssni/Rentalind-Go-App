package controllers

import (
	"net/http"
	"Rentalind-Go-App/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RegisterUser menangani registrasi pengguna baru
func RegisterUser(c echo.Context) error {
	// Mendapatkan instance database dari konteks
	db := c.Get("db").(*gorm.DB)

	// Mengikat data JSON dari request ke struct User
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
		})
	}

	// Meng-hash password dan menyimpan pengguna ke database
	user.PasswordHash = "hashed_password" // Disini Anda bisa menambahkan fungsi untuk hashing password
	result := db.Create(&user)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "User already exists or invalid data",
		})
	}

	// Mengembalikan respons sukses dengan informasi pengguna yang terdaftar
	return c.JSON(http.StatusCreated, echo.Map{
		"user_id": user.ID,
		"name": user.Name,
		"address": user.Address,
		"email":   user.Email,
	})
}

// LoginUser menangani login pengguna dan pembuatan token JWT
func LoginUser(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)

	// Mengikat data JSON dari request ke struct User
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
		})
	}

	// Memverifikasi pengguna berdasarkan email dan password hash
	dbUser := new(models.User)
	result := db.Where("email = ? AND password_hash = ?", user.Email, "hashed_password").First(&dbUser)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found or invalid credentials",
		})
	}

	// Membuat JWT token setelah autentikasi berhasil
	// Disini seharusnya Anda menambahkan logika untuk membuat JWT token yang aman

	db.Save(&dbUser)
	return c.JSON(http.StatusOK, echo.Map{
		"token": dbUser.JWTToken,
	})
}