package server

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pquerna/otp/totp"
)

func generateTOTPSecret(email string) (string, error) {
	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "altitude-test-task",
		AccountName: email,
	})
	if err != nil {
		return "", err
	}
	return secret.Secret(), nil
}

func generateUID() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func makeJwtToken(user *User) (string, error) {
	claims := jwt.MapClaims{
		"email":    user.Email,
		"is_admin": user.IsAdmin,
		//"expires":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	return tokenString, err
}

func processImage(c echo.Context, imageName string) error {
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
	extension := fmt.Sprintf(".%s", strings.Split(file.Filename, ".")[1])
	filePath := filepath.Join(uploadDir, imageName+extension)

	dst, err := os.Create(filePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return nil
}
