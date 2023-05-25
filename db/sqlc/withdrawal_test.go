package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

func TestCreateWithdrawals(t *testing.T) {
	wallet := createRandomWallets(t)
	ctx := context.Background()

	// Create a withdrawal
	arg := CreateWithdrawalsParams{
		UserID:      int32(wallet.UserID),
		WalletID:    int32(wallet.ID),
		Amount:      float64(utils.RandomMoney()),
		Description: utils.RandomString(255),
	}

	data, err := testQueries.CreateWithdrawals(ctx, arg)
	require.NoError(t, err)
	require.NotZero(t, data)
	require.Equal(t, wallet.UserID, data.UserID)
	require.Equal(t, arg.WalletID, data.WalletID)
	require.Equal(t, arg.Amount, data.Amount)
	require.NotEmpty(t, data.WithdrawalDate)
	require.Equal(t, arg.Description, data.Description)
}

func TestCreateWithdrawalsFailed(t *testing.T) {
	wallet := createRandomWallets(t)
	ctx := context.Background()

	// Create a withdrawal
	arg := CreateWithdrawalsParams{
		UserID:      0,
		WalletID:    int32(wallet.ID),
		Amount:      float64(utils.RandomMoney()),
		Description: utils.RandomString(255),
	}

	_, err := testQueries.CreateWithdrawals(ctx, arg)
	require.Error(t, err)
	require.EqualError(t, err, "pq: insert or update on table \"withdrawals\" violates foreign key constraint \"withdrawals_user_id_fkey\"")
}
