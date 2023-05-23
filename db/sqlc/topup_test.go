package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateTopUps(t *testing.T) {
	ctx := context.Background()
	user := createRandomUser(t)
	wallet := createRandomWallets(t)
	amount := float64(10000)
	description := "Top up for my account"

	// Act
	arg := CreateTopUpsParams{
		UserID:      int32(user.ID),
		WalletID:    int32(wallet.ID),
		Amount:      amount,
		Description: description,
	}
	topup, err := testQueries.CreateTopUps(ctx, arg)
	require.NoError(t, err)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, topup)
	assert.Equal(t, int32(user.ID), topup.UserID)
	assert.Equal(t, int32(wallet.ID), topup.WalletID)
	assert.Equal(t, amount, topup.Amount)
	assert.Equal(t, description, topup.Description)
}
