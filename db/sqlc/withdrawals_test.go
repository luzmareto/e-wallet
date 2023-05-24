package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateWithdrawals(t *testing.T) {
	ctx := context.Background()

	arg := CreateWithdrawalsParams{
		UserID:      1,
		WalletID:    1,
		Amount:      100.0,
		Description: "Test withdrawal",
	}

	withdrawal, err := testQueries.CreateWithdrawals(ctx, arg)

	assert.NoError(t, err)

	assert.NotZero(t, withdrawal.ID)
	assert.Equal(t, arg.UserID, withdrawal.UserID)
	assert.Equal(t, arg.WalletID, withdrawal.WalletID)
	assert.Equal(t, arg.Amount, withdrawal.Amount)
	assert.Equal(t, arg.Description, withdrawal.Description)

}
