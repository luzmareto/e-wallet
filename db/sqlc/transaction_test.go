package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransaction(t *testing.T) {
	ctx := context.Background()

	// Create a transaction
	arg := CreateTransactionParams{
		UserID:      1,
		WalletID:    1,
		Amount:      10000,
		Description: "Test transaction",
	}

	err := testQueries.CreateTransaction(ctx, arg)
	require.NoError(t, err)
}

func TestGetTransactionUserID(t *testing.T) {
	ctx := context.Background()

	// Get transactions by user ID
	userID := int32(1)
	transactions, err := testQueries.GetTransactionUserID(ctx, userID)
	require.NoError(t, err)
	require.NotNil(t, transactions)
}

func TestGetTransactionWalletID(t *testing.T) {
	ctx := context.Background()

	// Get transactions by wallet ID
	walletID := int32(1)
	transactions, err := testQueries.GetTransactionWalletID(ctx, walletID)
	require.NoError(t, err)
	require.NotNil(t, transactions)
}
