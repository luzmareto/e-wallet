package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	toLongPassword := RandomString(100)
	hashpassword, err := HashPassword(toLongPassword)
	require.Error(t, err)
	require.Empty(t, hashpassword)

	password := RandomString(8)

	hashedpassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedpassword)

	err = CheckPassword(password, hashedpassword)
	require.NoError(t, err)

	wrongpassword := RandomString(8)
	err = CheckPassword(wrongpassword, hashedpassword)
	require.EqualError(t, bcrypt.ErrMismatchedHashAndPassword, err.Error())
}
