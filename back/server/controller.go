package server

import (
	"log"
	"net/http"

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

	id, err := writeUserToDb(user, hashedPassword)
	if err != nil {
		return err
	}

	err = processImage(c, user.Email)
	if err != nil {
		return err
	}

	uid, _ := generateUID()
	writeUidToDb(id, uid)
	err = sendConfirmationEmail(uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "User registered successfully", "user": user})
}

func login(c echo.Context) error {
	req := new(LoginDTO)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"validation_errors": err.Error()})
	}

	user, err := getUserByEmail(req.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}

	if !user.IsEnabled {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Account not verified"})
	}

	if user.IsDeleted {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Account has been deleted"})
	}

	token, err := makeJwtToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func validate(c echo.Context) error {
	uid := c.Param("uid")
	conf, err := getConfirmationByUid(uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Invalid token"})
	}

	err = enableUserById(conf.UserId)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Failed to enable user"})
	}

	err = deleteUidFromDb(uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Failed to delete token"})
	}

	return c.NoContent(http.StatusOK)
}
