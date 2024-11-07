package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func googleAuthHandler(c echo.Context) error {
	var request struct {
		Token string `json:"token"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	oauth2Service, err := oauth2.NewService(context.Background(), option.WithAPIKey(GOOGLE_ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create OAuth2 service"})
	}

	tokenInfo, err := oauth2Service.Tokeninfo().IdToken(request.Token).Do()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token"})
	}

	email := tokenInfo.Email

	// If no account then register
	_, err = getUserByEmail(email)
	if err != nil {
		writePartialUserToDb(email)
	}

	jwtToken, err := makeJwtToken(&User{
		Email:   email,
		IsAdmin: false,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create JWT"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": jwtToken})
}
