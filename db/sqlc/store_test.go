package db

import (
	"context"
	"testing"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
	"github.com/stretchr/testify/require"
)

func TestTransferTransactions(t *testing.T) {
	ctx := context.TODO()

	user1 := createRandomMerchants(t)
	wallet1 := createRandomMerchants(t)
	wallet2 := createRandomMerchants(t)

	arg := CreateTransferParams{
		UserID:       int32(user1.ID),
		FromWalletID: int32(wallet1.ID),
		ToWalletID:   int32(wallet2.ID),
		Amount:       float64(utils.RandomMoney()),
		Description:  utils.RandomString(12),
	}

	// Pengujian
	result, err := testStore.TransferTransactions(ctx, arg)
	require.NoError(t, err) //no err = berhasil

	// Assersi
	if result.Transfer.ID == 0 {
		t.Error("Transfer ID should not be zero")
	}
	if result.FromWallet.ID == 0 {
		t.Error("FromWallet ID should not be zero")
	}
	if result.ToWallet.ID == 0 {
		t.Error("ToWallet ID should not be zero")
	}

}
