package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
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

func (srv *Server) createTables() {
	var queries []string
	db, err := pgx.Connect(context.Background(), CONN_STRING)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	queries = append(queries, "DROP TABLE IF EXISTS Users")
	queries = append(queries, "DROP TABLE IF EXISTS Admin")
	queries = append(queries,
		`CREATE TABLE Users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            lastname VARCHAR(100) NOT NULL,
            birthday DATE NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            password VARCHAR(100) NOT NULL,
            is_enabled BOOLEAN DEFAULT TRUE,
            is_deleted BOOLEAN DEFAULT FALSE,
            google_id VARCHAR(100) UNIQUE
        )`)
	queries = append(queries,
		`CREATE TABLE Admin (
			id SERIAL PRIMARY KEY,
			email VARCHAR(100) UNIQUE NOT NULL,
			password VARCHAR(100) NOT NULL
        )`)

	for _, query := range queries {
		if _, err := db.Exec(context.Background(), query); err != nil {
			log.Fatalf("failed executing query: %v, error: %v", query, err)
		}
	}
}

func (srv *Server) Run() {
	srv.createTables()

	e := echo.New()

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	v := validator.New()
	v.RegisterValidation("date", dateValidation)
	e.Validator = &CustomValidator{validator: v}

	//e.POST("/login", login)
	e.POST("/register", register)

	url := fmt.Sprintf("%s%s", srv.Ip, srv.Port)
	e.Logger.Fatal(e.Start(url))
}
