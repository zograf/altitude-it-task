package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
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

	return c.JSON(http.StatusOK, echo.Map{"token": token, "is_admin": user.IsAdmin})
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

func getUserDetails(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	email := claims["email"].(string)
	//isAdmin := claims["is_admin"].(bool)
	user, err := getUserInfoByEmail(email)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	info := &UserInfo{
		Name:     user.Name,
		LastName: user.LastName,
		Birthday: user.Birthday.String(),
		Email:    user.Email,
	}

	return c.JSON(http.StatusOK, echo.Map{"user": info})
}

func updateUserDetails(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	email := claims["email"].(string)
	user := new(UserInfo)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid data"})
	}

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if user.Email != email {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Wrong user"})
	}

	err := updateUserInDb(*user)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	err = processImage(c, user.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

func updatePassword(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	email := claims["email"].(string)
	dto := new(UpdatePasswordDTO)

	if err := c.Bind(dto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid data"})
	}

	if err := c.Validate(dto); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	user, err := getUserByEmail(email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.OldPassword))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid email or password"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	}

	err = updateUserPassword(string(hashedPassword), email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Failed to update password"})
	}

	return c.NoContent(http.StatusOK)
}

func getAllUsers(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)

	is_admin := claims["is_admin"].(bool)
	if !is_admin {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	email := c.QueryParam("email")
	birthdayFrom := c.QueryParam("birthdayFrom")
	birthdayTo := c.QueryParam("birthdayTo")
	e := c.QueryParam("enabled")

	var enabled *bool = nil
	if e == "true" {
		value := true
		enabled = &value
	} else if e == "false" {
		value := false
		enabled = &value
	}

	users, total, err := getUsers(email, birthdayFrom, birthdayTo, enabled, pageSize, offset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if users == nil {
		users = make([]User, 0)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"users":    users,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}
