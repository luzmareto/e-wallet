package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

func TestGetUserById(t *testing.T) {
	user := createUser(t)
	user1, err := testQueries.GetUserById(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user.ID, user1.ID)
	require.Equal(t, user.Email, user1.Email)
	require.Equal(t, user.PhoneNumber, user1.PhoneNumber)

	require.WithinDuration(t, user.RegistrationDate, user1.RegistrationDate, time.Second)
}

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
	user := createUser(t) // membuat user baru

	err := testQueries.DeleteUsers(context.Background(), user.ID) //menghapus user
	require.NoError(t, err)

	user1, err := testQueries.GetUserById(context.Background(), user.ID) //get user dari id
	//melakukan pengecekan apakah user sudah terhapus atau belum
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user1)
}
