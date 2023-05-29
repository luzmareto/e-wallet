package service

import (
	"context"
	"errors"
	"testing"
	"time"

	dbmocks "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/mocks"
	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var dummyUserWD = db.User{
	ID:               1,
	Role:             "user",
	Username:         "user",
	Password:         "user",
	Email:            "user@gmail.com",
	PhoneNumber:      "0123",
	IDCard:           "user.jgp",
	RegistrationDate: time.Now(),
}

var dummyWalletWD = db.Wallet{
	ID:       1,
	UserID:   int32(dummyUserWD.ID),
	Balance:  0,
	Currency: "IDR",
}

var dummyWd = db.Withdrawal{
	ID:             1,
	UserID:         int32(dummyUserWD.ID),
	WalletID:       int32(dummyWalletWD.ID),
	WithdrawalDate: time.Now(),
}

func TestCreateWithdrawals(t *testing.T) {
	arg := db.CreateWithdrawalsParams{
		UserID:      int32(dummyUserWD.ID),
		WalletID:    int32(dummyWalletWD.ID),
		Amount:      1000000,
		Description: "WD = Withdrawal",
	}

	dummyWd.Amount = arg.Amount
	dummyWd.Description = arg.Description

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
				mockStore.On("GetUserById", mock.Anything, mock.Anything).
					Return(dummyUserWD, nil)
				mockStore.On("CreateWithdrawals", mock.Anything, mock.Anything).
					Return(dummyWd, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				withdrawal, err := svc.CreateWithdrawals(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, withdrawal)
				require.Equal(t, arg.UserID, withdrawal.UserID)
				require.Equal(t, arg.WalletID, withdrawal.WalletID)
				require.Equal(t, arg.Description, withdrawal.Description)
				require.Equal(t, arg.Amount, withdrawal.Amount)
			},
		},
		{
			name: "User Not Found",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.Anything).
					Return(db.User{}, errors.New("user not found"))
			},
			checkresponse: func(t *testing.T, svc Service) {
				withdrawal, err := svc.CreateWithdrawals(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, withdrawal)
				require.IsType(t, &utils.CustomError{}, err)

				customErr, ok := err.(*utils.CustomError)
				require.True(t, ok)
				require.Equal(t, "user with id 1 not found", customErr.Msg)
				require.Equal(t, "user not found", customErr.Err.Error())
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.Anything).
					Return(dummyUserWD, nil)
				mockStore.On("CreateWithdrawals", mock.Anything, mock.Anything).
					Return(db.Withdrawal{}, errors.New("unexpected error"))
			},
			checkresponse: func(t *testing.T, svc Service) {
				withdrawal, err := svc.CreateWithdrawals(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, withdrawal)
				require.EqualError(t, err, "unexpected error")
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
