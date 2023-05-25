package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfers(t *testing.T) {
	ctx := context.Background()

	// Create a transfer
	arg := CreateTransfersParams{
		FromWalletID: 1,
		ToWalletID:   2,
		Amount:       10000,
		Description:  "Test transfer",
	}

	transfer, err := testQueries.CreateTransfers(ctx, arg)
	require.NoError(t, err)
	require.NotZero(t, transfer)
	require.Equal(t, arg.FromWalletID, transfer.FromWalletID)
	require.Equal(t, arg.ToWalletID, transfer.ToWalletID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.Equal(t, arg.Description, transfer.Description)

}

func TestGetTransfersByFromWalletID(t *testing.T) {
	ctx := context.Background()
	fromWalletID := int32(1)

	transfers, err := testQueries.GetTransfersByFromWalletID(ctx, fromWalletID)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	for _, transfer := range transfers {
		require.Equal(t, fromWalletID, transfer.FromWalletID)
	}

}

func TestGetTransfersByFromWalletIdAndToWalletId(t *testing.T) {
	ctx := context.Background()
	fromWalletID := int32(1)
	toWalletID := int32(2)

	arg := GetTransfersByFromWalletIdAndToWalletIdParams{
		FromWalletID: fromWalletID,
		ToWalletID:   toWalletID,
	}

	transfers, err := testQueries.GetTransfersByFromWalletIdAndToWalletId(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	for _, transfer := range transfers {
		require.Equal(t, fromWalletID, transfer.FromWalletID)
		require.Equal(t, toWalletID, transfer.ToWalletID)
	}

}
