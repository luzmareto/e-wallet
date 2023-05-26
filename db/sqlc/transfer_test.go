package db

import (
	"context"
	"testing"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfers(t *testing.T) {
	ctx := context.Background()

	wallet1 := createRandomWallets(t)
	wallet2 := createRandomWallets(t)
	// Create a transfer
	arg := CreateTransfersParams{
		FromWalletID: int32(wallet1.ID),
		ToWalletID:   int32(wallet2.ID),
		Amount:       float64(utils.RandomInt(10, 100)),
		Description:  utils.RandomString(255),
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
	// ctx := context.Background()
	// fromWalletID := int32(1)

	// transfers, err := testQueries.GetTransfersByFromWalletID(ctx, fromWalletID)
	// require.NoError(t, err)
	// require.NotEmpty(t, transfers)
	// for _, transfer := range transfers {
	// 	require.Equal(t, fromWalletID, transfer.FromWalletID)
	// }

	ctx := context.Background()

	// Get transactions by user ID
	userID := int32(1)
	transactions, err := testQueries.GetTransfersByFromWalletID(ctx, userID)
	require.NoError(t, err)
	require.NotNil(t, transactions)

}

func TestGetTransfersByFromWalletIdAndToWalletId(t *testing.T) {
	ctx := context.Background()

	// Get transactions by user ID
	// userID := createRandomWallets(t)
	transactions, err := testQueries.GetTransfersByFromWalletIdAndToWalletId(ctx, GetTransfersByFromWalletIdAndToWalletIdParams{})
	require.NoError(t, err)
	require.NotNil(t, transactions)
}
