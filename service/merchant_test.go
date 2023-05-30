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
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

func TestCreateMerchants(t *testing.T) {
	arg := db.CreateMerchantsParams{
		MerchantName: utils.RandomString(4),
		Description:  utils.RandomString(100),
		Website:      utils.RandomString(10),
		Address:      utils.RandomString(50),
	}
	dummyMerchant := db.Merchant{
		ID:           1,
		MerchantName: arg.MerchantName,
		Description:  arg.Description,
		Website:      arg.Website,
		Address:      arg.Address,
		Balance:      0,
	}
	testCase := []struct {
		name          string
		arg           db.CreateMerchantsParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("CreateMerchants", mock.Anything, mock.AnythingOfTypeArgument("db.CreateMerchantsParams")).
					Return(dummyMerchant, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				merchant, err := svc.CreateMerchants(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, merchant)
				require.NotZero(t, merchant.ID)
				require.Equal(t, arg.MerchantName, merchant.MerchantName)
				require.Equal(t, arg.Description, merchant.Description)
				require.Equal(t, arg.Website, merchant.Website)
				require.Equal(t, arg.Address, merchant.Address)
			},
		},
		{
			name: "unexpected_error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("CreateMerchants", mock.Anything, mock.AnythingOfTypeArgument("db.CreateMerchantsParams")).
					Return(db.Merchant{}, errors.New("unexpected error"))
			},
			checkresponse: func(t *testing.T, svc Service) {
				merchant, err := svc.CreateMerchants(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, merchant)
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

func TestDeleteMerchants(t *testing.T) {
	testCase := []struct {
		name          string
		arg           int64
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  1,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("DeleteMerchants", mock.Anything, mock.AnythingOfType("int64")).
					Return(nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.DeleteMerchants(context.Background(), 1)
				require.NoError(t, err)
			},
		},
		{
			name: "not_found",
			arg:  0,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("DeleteMerchants", mock.Anything, mock.AnythingOfType("int64")).
					Return(sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.DeleteMerchants(context.Background(), 0)
				require.Error(t, err)
				require.EqualError(t, err, fmt.Sprintf("merchant with id %d not found", 0))
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

func TestGetMerchantsById(t *testing.T) {
	dummyMerchant := db.Merchant{
		ID:           1,
		MerchantName: utils.RandomString(4),
		Description:  utils.RandomString(100),
		Website:      utils.RandomString(10),
		Address:      utils.RandomString(50),
		Balance:      0,
	}
	testCase := []struct {
		name          string
		arg           int64
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  1,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetMerchantsById", mock.Anything, mock.AnythingOfType("int64")).
					Return(dummyMerchant, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				merchant, err := svc.GetMerchantsById(context.Background(), 1)
				require.NoError(t, err)
				require.NotEmpty(t, merchant)
			},
		},
		{
			name: "not_found",
			arg:  0,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetMerchantsById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.Merchant{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				merchant, err := svc.GetMerchantsById(context.Background(), 0)
				require.Error(t, err)
				require.Empty(t, merchant)
				require.EqualError(t, err, fmt.Sprintf("merchant with id %d not found", 0))
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

func TestGetMerchantsByMerchantsName(t *testing.T) {
	dummyMerchant := db.Merchant{
		ID:           1,
		MerchantName: utils.RandomString(4),
		Description:  utils.RandomString(100),
		Website:      utils.RandomString(10),
		Address:      utils.RandomString(50),
		Balance:      0,
	}
	testCase := []struct {
		name          string
		arg           string
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  dummyMerchant.MerchantName,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetMerchantsByMerchantsName", mock.Anything, mock.AnythingOfType("string")).
					Return(dummyMerchant, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				merchant, err := svc.GetMerchantsByMerchantsName(context.Background(), dummyMerchant.MerchantName)
				require.NoError(t, err)
				require.NotEmpty(t, merchant)
			},
		},
		{
			name: "not_found",
			arg:  "test",
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetMerchantsByMerchantsName", mock.Anything, mock.AnythingOfType("string")).
					Return(db.Merchant{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				merchant, err := svc.GetMerchantsByMerchantsName(context.Background(), "test")
				require.Error(t, err)
				require.Empty(t, merchant)
				require.EqualError(t, err, fmt.Sprintf("merchant with merchantname %s not found", "test"))
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

func TestListMerchants(t *testing.T) {
	arg := db.ListMerchantsParams{
		Limit:  5,
		Offset: 5,
	}
	dummyListMerchants := []db.Merchant{
		{
			ID:           1,
			MerchantName: utils.RandomString(4),
			Description:  utils.RandomString(100),
			Website:      utils.RandomString(10),
			Address:      utils.RandomString(50),
			Balance:      0,
		},
		{
			ID:           2,
			MerchantName: utils.RandomString(4),
			Description:  utils.RandomString(100),
			Website:      utils.RandomString(10),
			Address:      utils.RandomString(50),
			Balance:      0,
		},
	}
	testCase := []struct {
		name          string
		arg           db.ListMerchantsParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("ListMerchants", mock.Anything, mock.AnythingOfTypeArgument("db.ListMerchantsParams")).
					Return(dummyListMerchants, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.ListMerchants(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, result)
			},
		},
		{
			name: "unexpected_error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("ListMerchants", mock.Anything, mock.AnythingOfTypeArgument("db.ListMerchantsParams")).
					Return([]db.Merchant{}, errors.New("unexpected error"))
			},
			checkresponse: func(t *testing.T, svc Service) {
				result, err := svc.ListMerchants(context.Background(), arg)
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

func TestUpdatMerchants(t *testing.T) {
	arg := db.UpdatMerchantsParams{
		ID:          1,
		Description: utils.RandomString(100),
		Address:     utils.RandomString(10),
	}
	dummyMerchant := db.Merchant{
		ID:           1,
		MerchantName: utils.RandomString(4),
		Description:  arg.Description,
		Website:      utils.RandomString(10),
		Address:      arg.Address,
		Balance:      0,
	}

	testCase := []struct {
		name          string
		arg           db.UpdatMerchantsParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("UpdatMerchants", mock.Anything, mock.AnythingOfTypeArgument("db.UpdatMerchantsParams")).
					Return(dummyMerchant, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				merchant, err := svc.UpdatMerchants(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, merchant)
				require.NotZero(t, merchant.ID)
				require.Equal(t, arg.Description, merchant.Description)
				require.Equal(t, arg.Address, merchant.Address)
			},
		},
		{
			name: "not_found",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("UpdatMerchants", mock.Anything, mock.AnythingOfTypeArgument("db.UpdatMerchantsParams")).
					Return(db.Merchant{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				merchant, err := svc.UpdatMerchants(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, merchant)
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

func TestAddMerchantBalance(t *testing.T) {
	dummyMerchant := db.Merchant{
		ID:           1,
		MerchantName: utils.RandomString(4),
		Description:  utils.RandomString(100),
		Website:      utils.RandomString(10),
		Address:      utils.RandomString(20),
		Balance:      0,
	}

	arg := db.AddMerchantBalanceParams{
		ID:      dummyMerchant.ID,
		Balance: 1000,
	}

	testCase := []struct {
		name          string
		arg           db.AddMerchantBalanceParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetMerchantsById", mock.Anything, mock.AnythingOfType("int64")).
					Return(dummyMerchant, nil)
				mockStore.On("AddMerchantBalance", mock.Anything, mock.AnythingOfTypeArgument("AddMerchantBalanceParams")).
					Return(dummyMerchant, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				merchant, err := svc.AddMerchantBalance(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, merchant)
			},
		},
		{
			name: "not_found",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetMerchantsById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.Merchant{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				merchant, err := svc.AddMerchantBalance(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, merchant)
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
