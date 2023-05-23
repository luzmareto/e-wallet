package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

func TestUpdateUsers(t *testing.T) {
	// Create user
	users := createUser(t)

	// Update user's email and phone number
	arg := UpdateUsersParams{
		ID:          users.ID,
		Email:       "presiden@gmail.com",
		PhoneNumber: "0123456",
	}

	// Call UpdateUsers function
	updatedUser, err := testQueries.UpdateUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	// Verify the updated user's fields
	require.Equal(t, arg.ID, updatedUser.ID)
	require.Equal(t, users.Username, updatedUser.Username)
	require.Equal(t, arg.Email, updatedUser.Email)
	require.Equal(t, arg.PhoneNumber, updatedUser.PhoneNumber)
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

func TestListUsers(t *testing.T) {
	n := 5
	for i := 0; i < n*2; i++ {
		createUser(t)
	}

	arg := ListUsersParams{
		Limit:  int32(n),
		Offset: 0,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestGetUserByUserName(t *testing.T) {
	user := createUser(t)
	user1, err := testQueries.GetUserByUserName(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user.ID, user1.ID)
	require.Equal(t, user.Username, user1.Username)
	require.Equal(t, user.Email, user1.Email)
	require.Equal(t, user.PhoneNumber, user1.PhoneNumber)

	require.WithinDuration(t, user.RegistrationDate, user1.RegistrationDate, time.Second)
}
