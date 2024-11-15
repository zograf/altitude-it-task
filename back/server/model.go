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

type UpdatePasswordDTO struct {
	OldPassword    string `json:"old_password" validate:"required"`
	NewPassword    string `json:"new_password" validate:"required"`
	RepeatPassword string `json:"repeat_password" validate:"required,eqfield=NewPassword"`
}

type User struct {
	ID           int
	Name         string
	LastName     string
	Birthday     time.Time
	Email        string
	Password     string
	IsEnabled    bool
	IsDeleted    bool
	GoogleID     *string
	IsAdmin      bool
	Is2FAEnabled bool
	TotpSecret   string
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

type UserInfo struct {
	Name         string `json:"name,omitempty" form:"name"`
	LastName     string `json:"last_name,omitempty" form:"last_name"`
	Email        string `json:"email,omitempty" form:"email"`
	Birthday     string `json:"birthday,omitempty" form:"birthday"`
	Is2FAEnabled bool   `json:"is_2fa_enabled,omitempty" form:"is_2fa_enabled"`
}

type TotpDTO struct {
	Email string `json:"email"`
	Totp  string `json:"totp_code"`
}
