package server

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func register(c echo.Context) error {
	user := new(RegisterDTO)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid data"})
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	}

	err = writeUserToDb(user, hashedPassword)
	if err != nil {
		return err
	}

	err = processImage(c)
	if err != nil {
		return err
	}

	//_, err = getUserByEmail(user.Email)
	//if err != nil {
	//	return err
	//}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User registered successfully", "user": user})
}

func processImage(c echo.Context) error {
	file, err := c.FormFile("image")
	// I don't need an image file
	if err != nil {
		return nil
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to open image file"})
	}
	defer src.Close()

	uploadDir := "./img"
	filePath := filepath.Join(uploadDir, file.Filename)

	dst, err := os.Create(filePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save image file"})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save image data"})
	}

	return nil
}
