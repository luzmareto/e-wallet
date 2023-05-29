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

var dummyUserTopUp = db.User{
	ID:               2,
	Role:             "user",
	Username:         "user",
	Password:         "user",
	Email:            "user@gmai.com",
	PhoneNumber:      "011111",
	IDCard:           "666.jpg",
	RegistrationDate: time.Now(),
}

var dummyWalletTopUp = db.Wallet{
	ID:       1,
	UserID:   int32(dummyUserTopUp.ID),
	Balance:  0,
	Currency: "IDR",
}

var dummyTopUp = db.Topup{
	ID:        1,
	UserID:    int32(dummyUserTopUp.ID),
	WalletID:  int32(dummyWalletTopUp.ID),
	TopupDate: time.Now(),
}

func TestCreateTopUps(t *testing.T) {
	arg := db.CreateTopUpsParams{
		UserID:      int32(dummyUserTopUp.ID),
		WalletID:    int32(dummyWalletTopUp.ID),
		Amount:      1000000,
		Description: "top up",
	}

	dummyTopUp.Amount = arg.Amount
	dummyTopUp.Description = arg.Description

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
				mockStore.On("CreateTopUps", mock.Anything, mock.Anything).
					Return(dummyTopUp, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				topUp, err := svc.CreateTopUps(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, topUp)
				require.Equal(t, arg.UserID, topUp.UserID)
				require.Equal(t, arg.WalletID, topUp.WalletID)
				require.Equal(t, arg.Description, topUp.Description)
				require.Equal(t, arg.Amount, topUp.Amount)
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("CreateTopUps", mock.Anything, mock.Anything).
					Return(db.Topup{}, errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				topUp, err := svc.CreateTopUps(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, topUp)
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
