package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomMerchants(t *testing.T) Merchant {
	ctx := context.Background()

	// Create a sample merchant for testing
	arg := CreateMerchantsParams{
		MerchantName: "Test Merchant",
		Description:  "Test Description",
		Website:      "testmerchant.com",
		Address:      "Test Address",
	}

	// Call the CreateMerchants function
	user, err := testQueries.CreateMerchants(ctx, arg)
	require.NoError(t, err)
	require.NotZero(t, arg)
	require.Equal(t, arg.MerchantName, user.MerchantName)
	require.Equal(t, arg.Description, user.Description)
	require.Equal(t, arg.Website, user.Website)
	require.Equal(t, arg.Address, user.Address)

	return user
}

func TestCreateMerchants(t *testing.T) {
	createRandomMerchants(t)
}

func TestDeleteMerchants(t *testing.T) {
	user := createRandomMerchants(t) //manipulasi data merchant
	ctx := context.Background()

	err := testQueries.DeleteMerchants(ctx, user.ID) //menghapus user
	require.NoError(t, err)

	user1, err := testQueries.GetMerchantsById(ctx, user.ID) //get user dari id
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user1)
}

func TestGetMerchantsById(t *testing.T) {
	user := createRandomMerchants(t)
	ctx := context.Background()

	// Get the merchant by ID
	merchant, err := testQueries.GetMerchantsById(ctx, user.ID)
	require.NoError(t, err)
	require.NotZero(t, merchant)
	require.Equal(t, user.MerchantName, merchant.MerchantName)
	require.Equal(t, user.Description, merchant.Description)
	require.Equal(t, user.Website, merchant.Website)
	require.Equal(t, user.Address, merchant.Address)
}

func TestGetMerchantsByIdNotFound(t *testing.T) {
	ctx := context.Background()

	// Get the merchant by ID
	merchant, err := testQueries.GetMerchantsById(ctx, 12345)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, merchant)
}

func TestGetMerchantsByMerchantsName(t *testing.T) {
	user := createRandomMerchants(t)
	ctx := context.Background()

	// Get the merchant by name
	merchant, err := testQueries.GetMerchantsByMerchantsName(ctx, user.MerchantName)
	require.NoError(t, err)
	require.NotZero(t, merchant)
	require.Equal(t, user.MerchantName, merchant.MerchantName)
	require.Equal(t, user.Description, merchant.Description)
	require.Equal(t, user.Website, merchant.Website)
	require.Equal(t, user.Address, merchant.Address)
}

func TestGetMerchantsByMerchantsNameNotFound(t *testing.T) {
	ctx := context.Background()

	// Get the merchant by name
	merchant, err := testQueries.GetMerchantsByMerchantsName(ctx, "merchant name that does not exist")
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, merchant)
}

func TestListMerchants(t *testing.T) {
	n := 5
	for i := 0; i < n*2; i++ {
		createRandomMerchants(t)
	}

	arg := ListMerchantsParams{
		Limit:  int32(n),
		Offset: 0,
	}

	merchant, err := testQueries.ListMerchants(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, merchant)

	for _, user := range merchant {
		require.NotEmpty(t, user)
	}
}

func TestUpdate(t *testing.T) {
	merchant := createRandomMerchants(t)
	ctx := context.Background()

	arg := UpdatMerchantsParams{
		ID:          merchant.ID,
		Description: "lindungi koruptor",
		Address:     "Depok Aja",
	}

	updatedMerchants, err := testQueries.UpdatMerchants(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedMerchants)

	require.Equal(t, arg.ID, updatedMerchants.ID)
	require.Equal(t, arg.Description, updatedMerchants.Description)
	require.Equal(t, arg.Address, updatedMerchants.Address)
}
