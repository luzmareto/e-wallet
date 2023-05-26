package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

func createRandomUser1(t *testing.T) User {
	username := utils.RandomOwner()

	arg := CreateUsersParams{
		Username: username,
		// Set other required fields as needed
	}

	user, err := testQueries.CreateUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, user)

	return user
}

func createRandomSession(t *testing.T) {
	ctx := context.Background()

	// Create a random user
	user := createRandomUser1(t)

	// Create a session
	arg := CreateSessionParams{
		ID:           [16]byte{byte(utils.RandomInt(1, 1000))},
		Username:     user.Username,
		RefreshToken: "refresh-token",
		UserAgent:    "user-agent",
		ClientIp:     utils.RandomString(100),
		IsBlocked:    false,
		ExpiredAt:    time.Now().AddDate(1, 0, 0),
	}

	data, err := testQueries.CreateSession(ctx, arg)
	require.NoError(t, err)
	require.NotZero(t, data)
	require.Equal(t, arg.Username, data.Username)
	require.Equal(t, arg.RefreshToken, data.RefreshToken)
	require.Equal(t, arg.UserAgent, data.UserAgent)
	require.Equal(t, arg.ClientIp, data.ClientIp)
	require.Equal(t, false, data.IsBlocked)
	// require.Equal(t, time.Now().AddDate(1, 0, 0), data.ExpiredAt)
}

func TestCreateSession(t *testing.T) {
	createRandomSession(t)
}
func TestGetSessions(t *testing.T) {
	user := createRandomUser(t)

	arg := CreateSessionParams{
		ID:           [16]byte{byte(utils.RandomInt(1, 10000000))},
		Username:     user.Username,
		RefreshToken: "refresh-token",
		UserAgent:    "user-agent",
		ClientIp:     utils.RandomString(100),
		IsBlocked:    false,
		ExpiredAt:    time.Now().AddDate(1, 0, 0),
	}

	_, err := testQueries.CreateSession(context.Background(), arg)
	require.NoError(t, err)

	session, err := testQueries.GetSessions(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotZero(t, session)
	require.Equal(t, arg.Username, session.Username)
	require.Equal(t, arg.RefreshToken, session.RefreshToken)
	require.Equal(t, arg.UserAgent, session.UserAgent)
	require.Equal(t, arg.ClientIp, session.ClientIp)
	require.Equal(t, arg.IsBlocked, session.IsBlocked)
	require.WithinDuration(t, arg.ExpiredAt, session.ExpiredAt, time.Second)

	// Get a session that does not exist
	_, err = testQueries.GetSessions(context.Background(), uuid.New())
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
