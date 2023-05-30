package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbmocks "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/mocks"
	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

func TestTopupTransactions(t *testing.T) {
	dummyUser := db.User{ID: 1}
	dummyWallet := db.Wallet{
		ID:       1,
		UserID:   int32(dummyUser.ID),
		Balance:  0,
		Currency: utils.RandomCurrency(),
	}
	arg := db.CreateTopUpsParams{
		UserID:      int32(dummyUser.ID),
		WalletID:    int32(dummyWallet.ID),
		Amount:      float64(utils.RandomMoney()),
		Description: utils.RandomString(200),
	}
	topUpsresult := db.TopupResult{
		Topup: db.Topup{
			ID:          1,
			UserID:      int32(dummyUser.ID),
			WalletID:    int32(dummyWallet.ID),
			Amount:      arg.Amount,
			TopupDate:   time.Now(),
			Description: arg.Description,
		},
		Wallet: dummyWallet,
	}
	testCase := []struct {
		name          string
		arg           db.CreateTopUpsParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(dummyWallet, nil)
				mockStore.On("TopupTransactions", mock.Anything, mock.AnythingOfTypeArgument("CreateTopUpsParams")).
					Return(topUpsresult, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.TopupTransactions(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, result)
			},
		},
		{
			name: "Not Found",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.Wallet{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.TopupTransactions(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, result)
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.Wallet{}, errors.New("unexpected error"))
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.TopupTransactions(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, result)
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}

			svc := New(nil)
			svc.SetStore(mockStore)

			tc.buildStubs(mockStore)
			tc.checkresponse(t, svc)

			mockStore.AssertExpectations(t)
		})
	}
}

func TestWithdrawalTransactions(t *testing.T) {
	dummyUser := db.User{ID: 1}
	dummyWallet := db.Wallet{
		ID:       1,
		UserID:   int32(dummyUser.ID),
		Balance:  10000,
		Currency: utils.RandomCurrency(),
	}
	arg := db.CreateWithdrawalsParams{
		UserID:      int32(dummyUser.ID),
		WalletID:    int32(dummyWallet.ID),
		Amount:      5000,
		Description: utils.RandomString(100),
	}
	wdResult := db.WithdrawalResult{
		Withdrawal: db.Withdrawal{
			ID:             1,
			UserID:         arg.UserID,
			WalletID:       arg.WalletID,
			Amount:         arg.Amount,
			WithdrawalDate: time.Now(),
			Description:    arg.Description,
		},
		Wallet: dummyWallet,
	}

	testCase := []struct {
		name          string
		arg           db.CreateWithdrawalsParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(dummyWallet, nil)
				mockStore.On("GetUserById", mock.Anything, mock.AnythingOfType("int64")).
					Return(dummyUser, nil)
				mockStore.On("WithdrawalTransactions", mock.Anything, mock.AnythingOfTypeArgument("db.CreateWithdrawalsParams")).
					Return(wdResult, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.WithdrawalTransactions(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, result)
			},
		},
		{
			name: "wallet not found",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.Wallet{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.WithdrawalTransactions(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, result)
				require.EqualError(t, err, fmt.Sprintf("wallet with id %d not found", arg.WalletID))
			},
		},
		{
			name: "user not found",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(dummyWallet, nil)
				mockStore.On("GetUserById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.WithdrawalTransactions(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, result)
				require.EqualError(t, err, fmt.Sprintf("user with id %d not found", arg.WalletID))
			},
		},
		{
			name: "unauthorized error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(dummyWallet, nil)
				mockStore.On("GetUserById", mock.Anything, mock.AnythingOfType("int64")).
					Return(dummyUser, nil)
				mockStore.On("WithdrawalTransactions", mock.Anything, mock.AnythingOfTypeArgument("db.CreateWithdrawalsParams")).
					Return(db.WithdrawalResult{}, fmt.Errorf("wallet with id %d does not belong to user with id %d", dummyWallet.ID, dummyUser.ID))
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.WithdrawalTransactions(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, result)
				require.EqualError(t, err, fmt.Sprintf("wallet with id %d does not belong to user with id %d", dummyWallet.ID, dummyUser.ID))
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}

			svc := New(nil)
			svc.SetStore(mockStore)

			tc.buildStubs(mockStore)
			tc.checkresponse(t, svc)

			mockStore.AssertExpectations(t)
		})
	}
}

func TestTransferTransactions(t *testing.T) {
	user1 := db.User{ID: 1}
	user2 := db.User{ID: 1}
	wallet1 := db.Wallet{
		ID:       1,
		UserID:   int32(user1.ID),
		Balance:  500000,
		Currency: "IDR",
	}
	wallet2 := db.Wallet{
		ID:       2,
		UserID:   int32(user2.ID),
		Balance:  10000,
		Currency: "IDR",
	}
	arg := db.CreateTransferParams{
		UserID:       int32(user1.ID),
		FromWalletID: int32(wallet1.ID),
		ToWalletID:   int32(wallet2.ID),
		Amount:       100000,
		Description:  utils.RandomString(100),
	}
	tfResult := db.TransferResult{
		Transfer: db.Transfer{
			ID:           1,
			FromWalletID: arg.FromWalletID,
			ToWalletID:   arg.ToWalletID,
			Amount:       arg.Amount,
			TransferDate: time.Now(),
			Description:  arg.Description,
		},
		FromWallet: wallet1,
		ToWallet:   wallet2,
	}
	testCase := []struct {
		name          string
		arg           db.CreateTransferParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(wallet1, nil)
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(wallet2, nil)
				mockStore.On("TransferTransactions", mock.Anything, mock.AnythingOfType("db.CreateTransferParams")).
					Return(tfResult, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.TransferTransactions(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, result)
			},
		},
		{
			name: "unexpected-error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(wallet1, nil)
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(wallet2, nil)
				mockStore.On("TransferTransactions", mock.Anything, mock.AnythingOfType("db.CreateTransferParams")).
					Return(db.TransferResult{}, errors.New("unexpected error"))
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.TransferTransactions(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, result)
			},
		},
		{
			name: "wallet_from_not_found",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.Wallet{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.TransferTransactions(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, result)
				require.EqualError(t, err, fmt.Sprintf("wallet with id %d not found", arg.FromWalletID))
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}

			svc := New(nil)
			svc.SetStore(mockStore)

			tc.buildStubs(mockStore)
			tc.checkresponse(t, svc)

			mockStore.AssertExpectations(t)
		})
	}
}

func TestMerchantPaymentTransactions(t *testing.T) {
	user1 := db.User{ID: 1}
	wallet1 := db.Wallet{
		ID:      1,
		UserID:  int32(user1.ID),
		Balance: 500000,
	}
	arg1 := db.CreateTransactionParams{
		UserID:          int32(user1.ID),
		WalletID:        int32(wallet1.ID),
		Amount:          1000,
		Description:     utils.RandomString(200),
		TransactionType: "PAYMENT",
	}
	merchant := db.Merchant{ID: 1}
	arg2 := merchant.ID

	testCase := []struct {
		name          string
		arg1          db.CreateTransactionParams
		arg2          int64
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg1: arg1,
			arg2: arg2,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(wallet1, nil)
				mockStore.On("MerchantPaymentTransactions", mock.Anything, mock.AnythingOfTypeArgument("db.CreateTransactionParams"), mock.AnythingOfType("int64")).
					Return(nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.MerchantPaymentTransactions(context.Background(), arg1, arg2)
				require.NoError(t, err)
			},
		},
		{
			name: "wallet_not_found",
			arg1: arg1,
			arg2: arg2,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.Wallet{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.MerchantPaymentTransactions(context.Background(), arg1, arg2)
				require.Error(t, err)
				require.EqualError(t, err, fmt.Sprintf("wallet with id %d not found", arg1.WalletID))
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}

			svc := New(nil)
			svc.SetStore(mockStore)

			tc.buildStubs(mockStore)
			tc.checkresponse(t, svc)

			mockStore.AssertExpectations(t)
		})
	}
}

func TestWalletHistoryGenerateCSV(t *testing.T) {
	arg := db.GetTransactionWalletByidAndUserIDParams{
		WalletID: 1,
		UserID:   1,
	}

	csvresult := db.WalletHistoryResult{
		Transactions: []db.Transaction{
			{
				ID:              1,
				UserID:          1,
				WalletID:        1,
				Amount:          1000,
				TransactionDate: time.Now().String(),
				TransactionType: "TRANSFER",
				Description:     utils.RandomString(100),
			},
		},
		Transfers: []db.Transfer{
			{
				ID:           1,
				FromWalletID: 1,
				ToWalletID:   2,
				Amount:       1000,
				TransferDate: time.Now(),
				Description:  utils.RandomString(100),
			},
		},
	}
	testCase := []struct {
		name          string
		arg           db.GetTransactionWalletByidAndUserIDParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("WalletHistoryGenerateCSV", mock.Anything, mock.AnythingOfTypeArgument("db.GetTransactionWalletByidAndUserIDParams")).
					Return(db.WalletHistoryResult{}, errors.New("unexpected error"))
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.WalletHistoryGenerateCSV(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, result)
			},
		},
		{
			name: "unexpected_error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("WalletHistoryGenerateCSV", mock.Anything, mock.AnythingOfTypeArgument("db.GetTransactionWalletByidAndUserIDParams")).
					Return(csvresult, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.WalletHistoryGenerateCSV(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, result)
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}

			svc := New(nil)
			svc.SetStore(mockStore)

			tc.buildStubs(mockStore)
			tc.checkresponse(t, svc)

			mockStore.AssertExpectations(t)
		})
	}
}
