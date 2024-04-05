package api

import (
	"database/sql"
	"fmt"
	"net/http"
	db "simple_bank/db/sqlc"
	"simple_bank/internal/dto"
	"simple_bank/internal/token"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (server *Server) createTransfer(c *gin.Context) {
	var req dto.TransferRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fromAccount, valid := server.validAccount(c, req.FromAccount, req.Currency)
	if !valid {
		return
	}

	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromAccount.Owner != authPayload.Username {
		err := errors.New("from account doesn't belong to the authenticated user")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	_, valid = server.validAccount(c, req.ToAccount, req.Currency)

	if !valid {
		return
	}

	arg := db.TransferTxParam{
		FromAccountId: req.FromAccount,
		ToAccountId:   req.ToAccount,
		Amount:        req.Amount,
	}

	account, err := server.store.TransferTx(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, account)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) (db.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}

	return account, true
}
