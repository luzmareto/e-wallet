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

func TestMerchantPaymentTransactions(t *testing.T) {
	user := createRandomUser(t)
	wallet := createRandomWallets(t)
	merchant := createRandomMerchants(t)

	arg := CreateTransactionParams{
		UserID:          int32(user.ID),
		WalletID:        int32(wallet.ID),
		Amount:          10000,
		Description:     utils.RandomString(100),
		TransactionType: "PAYMENT",
	}

	err := testStore.MerchantPaymentTransactions(context.Background(), arg, merchant.ID)
	require.NoError(t, err)

}

func TestTopupTransactions(t *testing.T) {
	user := createRandomUser(t)
	wallet := createRandomWallets(t)
	arg := CreateTopUpsParams{
		UserID:      int32(user.ID),
		WalletID:    int32(wallet.ID),
		Amount:      10000,
		Description: utils.RandomString(100),
	}
	result, err := testStore.TopupTransactions(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.NotEmpty(t, result.Topup)
	require.NotEmpty(t, result.Wallet)
}

func TestWithdrawalTransactions(t *testing.T) {
	user := createRandomUser(t)
	wallet := createRandomWallets(t)
	arg := CreateWithdrawalsParams{
		UserID:      int32(user.ID),
		WalletID:    int32(wallet.ID),
		Amount:      10000,
		Description: utils.RandomString(100),
	}

	result, err := testStore.WithdrawalTransactions(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.NotEmpty(t, result.Wallet)
	require.NotEmpty(t, result.Withdrawal)
}

func TestWalletHistoryGenerateCSV(t *testing.T) {
	user := createRandomUser(t)
	wallet := createRandomWallets(t)
	arg := GetTransactionWalletByidAndUserIDParams{
		WalletID: int32(wallet.ID),
		UserID:   int32(user.ID),
	}

	result, err := testStore.WalletHistoryGenerateCSV(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}
