package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

func TestTransferTransactions(t *testing.T) {
	ctx := context.TODO()

	user1 := createRandomUser(t)
	wallet1 := createRandomWallets(t)
	wallet2 := createRandomWallets(t)

	arg := CreateTransferParams{
		UserID:       int32(user1.ID),
		FromWalletID: int32(wallet1.ID),
		ToWalletID:   int32(wallet2.ID),
		Amount:       float64(utils.RandomInt(10, 100)),
		Description:  utils.RandomString(12),
	}

	// Pengujian
	result, err := testStore.TransferTransactions(ctx, arg)
	require.NoError(t, err) //no err = berhasil
	require.NotEmpty(t, result)
	require.Equal(t, int64(arg.FromWalletID), result.FromWallet.ID)
	require.Equal(t, int64(arg.ToWalletID), result.ToWallet.ID)

}
