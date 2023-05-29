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

func TestCreateTransfers(t *testing.T) {
	wallet1 := db.Wallet{
		ID:      1,
		Balance: 1000,
	}

	wallet2 := db.Wallet{
		ID:      2,
		Balance: 1000,
	}
	arg := db.CreateTransfersParams{
		FromWalletID: int32(wallet1.ID),
		ToWalletID:   int32(wallet2.ID),
		Amount:       500,
		Description:  utils.RandomString(10),
	}

	transfer := db.Transfer{
		ID:           1,
		FromWalletID: int32(wallet1.ID),
		ToWalletID:   int32(wallet2.ID),
		Amount:       arg.Amount,
		TransferDate: time.Now(),
		Description:  arg.Description,
	}

	testCase := []struct {
		name          string
		arg           db.CreateTransfersParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("CreateTransfers", mock.Anything, mock.Anything).
					Return(transfer, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				transfer, err := svc.CreateTransfers(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, transfer)
				require.NotZero(t, transfer.ID)
				require.Equal(t, arg.FromWalletID, transfer.FromWalletID)
				require.Equal(t, arg.ToWalletID, transfer.ToWalletID)
				require.Equal(t, arg.Amount, transfer.Amount)
				require.Equal(t, arg.Description, transfer.Description)
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("CreateTransfers", mock.Anything, mock.Anything).
					Return(db.Transfer{}, errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				transfer, err := svc.CreateTransfers(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, transfer)
				require.Zero(t, transfer.ID)
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

func TestGetTransfersByFromWalletID(t *testing.T) {
	result := []db.Transfer{
		{
			ID:           1,
			FromWalletID: 1,
			ToWalletID:   2,
			Amount:       1000,
			TransferDate: time.Now(),
			Description:  "test description",
		},
		{
			ID:           1,
			FromWalletID: 2,
			ToWalletID:   1,
			Amount:       1000,
			TransferDate: time.Now(),
			Description:  "test description",
		},
	}
	arg := int32(1)
	testCase := []struct {
		name          string
		arg           int32
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransfersByFromWalletID", mock.Anything, mock.AnythingOfType("int32")).
					Return(result, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				transfer, err := svc.GetTransfersByFromWalletID(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, transfer)
				require.GreaterOrEqual(t, len(result), 1)
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransfersByFromWalletID", mock.Anything, mock.AnythingOfType("int32")).
					Return([]db.Transfer{}, errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				transfer, err := svc.GetTransfersByFromWalletID(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, transfer)
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

func TestGetTransfersByFromWalletIdAndToWalletId(t *testing.T) {
	result := []db.Transfer{
		{
			ID:           1,
			FromWalletID: 1,
			ToWalletID:   2,
			Amount:       1000,
			TransferDate: time.Now(),
			Description:  "test description",
		},
		{
			ID:           1,
			FromWalletID: 2,
			ToWalletID:   1,
			Amount:       1000,
			TransferDate: time.Now(),
			Description:  "test description",
		},
	}
	arg := db.GetTransfersByFromWalletIdAndToWalletIdParams{
		FromWalletID: int32(result[0].ID),
		ToWalletID:   int32(result[1].ID),
	}
	testCase := []struct {
		name          string
		arg           db.GetTransfersByFromWalletIdAndToWalletIdParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransfersByFromWalletIdAndToWalletId", mock.Anything, mock.Anything).
					Return(result, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				transfer, err := svc.GetTransfersByFromWalletIdAndToWalletId(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, transfer)
				require.GreaterOrEqual(t, len(result), 1)
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetTransfersByFromWalletIdAndToWalletId", mock.Anything, mock.Anything).
					Return([]db.Transfer{}, errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				transfer, err := svc.GetTransfersByFromWalletIdAndToWalletId(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, transfer)
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
