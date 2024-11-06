package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Validator

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func dateValidation(fl validator.FieldLevel) bool {
	date := fl.Field().String()
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

// Server

type Server struct {
	Port string
	Ip   string
}

func New() *Server {
	p := ":8080"
	ip := "localhost"

	srv := &Server{
		Port: p,
		Ip:   ip,
	}

	return srv
}

func (srv *Server) Run() {
	e := echo.New()

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v := validator.New()
	v.RegisterValidation("date", dateValidation)
	e.Validator = &CustomValidator{validator: v}

	//e.POST("/login", login)
	e.POST("/register", register)

	url := fmt.Sprintf("%s%s", srv.Ip, srv.Port)
	e.Logger.Fatal(e.Start(url))
}
