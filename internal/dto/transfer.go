package dto

type TransferRequest struct {
	FromAccount int64  `json:"from_account" binding:"required,min=1"`
	ToAccount   int64  `json:"to_account" binding:"required,min=1"`
	Amount      int64  `json:"amount" binding:"required,gt=0"`
	Currency    string `json:"currency" binding:"required,currency"`
}
