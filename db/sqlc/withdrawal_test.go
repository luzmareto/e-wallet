package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateWithdrawals(t *testing.T) {
	user := createRandomMerchants(t)
	ctx := context.Background()

	// Create a withdrawal
	arg := CreateWithdrawalsParams{
		UserID:      int32(user.ID),
		WalletID:    1,
		Amount:      10000,
		Description: "Test withdrawal",
	}

	data, err := testQueries.CreateWithdrawals(ctx, arg)
	require.NoError(t, err)
	require.NotZero(t, data)
	require.Equal(t, user.ID, int64(data.UserID))
	require.Equal(t, int32(1), data.WalletID)
	require.Equal(t, float64(10000), data.Amount)
	require.NotEmpty(t, data.WithdrawalDate)
	require.Equal(t, "Test withdrawal", data.Description)
}

func TestCreateWithdrawalsNotFound(t *testing.T) {
	ctx := context.Background()

	// Create a withdrawal for a user that does not exist
	arg := CreateWithdrawalsParams{
		UserID:      12345,
		WalletID:    1,
		Amount:      10000,
		Description: "Test withdrawal",
	}

	_, err := testQueries.CreateWithdrawals(ctx, arg)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
