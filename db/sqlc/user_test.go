package db

import (
	"context"
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
