package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

func createUser(t *testing.T) User {
	ctx := context.Background()
	hashed, err := utils.HashPassword(utils.RandomString(12))
	require.NoError(t, err)
	arg := CreateUsersParams{
		Username:    utils.RandomOwner(),
		Password:    hashed,
		Email:       utils.RandomEmail(),
		PhoneNumber: utils.RandomString(12),
	}
	user, err := testQueries.CreateUsers(ctx, arg)
	require.NoError(t, err)                             // check no error
	require.Equal(t, arg.Username, user.Username)       // check must must be equal
	require.Equal(t, arg.Password, user.Password)       //
	require.Equal(t, arg.PhoneNumber, user.PhoneNumber) // check must must be equal
	require.Equal(t, arg.Email, user.Email)             // check must must be equal
	// require.WithinDuration(t, time.Now(), user.RegistrationDate, time.Millisecond) /// check must must be in duration

	return user
}
func TestCreateUser(t *testing.T) {
	createUser(t)
}

func TestDeleteUser(t *testing.T) {
	/*
		ALUR :
			1. Membuat User baru
			2. Menghapus data dari User baru
			3. Melakukan pengecekan
	*/

	user := createUser(t) // membuat user baru

	err := testQueries.DeleteUsers(context.Background(), user.ID) //menghapus user
	require.NoError(t, err)

	user1, err := testQueries.GetUserById(context.Background(), user.ID) //get user dari id
	//melakukan pengecekan apakah user sudah terhapus atau belum
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user1)
}
