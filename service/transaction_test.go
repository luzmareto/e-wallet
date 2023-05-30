package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbmocks "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/mocks"
	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

func TestCreateTransaction(t *testing.T) {
	dummyWallet := db.Wallet{ID: 1}
	dummyUser := db.User{ID: 1}
	arg := db.CreateTransactionParams{
		UserID:          int32(dummyUser.ID),
		WalletID:        int32(dummyWallet.ID),
		Amount:          float64(utils.RandomMoney()),
		Description:     utils.RandomString(255),
		TransactionType: utils.RandomTransactionTypes(),
	}
	testCase := []struct {
		name          string
		arg           db.CreateTransactionParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("CreateTransaction", mock.Anything, mock.AnythingOfTypeArgument("db.CreateTransactionParams")).
					Return(nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.CreateTransaction(context.Background(), arg)
				require.NoError(t, err)
			},
		},
		{
			name: "Unexpected Erro",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("CreateTransaction", mock.Anything, mock.AnythingOfTypeArgument("db.CreateTransactionParams")).
					Return(errors.New("unexpected error"))
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.CreateTransaction(context.Background(), arg)
				require.Error(t, err)
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

func TestGetTransactionUserID(t *testing.T) {
	dummyWallet := db.Wallet{ID: 1}
	dummyUser := db.User{ID: 1}
	arg := db.CreateTransactionParams{
		UserID:          int32(dummyUser.ID),
		WalletID:        int32(dummyWallet.ID),
		Amount:          0,
		Description:     utils.RandomString(255),
		TransactionType: utils.RandomTransactionTypes(),
	}
	dummyTransactions := []db.Transaction{
		{
			ID:              1,
			UserID:          int32(arg.UserID),
			WalletID:        int32(arg.WalletID),
			Amount:          arg.Amount,
			TransactionDate: time.Now().String(),
			Description:     arg.Description,
			TransactionType: arg.TransactionType,
		},
		{
			ID:              2,
			UserID:          int32(arg.UserID),
			WalletID:        int32(arg.WalletID),
			Amount:          arg.Amount,
			TransactionDate: time.Now().String(),
			Description:     arg.Description,
			TransactionType: arg.TransactionType,
		},
	}
	testCase := []struct {
		name          string
		arg           int32
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  int32(dummyUser.ID),
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransactionUserID", mock.Anything, mock.AnythingOfType("int32")).
					Return(dummyTransactions, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				transactions, err := svc.GetTransactionUserID(context.Background(), int32(dummyUser.ID))
				require.NoError(t, err)
				require.NotEmpty(t, transactions)
				require.GreaterOrEqual(t, len(transactions), len(dummyTransactions))
			},
		},
		{
			name: "Unexpected Erro",
			arg:  0,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransactionUserID", mock.Anything, mock.AnythingOfType("int32")).
					Return([]db.Transaction{}, errors.New("unexpected error"))
			},
			checkresponse: func(t *testing.T, svc Service) {
				transactions, err := svc.GetTransactionUserID(context.Background(), int32(0))
				require.Error(t, err)
				require.Empty(t, transactions)
				require.Equal(t, len(transactions), 0)
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

func TestGetTransactionWalletID(t *testing.T) {
	dummyWallet := db.Wallet{ID: 1}
	dummyUser := db.User{ID: 1}
	arg := db.CreateTransactionParams{
		UserID:          int32(dummyUser.ID),
		WalletID:        int32(dummyWallet.ID),
		Amount:          0,
		Description:     utils.RandomString(255),
		TransactionType: utils.RandomTransactionTypes(),
	}
	dummyTransactions := []db.Transaction{
		{
			ID:              1,
			UserID:          int32(arg.UserID),
			WalletID:        int32(arg.WalletID),
			Amount:          arg.Amount,
			TransactionDate: time.Now().String(),
			Description:     arg.Description,
			TransactionType: arg.TransactionType,
		},
		{
			ID:              2,
			UserID:          int32(arg.UserID),
			WalletID:        int32(arg.WalletID),
			Amount:          arg.Amount,
			TransactionDate: time.Now().String(),
			Description:     arg.Description,
			TransactionType: arg.TransactionType,
		},
	}
	testCase := []struct {
		name          string
		arg           int32
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  int32(dummyUser.ID),
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransactionWalletID", mock.Anything, mock.AnythingOfType("int32")).
					Return(dummyTransactions, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				transactions, err := svc.GetTransactionWalletID(context.Background(), int32(dummyWallet.ID))
				require.NoError(t, err)
				require.NotEmpty(t, transactions)
				require.GreaterOrEqual(t, len(transactions), len(dummyTransactions))
			},
		},
		{
			name: "Unexpected Erro",
			arg:  0,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransactionWalletID", mock.Anything, mock.AnythingOfType("int32")).
					Return([]db.Transaction{}, errors.New("unexpected error"))
			},
			checkresponse: func(t *testing.T, svc Service) {
				transactions, err := svc.GetTransactionWalletID(context.Background(), int32(0))
				require.Error(t, err)
				require.Empty(t, transactions)
				require.Equal(t, len(transactions), 0)
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

func TestGetTransactionWalletByidAndUserID(t *testing.T) {
	dummyWallet := db.Wallet{ID: 1}
	dummyUser := db.User{ID: 1}
	arg := db.GetTransactionWalletByidAndUserIDParams{
		UserID:   int32(dummyUser.ID),
		WalletID: int32(dummyWallet.ID),
	}
	dummyTransactions := []db.Transaction{
		{
			ID:              1,
			UserID:          int32(arg.UserID),
			WalletID:        int32(arg.WalletID),
			Amount:          float64(utils.RandomMoney()),
			TransactionDate: time.Now().String(),
			Description:     utils.RandomString(255),
			TransactionType: utils.RandomTransactionTypes(),
		},
		{
			ID:              2,
			UserID:          int32(arg.UserID),
			WalletID:        int32(arg.WalletID),
			Amount:          float64(utils.RandomMoney()),
			TransactionDate: time.Now().String(),
			Description:     utils.RandomString(255),
			TransactionType: utils.RandomTransactionTypes(),
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
				mockStore.On("GetTransactionWalletByidAndUserID", mock.Anything, mock.AnythingOfTypeArgument("db.GetTransactionWalletByidAndUserIDParams")).
					Return(dummyTransactions, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				transactions, err := svc.GetTransactionWalletByidAndUserID(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, transactions)
				require.GreaterOrEqual(t, len(transactions), len(dummyTransactions))
			},
		},
		{
			name: "Unexpected Erro",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransactionWalletByidAndUserID", mock.Anything, mock.AnythingOfTypeArgument("db.GetTransactionWalletByidAndUserIDParams")).
					Return([]db.Transaction{}, errors.New("unexpected error"))
			},
			checkresponse: func(t *testing.T, svc Service) {
				transactions, err := svc.GetTransactionWalletByidAndUserID(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, transactions)
				require.Equal(t, len(transactions), 0)
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
