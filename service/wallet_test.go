package service

import (
	"context"
	"errors"
	"testing"
	"time"

	dbmocks "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/mocks"
	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var dummyUserWallet = db.User{
	ID:               2,
	Role:             "user",
	Username:         "user",
	Password:         "user",
	Email:            "user@gmai.com",
	PhoneNumber:      "011111",
	IDCard:           "666.jpg",
	RegistrationDate: time.Now(),
}

var dummyWallet = db.Wallet{
	ID:       1,
	UserID:   int32(dummyUserWallet.ID),
	Balance:  0,
	Currency: "",
}

func TestCreateWallets(t *testing.T) {
	arg := db.CreateWalletsParams{
		UserID:   int32(dummyUserWallet.ID),
		Balance:  0,
		Currency: "IDR",
	}

	testCase := []struct {
		name          string
		arg           db.CreateWalletsParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkResponse func(t *testing.T, svc Service, wallet db.Wallet, err error)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, int64(arg.UserID)).Return(dummyUserWallet, nil)
				mockStore.On("CreateWallets", mock.Anything, mock.Anything).
					Run(func(args mock.Arguments) {
						createWalletsParams := args.Get(1).(db.CreateWalletsParams)
						require.Equal(t, CurrencyIDR, createWalletsParams.Currency) // Assert the Currency value
					}).
					Return(db.Wallet{
						UserID:   arg.UserID,
						Balance:  arg.Balance,
						Currency: arg.Currency,
						// Add any additional fields you need for testing
					}, nil)
			},
			checkResponse: func(t *testing.T, svc Service, wallet db.Wallet, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, wallet)
				require.Equal(t, arg.UserID, wallet.UserID)
				require.Equal(t, arg.Balance, wallet.Balance)
				require.Equal(t, arg.Currency, wallet.Currency)
			},
		},
		{
			name: "User Not Found",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, int64(arg.UserID)).Return(db.User{}, errors.New("user not found"))
			},
			checkResponse: func(t *testing.T, svc Service, wallet db.Wallet, err error) {
				require.Error(t, err)
				require.Empty(t, wallet)
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, int64(arg.UserID)).Return(dummyUserWallet, nil)
				mockStore.On("CreateWallets", mock.Anything, mock.Anything).Return(db.Wallet{}, errors.New("unexpected error"))
			},
			checkResponse: func(t *testing.T, svc Service, wallet db.Wallet, err error) {
				require.Error(t, err)
				require.Empty(t, wallet)
				// Add any additional assertions you need for testing
			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}
			svc := New(nil)
			svc.SetStore(mockStore)
			tc.buildStubs(mockStore)
			wallet, err := svc.CreateWallets(context.Background(), tc.arg)
			tc.checkResponse(t, svc, wallet, err)
			mockStore.AssertExpectations(t)
		})
	}

}

// var dummyUser1 = []db.User{
// 	{
// 		ID:               1,
// 		Role:             "admin",
// 		Username:         "admin",
// 		Password:         "password-admin",
// 		Email:            "admin@mail.com",
// 		PhoneNumber:      "012345678",
// 		IDCard:           "jhkhkj.jpg",
// 		RegistrationDate: time.Now(),
// 	},
// 	{
// 		ID:               2,
// 		Role:             "user",
// 		Username:         "user",
// 		Password:         "password-user",
// 		Email:            "user@mail.com",
// 		PhoneNumber:      "012345678",
// 		IDCard:           "jhkhkj.jpg",
// 		RegistrationDate: time.Now(),
// 	},
// }

func TestGetWalletById(t *testing.T) {
	expectedWallet := db.Wallet{
		ID:       dummyWallet.ID,
		UserID:   int32(dummyUserWallet.ID),
		Balance:  1000,
		Currency: "IDR",
	}

	testCase := []struct {
		name          string
		id            int64
		buildStubs    func(mockStore *dbmocks.Store)
		checkResponse func(t *testing.T, svc Service, wallet db.Wallet, err error)
	}{
		{
			name: "OK",
			id:   dummyWallet.ID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, dummyWallet.ID).Return(expectedWallet, nil)
			},
			checkResponse: func(t *testing.T, svc Service, wallet db.Wallet, err error) {
				require.NoError(t, err)
				require.Equal(t, expectedWallet, wallet)
			},
		},
		{
			name: "Wallet Not Found",
			id:   999, // Non-existing wallet ID
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, int64(999)).Return(db.Wallet{}, errors.New("not found"))
			},
			checkResponse: func(t *testing.T, svc Service, wallet db.Wallet, err error) {
				require.Error(t, err)
				require.Empty(t, wallet)
			},
		},
		{
			name: "Unexpected Error",
			id:   dummyWallet.ID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetWalletById", mock.Anything, dummyWallet.ID).Return(db.Wallet{}, errors.New("unexpected error"))
			},
			checkResponse: func(t *testing.T, svc Service, wallet db.Wallet, err error) {
				require.Error(t, err)
				require.Empty(t, wallet)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}
			svc := New(nil)
			svc.SetStore(mockStore)
			tc.buildStubs(mockStore)
			wallet, err := svc.GetWalletById(context.Background(), tc.id)
			tc.checkResponse(t, svc, wallet, err)
			mockStore.AssertExpectations(t)
		})
	}
}

func TestAddWalletBalance(t *testing.T) {
	arg := db.AddWalletBalanceParams{
		ID:      dummyUserWallet.ID,
		Balance: 0,
	}

	expectedWallet := db.Wallet{
		ID:       arg.ID,
		UserID:   2, // Assuming the user ID
		Balance:  500,
		Currency: "IDR",
	}

	testCase := []struct {
		name          string
		arg           db.AddWalletBalanceParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkResponse func(t *testing.T, svc Service, wallet db.Wallet, err error)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("AddWalletBalance", mock.Anything, arg).Return(expectedWallet, nil)
			},
			checkResponse: func(t *testing.T, svc Service, wallet db.Wallet, err error) {
				require.NoError(t, err)
				require.Equal(t, expectedWallet, wallet)
			},
		},
		{
			name: "Wallet Not Found",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("AddWalletBalance", mock.Anything, arg).Return(db.Wallet{}, errors.New("wallet not found"))
			},
			checkResponse: func(t *testing.T, svc Service, wallet db.Wallet, err error) {
				require.Error(t, err)
				require.Empty(t, wallet)
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("AddWalletBalance", mock.Anything, arg).Return(db.Wallet{}, errors.New("unexpected error"))
			},
			checkResponse: func(t *testing.T, svc Service, wallet db.Wallet, err error) {
				require.Error(t, err)
				require.Empty(t, wallet)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}
			svc := New(nil)
			svc.SetStore(mockStore)
			tc.buildStubs(mockStore)
			wallet, err := svc.AddWalletBalance(context.Background(), tc.arg)
			tc.checkResponse(t, svc, wallet, err)
			mockStore.AssertExpectations(t)
		})
	}
}
