package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

func createRandomWallets(t *testing.T) Wallet {
	ctx := context.Background()
	user := createRandomUser(t)

	balance := float64(utils.RandomMoney())

	arg := CreateWalletsParams{
		UserID:   int32(user.ID),
		Balance:  balance,
		Currency: utils.RandomCurrency(),
	}
	wallet, err := testQueries.CreateWallets(ctx, arg)

	require.NoError(t, err)
	assert.NotZero(t, wallet.ID)
	assert.Equal(t, arg.UserID, wallet.UserID)
	assert.Equal(t, arg.Balance, wallet.Balance)
	assert.Equal(t, arg.Currency, wallet.Currency)

	return wallet
}

func TestCreateWallets(t *testing.T) {
	createRandomWallets(t)
}

func TestAddWalletBalance(t *testing.T) {
	ctx := context.Background()

	// Create a user
	user := createRandomUser(t)

	balance := float64(utils.RandomMoney())

	// Create a wallet for the user
	wallet, err := testQueries.CreateWallets(ctx, CreateWalletsParams{
		UserID:   int32(user.ID),
		Balance:  balance,
		Currency: utils.RandomCurrency(),
	})
	require.NoError(t, err)

	// Add balance to the wallet
	arg := AddWalletBalanceParams{
		ID:      wallet.ID,
		Balance: balance,
	}
	updatedWallet, err := testQueries.AddWalletBalance(ctx, arg)
	require.NoError(t, err)

	// Assert that the balance has been updated
	assert.Equal(t, updatedWallet.Balance, wallet.Balance+arg.Balance)
}

func TestGetWalletById(t *testing.T) {
	ctx := context.Background()

	// Create a user
	user := createRandomUser(t)

	balance := float64(utils.RandomMoney())

	// Create a wallet for the user
	wallet, err := testQueries.CreateWallets(ctx, CreateWalletsParams{
		UserID:   int32(user.ID),
		Balance:  balance,
		Currency: utils.RandomCurrency(),
	})
	require.NoError(t, err)

	// Get the wallet by ID
	gotWallet, err := testQueries.GetWalletById(ctx, wallet.ID)
	require.NoError(t, err)

	// Assert that the wallet is the same as the one created
	assert.Equal(t, gotWallet.ID, wallet.ID)
	assert.Equal(t, gotWallet.UserID, wallet.UserID)
	assert.Equal(t, gotWallet.Balance, wallet.Balance)
	assert.Equal(t, gotWallet.Currency, wallet.Currency)
}
