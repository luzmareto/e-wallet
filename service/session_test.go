package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	dbmocks "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/mocks"
	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var dummySession = []db.Session{
	{
		ID:           [16]byte{1},
		Username:     "user",
		RefreshToken: "user",
		UserAgent:    "user",
		ClientIp:     "user",
		IsBlocked:    false,
		ExpiredAt:    time.Now(),
		CreatedAt:    time.Now(),
	},
	{
		ID:           [16]byte{2},
		Username:     "user",
		RefreshToken: "user",
		UserAgent:    "user",
		ClientIp:     "user",
		IsBlocked:    false,
		ExpiredAt:    time.Now(),
		CreatedAt:    time.Now(),
	},
}

func TestCreateSession(t *testing.T) {
	testCase := []struct {
		name          string
		arg           db.CreateSessionParams
		buildStubs    func(mockStore *dbmocks.Store, session db.Session)
		checkresponse func(t *testing.T, svc Service, arg db.CreateSessionParams)
	}{
		{
			name: "OK",
			arg: db.CreateSessionParams{
				ID:           dummySession[0].ID,
				Username:     dummySession[0].Username,
				RefreshToken: dummySession[0].RefreshToken,
				UserAgent:    dummySession[0].UserAgent,
				ClientIp:     dummySession[0].ClientIp,
				IsBlocked:    false,
				ExpiredAt:    time.Now(),
			},
			buildStubs: func(mockStore *dbmocks.Store, session db.Session) {
				mockStore.On("CreateSession", mock.Anything, mock.Anything).
					Return(session, nil)
			},
			checkresponse: func(t *testing.T, svc Service, arg db.CreateSessionParams) {
				session, err := svc.CreateSession(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, session)
				require.Equal(t, arg.ID, session.ID)
				require.Equal(t, arg.Username, session.Username)
				require.Equal(t, arg.RefreshToken, session.RefreshToken)
				require.Equal(t, arg.UserAgent, session.UserAgent)
				require.Equal(t, arg.ClientIp, session.ClientIp)
				require.Equal(t, arg.IsBlocked, session.IsBlocked)
				require.WithinDuration(t, arg.ExpiredAt, session.ExpiredAt, time.Second)
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}
			svc := New(nil)
			svc.SetStore(mockStore)

			tc.buildStubs(mockStore, dummySession[0])
			tc.checkresponse(t, svc, tc.arg)

			mockStore.AssertExpectations(t)
		})
	}
}

func TestGetSessions(t *testing.T) {
	testCase := []struct {
		name          string
		id            uuid.UUID
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			id:   dummySession[0].ID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetSessions", mock.Anything, mock.Anything).
					Return(dummySession[0], nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				sessionList, err := svc.GetSessions(context.Background(), dummySession[0].ID)
				require.NoError(t, err)
				require.NotEmpty(t, sessionList)

				session := sessionList
				require.Equal(t, dummySession[0].ID, session.ID)
				require.Equal(t, dummySession[0].Username, session.Username)
				require.Equal(t, dummySession[0].RefreshToken, session.RefreshToken)
				require.Equal(t, dummySession[0].UserAgent, session.UserAgent)
				require.Equal(t, dummySession[0].ClientIp, session.ClientIp)
				require.Equal(t, dummySession[0].IsBlocked, session.IsBlocked)
				require.WithinDuration(t, dummySession[0].ExpiredAt, session.ExpiredAt, time.Second)
			},
		},
		{
			name: "Not Found",
			id:   dummySession[0].ID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetSessions", mock.Anything, mock.Anything).
					Return(db.Session{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				sessionList, err := svc.GetSessions(context.Background(), dummySession[0].ID)
				require.Error(t, err)
				require.Empty(t, sessionList)
				require.EqualError(t, err, "sql: no rows in result set")

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
