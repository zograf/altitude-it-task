package server

import "time"

type RegisterDTO struct {
	Name           string `form:"name" validate:"required"`
	LastName       string `form:"last_name" validate:"required"`
	Email          string `form:"email" validate:"required,email"`
	Password       string `form:"password" validate:"required"`
	RepeatPassword string `form:"repeat_password" validate:"required,eqfield=Password"`
	Birthday       string `form:"birthday" validate:"required,date"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID        int
	Name      string
	LastName  string
	Birthday  time.Time
	Email     string
	Password  string
	IsEnabled bool
	IsDeleted bool
	GoogleID  *string
	IsAdmin   bool
}

type Confirmation struct {
	Uid    string
	UserId int
}

type EmailPayload struct {
	From     EmailAddress   `json:"from"`
	To       []EmailAddress `json:"to"`
	Subject  string         `json:"subject"`
	HTML     string         `json:"HTML"`
	Category string         `json:"category"`
}

type EmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}
