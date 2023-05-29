package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbmocks "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/mocks"
	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
)

func TestCreateWallets(t *testing.T) {
	user := db.User{
		ID: 1,
	}
	dummywallet := db.Wallet{
		ID:       1,
		UserID:   int32(user.ID),
		Balance:  0,
		Currency: "IDR",
	}
	arg := db.CreateWalletsParams{
		UserID: int32(user.ID),
	}
	testCase := []struct {
		name          string
		arg           db.CreateWalletsParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.AnythingOfType("int64")).
					Return(user, nil)
				mockStore.On("CreateWallets", mock.Anything, mock.Anything).
					Return(dummywallet, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				wallet, err := svc.CreateWallets(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, wallet)
				require.Equal(t, arg.UserID, wallet.UserID)
				require.Equal(t, dummywallet.Balance, wallet.Balance)
				require.Equal(t, dummywallet.Currency, wallet.Currency)
			},
		},
		{
			name: "Not Found",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				wallet, err := svc.CreateWallets(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, wallet)
				require.EqualError(t, err, fmt.Sprintf("user with id %d not found", arg.UserID))
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.AnythingOfType("int64")).
					Return(user, nil)
				mockStore.On("CreateWallets", mock.Anything, mock.Anything).
					Return(db.Wallet{}, errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				wallet, err := svc.CreateWallets(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, wallet)
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

func TestGetWalletByID(t *testing.T) {
	user := db.User{
		ID: 1,
	}
	dummywallet := db.Wallet{
		ID:       1,
		UserID:   int32(user.ID),
		Balance:  0,
		Currency: "IDR",
	}

	testCase := []struct {
		name          string
		arg           int64
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  dummywallet.ID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(dummywallet, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				wallet, err := svc.GetWalletById(context.Background(), dummywallet.ID)
				require.NoError(t, err)
				require.NotEmpty(t, wallet)
				require.Equal(t, dummywallet.UserID, wallet.UserID)
				require.Equal(t, dummywallet.Balance, wallet.Balance)
				require.Equal(t, dummywallet.Currency, wallet.Currency)
			},
		},
		{
			name: "Unexpected Error",
			arg:  dummywallet.ID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.Wallet{}, errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				wallet, err := svc.GetWalletById(context.Background(), dummywallet.ID)
				require.Error(t, err)
				require.Empty(t, wallet)
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

func TestAddWalletBalance(t *testing.T) {
	user := db.User{
		ID: 1,
	}
	dummywallet := db.Wallet{
		ID:       1,
		UserID:   int32(user.ID),
		Balance:  0,
		Currency: "IDR",
	}

	arg := db.AddWalletBalanceParams{
		ID:      dummywallet.ID,
		Balance: 1000,
	}

	testCase := []struct {
		name          string
		arg           db.AddWalletBalanceParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("AddWalletBalance", mock.Anything, mock.Anything).
					Return(dummywallet, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				wallet, err := svc.AddWalletBalance(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, wallet)
				require.Equal(t, dummywallet.UserID, wallet.UserID)
				require.Equal(t, dummywallet.Balance, wallet.Balance)
				require.Equal(t, dummywallet.Currency, wallet.Currency)
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("AddWalletBalance", mock.Anything, mock.Anything).
					Return(db.Wallet{}, errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				wallet, err := svc.AddWalletBalance(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, wallet)
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
