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

var dummyUserTf = db.User{
	ID:               1,
	Role:             "user",
	Username:         "user",
	Password:         "user",
	Email:            "user@gmail.com",
	PhoneNumber:      "1111",
	IDCard:           "user.jpg",
	RegistrationDate: time.Now(),
}

var dummyTransfer = db.Transfer{
	ID:           dummyUserTf.ID,
	FromWalletID: 2,
	ToWalletID:   3,
	Amount:       500000,
	TransferDate: time.Now(),
	Description:  "TF",
}

func TestCreateTransfers(t *testing.T) {
	arg := db.CreateTransfersParams{
		FromWalletID: dummyTransfer.FromWalletID,
		ToWalletID:   dummyTransfer.ToWalletID,
		Amount:       500000,
		Description:  "TF",
	}

	testCase := []struct {
		name          string
		arg           db.CreateTransfersParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkResponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("CreateTransfers", mock.Anything, arg).
					Return(dummyTransfer, nil)
			},
			checkResponse: func(t *testing.T, svc Service) {
				tf, err := svc.CreateTransfers(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, tf)
				require.Equal(t, arg.FromWalletID, tf.FromWalletID)
				require.Equal(t, arg.ToWalletID, tf.ToWalletID)
				require.Equal(t, arg.Amount, tf.Amount)
				require.Equal(t, arg.Description, tf.Description)
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("CreateTransfers", mock.Anything, arg).
					Return(db.Transfer{}, errors.New("unexpected error"))
			},
			checkResponse: func(t *testing.T, svc Service) {
				tf, err := svc.CreateTransfers(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, tf)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}

			svc := New(nil)
			svc.SetStore(mockStore)

			tc.buildStubs(mockStore)
			tc.checkResponse(t, svc)

			mockStore.AssertExpectations(t)
		})
	}
}

func TestGetTransfersByFromWalletID(t *testing.T) {
	testCase := []struct {
		name          string
		fromWalletID  int32
		buildStubs    func(mockStore *dbmocks.Store)
		checkResponse func(t *testing.T, svc Service)
	}{
		{
			name:         "OK",
			fromWalletID: dummyTransfer.FromWalletID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransfersByFromWalletID", mock.Anything, dummyTransfer.FromWalletID).
					Return([]db.Transfer{dummyTransfer}, nil)
			},
			checkResponse: func(t *testing.T, svc Service) {
				transfers, err := svc.GetTransfersByFromWalletID(context.Background(), dummyTransfer.FromWalletID)
				require.NoError(t, err)
				require.NotEmpty(t, transfers)
				require.Len(t, transfers, 1)

				transfer := transfers[0]
				require.Equal(t, dummyTransfer.ID, transfer.ID)
				require.Equal(t, dummyTransfer.FromWalletID, transfer.FromWalletID)
				require.Equal(t, dummyTransfer.ToWalletID, transfer.ToWalletID)
				require.Equal(t, dummyTransfer.Amount, transfer.Amount)
				require.Equal(t, dummyTransfer.Description, transfer.Description)
			},
		},
		{
			name:         "Unexpected Error",
			fromWalletID: dummyTransfer.FromWalletID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransfersByFromWalletID", mock.Anything, dummyTransfer.FromWalletID).
					Return(nil, errors.New("unexpected error"))
			},
			checkResponse: func(t *testing.T, svc Service) {
				transfers, err := svc.GetTransfersByFromWalletID(context.Background(), dummyTransfer.FromWalletID)
				require.Error(t, err)
				require.Empty(t, transfers)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}

			svc := New(nil)
			svc.SetStore(mockStore)

			tc.buildStubs(mockStore)
			tc.checkResponse(t, svc)

			mockStore.AssertExpectations(t)
		})
	}
}

func TestGetTransfersByFromWalletIdAndToWalletId(t *testing.T) {
	arg := db.GetTransfersByFromWalletIdAndToWalletIdParams{
		FromWalletID: dummyTransfer.FromWalletID,
		ToWalletID:   dummyTransfer.ToWalletID,
	}

	testCase := []struct {
		name          string
		arg           db.GetTransfersByFromWalletIdAndToWalletIdParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkResponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransfersByFromWalletIdAndToWalletId", mock.Anything, arg).
					Return([]db.Transfer{dummyTransfer}, nil)
			},
			checkResponse: func(t *testing.T, svc Service) {
				transfers, err := svc.GetTransfersByFromWalletIdAndToWalletId(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, transfers)
				require.Len(t, transfers, 1)

				transfer := transfers[0]
				require.Equal(t, dummyTransfer.ID, transfer.ID)
				require.Equal(t, dummyTransfer.FromWalletID, transfer.FromWalletID)
				require.Equal(t, dummyTransfer.ToWalletID, transfer.ToWalletID)
				require.Equal(t, dummyTransfer.Amount, transfer.Amount)
				require.Equal(t, dummyTransfer.Description, transfer.Description)
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransfersByFromWalletIdAndToWalletId", mock.Anything, arg).
					Return(nil, errors.New("unexpected error"))
			},
			checkResponse: func(t *testing.T, svc Service) {
				transfers, err := svc.GetTransfersByFromWalletIdAndToWalletId(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, transfers)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}

			svc := New(nil)
			svc.SetStore(mockStore)

			tc.buildStubs(mockStore)
			tc.checkResponse(t, svc)

			mockStore.AssertExpectations(t)
		})
	}
}
