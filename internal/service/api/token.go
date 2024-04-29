package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"simple_bank/internal/dto"

	"github.com/gin-gonic/gin"
)

func (server *Server) RenewAccessToken(c *gin.Context) {
	var req dto.RenewAccessTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	session, err := server.store.GetSessionById(c, payload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if session.IsBlocked {
		err := fmt.Errorf("session is blocked")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	if session.Username == payload.Username {
		err := fmt.Errorf("refresh token is for another user")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	if session.RefreshToken == req.RefreshToken {
		err := fmt.Errorf("refresh token is revoked")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return

	}
	accessToken, accessTokenPayload, err := server.tokenMaker.CreateToken(
		payload.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, dto.RenewAccessTokenResponse{
		AccessToken:        accessToken,
		AccessTokenExpires: accessTokenPayload.ExpiredAt,
	})
}
