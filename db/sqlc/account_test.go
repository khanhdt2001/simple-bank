package db

import (
	"context"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	account1, err := testQueries.GetAccountForUpdate(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	require.Equal(t, account.ID, account1.ID)
	require.Equal(t, account.Owner, account1.Owner)
	require.Equal(t, account.Balance, account1.Balance)
	require.Equal(t, account.Currency, account1.Currency)
	require.WithinDuration(t, account.CreatedAt, account1.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)

	}
	args := ListAccountParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
		require.Equal(t, lastAccount.Owner, acc.Owner)
	}
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	account1, err := testQueries.UpdateAccount(context.Background(), UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	require.Equal(t, account.ID, account1.ID)
	require.Equal(t, account.Owner, account1.Owner)
	require.NotEqual(t, account.Balance, account1.Balance)
	require.Equal(t, account.Currency, account1.Currency)
	require.WithinDuration(t, account.CreatedAt, account1.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccountForUpdate(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}
