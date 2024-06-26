package api

import (
	"database/sql"
	"errors"
	"net/http"
	db "simple_bank/db/sqlc"
	"simple_bank/internal/dto"
	"simple_bank/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func newUserResponse(user db.User) dto.UserResponse {
	return dto.UserResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}
func (server *Server) createUser(c *gin.Context) {
	var req dto.CreateUserRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Username:       req.UserName,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(c, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				c.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, newUserResponse(user))
}

func (server *Server) loginUser(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.GetUser(c, req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := dto.UserLoginResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	c.JSON(http.StatusOK, rsp)
}
