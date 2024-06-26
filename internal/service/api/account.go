package api

import (
	"database/sql"
	"errors"
	"net/http"
	db "simple_bank/db/sqlc"
	"simple_bank/internal/dto"
	"simple_bank/internal/token"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (server *Server) createAccount(c *gin.Context) {
	var req dto.CreateAccountRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Balance:  0,
		Currency: req.Currency,
	}

	account, err := server.store.CreateAccount(c, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				c.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, account)
}

func (server *Server) getAccount(c *gin.Context) {
	var req dto.GetAccountRequest
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	account, err := server.store.GetAccount(c, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, account)
}

func (server *Server) listAccount(c *gin.Context) {
	var req dto.ListAccountsRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)

	accounts, err := server.store.ListAccount(c, db.ListAccountParams{
		Owner:  authPayload.Username,
		Limit:  int32(req.PageSize),
		Offset: int32(req.PageId-1) * int32(req.PageSize),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, accounts)
}
