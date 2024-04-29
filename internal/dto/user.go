package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserRequest struct {
	UserName string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UserResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserLoginResponse struct {
	SessionId           uuid.UUID    `json:"session_id"`
	AccessToken         string       `json:"access_token"`
	AccessTokenExpires  time.Time    `json:"access_token_expires"`
	RefreshToken        string       `json:"refresh_token"`
	RefreshTokenExpires time.Time    `json:"refresh_token_expires"`
	User                UserResponse `json:"user"`
}
