package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	dbmocks "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/mocks"
	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
)

var dummyUser = []db.User{
	{
		ID:               1,
		Role:             "admin",
		Username:         "admin",
		Password:         "password-admin",
		Email:            "admin@mail.com",
		PhoneNumber:      "012345678",
		IDCard:           "jhkhkj.jpg",
		RegistrationDate: time.Now(),
	},
	{
		ID:               2,
		Role:             "user",
		Username:         "user",
		Password:         "password-user",
		Email:            "user@mail.com",
		PhoneNumber:      "012345678",
		IDCard:           "jhkhkj.jpg",
		RegistrationDate: time.Now(),
	},
	{
		ID:               3,
		Role:             "user3",
		Username:         "user3",
		Password:         "password-user3",
		Email:            "user5@mail.com",
		PhoneNumber:      "012345678",
		IDCard:           "jhkhkj.jpg",
		RegistrationDate: time.Now(),
	},
	{
		ID:               4,
		Role:             "user4",
		Username:         "user4",
		Password:         "password-user4",
		Email:            "user@mail.com",
		PhoneNumber:      "012345678",
		IDCard:           "jhkhkj.jpg",
		RegistrationDate: time.Now(),
	},
	{
		ID:               5,
		Role:             "user5",
		Username:         "user5",
		Password:         "password-user5",
		Email:            "user@mail.com",
		PhoneNumber:      "012345678",
		IDCard:           "jhkhkj.jpg",
		RegistrationDate: time.Now(),
	},
}

func TestCreateUser(t *testing.T) {
	testCase := []struct {
		name          string
		arg           db.CreateUsersParams
		buildStubs    func(mockStore *dbmocks.Store, user db.User)
		checkresponse func(t *testing.T, svc Service, arg db.CreateUsersParams)
	}{
		{
			name: "OK",
			arg: db.CreateUsersParams{
				Username:    dummyUser[0].Username,
				Password:    dummyUser[0].Password,
				Email:       dummyUser[0].Email,
				PhoneNumber: dummyUser[0].PhoneNumber,
			},
			buildStubs: func(mockStore *dbmocks.Store, user db.User) {
				mockStore.On("CreateUsers", mock.Anything, mock.Anything).
					Return(user, nil)
			},
			checkresponse: func(t *testing.T, svc Service, arg db.CreateUsersParams) {
				user, err := svc.CreateUsers(context.Background(), arg)
				require.NoError(t, err)
				require.NotEmpty(t, user)
				require.Equal(t, arg.Username, user.Username)
				require.Equal(t, arg.Password, user.Password)
				require.Equal(t, arg.Email, user.Email)
				require.Equal(t, arg.PhoneNumber, user.PhoneNumber)
			},
		},
		{
			name: "Error",
			arg: db.CreateUsersParams{
				Username:    dummyUser[0].Username,
				Password:    dummyUser[0].Password,
				Email:       dummyUser[0].Email,
				PhoneNumber: dummyUser[0].PhoneNumber,
			},
			buildStubs: func(mockStore *dbmocks.Store, user db.User) {
				mockStore.On("CreateUsers", mock.Anything, mock.Anything).
					Return(user, errors.New("error creating user"))
			},
			checkresponse: func(t *testing.T, svc Service, arg db.CreateUsersParams) {
				user, err := svc.CreateUsers(context.Background(), arg)
				require.Error(t, err)
				require.Empty(t, user)

			},
		},
	}

	for i := range testCase {
		tc := testCase[i]
		t.Run(tc.name, func(t *testing.T) {
			mockStore := &dbmocks.Store{}
			svc := New(nil)
			svc.SetStore(mockStore)

			tc.buildStubs(mockStore, dummyUser[0])
			tc.checkresponse(t, svc, tc.arg)

			mockStore.AssertExpectations(t)
		})
	}

}

func TestDeleteUser(t *testing.T) {
	testCase := []struct {
		name          string
		useID         int64
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name:  "OK",
			useID: dummyUser[0].ID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.User{}, nil)
				mockStore.On("DeleteUsers", mock.Anything, mock.AnythingOfType("int64")).
					Return(nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.DeleteUsers(context.Background(), dummyUser[0].ID)
				require.NoError(t, err)
			},
		},
		{
			name:  "Error",
			useID: dummyUser[0].ID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.DeleteUsers(context.Background(), dummyUser[0].ID)
				require.Error(t, err)
				require.EqualError(t, err, fmt.Sprintf("user with id %d not found", dummyUser[0].ID))
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

func TestGetUserByID(t *testing.T) {
	testCase := []struct {
		name          string
		useID         int64
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name:  "OK",
			useID: dummyUser[0].ID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.AnythingOfType("int64")).
					Return(dummyUser[0], nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				user, err := svc.GetUserById(context.Background(), dummyUser[0].ID)
				require.NoError(t, err)
				require.NotEmpty(t, user)
				require.Equal(t, dummyUser[0].ID, user.ID)
				require.Equal(t, dummyUser[0].Role, user.Role)
				require.Equal(t, dummyUser[0].Username, user.Username)
				require.Equal(t, dummyUser[0].Password, user.Password)
				require.Equal(t, dummyUser[0].Email, user.Email)
				require.Equal(t, dummyUser[0].PhoneNumber, user.PhoneNumber)
				require.Equal(t, dummyUser[0].IDCard, user.IDCard)
				require.Equal(t, dummyUser[0].RegistrationDate, user.RegistrationDate)
			},
		},
		{
			name:  "Not Found",
			useID: dummyUser[0].ID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				user, err := svc.GetUserById(context.Background(), dummyUser[0].ID)
				require.Error(t, err)
				require.Empty(t, user)
				require.EqualError(t, err, fmt.Sprintf("user with id %d not found", dummyUser[0].ID))
			},
		},
		{
			name:  "Unexpected error",
			useID: dummyUser[0].ID,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.AnythingOfType("int64")).
					Return(db.User{}, errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				user, err := svc.GetUserById(context.Background(), dummyUser[0].ID)
				require.Error(t, err)
				require.Empty(t, user)
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

func TestGetUserByUsername(t *testing.T) {
	testCase := []struct {
		name          string
		username      string
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name:     "OK",
			username: dummyUser[0].Username,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserByUserName", mock.Anything, mock.AnythingOfType("string")).
					Return(dummyUser[0], nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				user, err := svc.GetUserByUserName(context.Background(), dummyUser[0].Username)
				require.NoError(t, err)
				require.NotEmpty(t, user)
				require.Equal(t, dummyUser[0].ID, user.ID)
				require.Equal(t, dummyUser[0].Role, user.Role)
				require.Equal(t, dummyUser[0].Username, user.Username)
				require.Equal(t, dummyUser[0].Password, user.Password)
				require.Equal(t, dummyUser[0].Email, user.Email)
				require.Equal(t, dummyUser[0].PhoneNumber, user.PhoneNumber)
				require.Equal(t, dummyUser[0].IDCard, user.IDCard)
				require.Equal(t, dummyUser[0].RegistrationDate, user.RegistrationDate)
			},
		},
		{
			name:     "Not Found",
			username: dummyUser[0].Username,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserByUserName", mock.Anything, mock.AnythingOfType("string")).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				user, err := svc.GetUserByUserName(context.Background(), dummyUser[0].Username)
				require.Error(t, err)
				require.Empty(t, user)
				require.EqualError(t, err, fmt.Sprintf("user with username %s not found", dummyUser[0].Username))
			},
		},
		{
			name:     "Unexpected error",
			username: dummyUser[0].Username,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserByUserName", mock.Anything, mock.AnythingOfType("string")).
					Return(db.User{}, errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				user, err := svc.GetUserByUserName(context.Background(), dummyUser[0].Username)
				require.Error(t, err)
				require.Empty(t, user)
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

func TestListUser(t *testing.T) {
	argListUsers := db.ListUsersParams{
		Limit:  5,
		Offset: 5,
	}
	testCase := []struct {
		name          string
		arg           db.ListUsersParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  argListUsers,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("ListUsers", mock.Anything, mock.Anything).
					Return(dummyUser, nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				users, err := svc.ListUsers(context.Background(), argListUsers)
				require.NoError(t, err)
				require.NotEmpty(t, users)
				require.Equal(t, argListUsers.Limit, int32(len(users)))
			},
		},
		{
			name: "Unexpected Error",
			arg:  argListUsers,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("ListUsers", mock.Anything, mock.Anything).
					Return([]db.User{}, errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				users, err := svc.ListUsers(context.Background(), argListUsers)
				require.Error(t, err)
				require.Empty(t, users)
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

func TestUpdateUsers(t *testing.T) {
	argUpdateUser := db.UpdateUsersParams{
		ID:          dummyUser[1].ID,
		Email:       "abc@mail.com",
		PhoneNumber: "0812345678910",
	}
	testCase := []struct {
		name          string
		arg           db.UpdateUsersParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  argUpdateUser,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.Anything).
					Return(dummyUser[1], nil)
				mockStore.On("UpdateUsers", mock.Anything, mock.Anything).
					Return(dummyUser[1], nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				user, err := svc.UpdateUsers(context.Background(), argUpdateUser)
				require.NoError(t, err)
				require.NotEmpty(t, user)
				require.Equal(t, argUpdateUser.ID, user.ID)
			},
		},
		{
			name: "Not found",
			arg:  argUpdateUser,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.Anything).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				user, err := svc.UpdateUsers(context.Background(), argUpdateUser)
				require.Error(t, err)
				require.Empty(t, user)
				require.EqualError(t, err, fmt.Sprintf("user with id %d not found", argUpdateUser.ID))
			},
		},
		{
			name: "Unexpected Error",
			arg:  argUpdateUser,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.Anything).
					Return(dummyUser[1], nil)
				mockStore.On("UpdateUsers", mock.Anything, mock.Anything).
					Return(db.User{}, errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				user, err := svc.UpdateUsers(context.Background(), argUpdateUser)
				require.Error(t, err)
				require.Empty(t, user)
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

func TestUpdateUsersPassword(t *testing.T) {
	arg := db.UpdateUsersPasswordParams{
		ID:       dummyUser[1].ID,
		Password: "abcdefghijkl",
	}
	testCase := []struct {
		name          string
		arg           db.UpdateUsersPasswordParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.Anything).
					Return(db.User{}, nil)
				mockStore.On("UpdateUsersPassword", mock.Anything, mock.Anything).
					Return(nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.UpdateUsersPassword(context.Background(), arg)
				require.NoError(t, err)
			},
		},
		{
			name: "Not Found",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.Anything).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.UpdateUsersPassword(context.Background(), arg)
				require.Error(t, err)
				require.EqualError(t, err, fmt.Sprintf("user with id %d not found", arg.ID))
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.Anything).
					Return(db.User{}, nil)
				mockStore.On("UpdateUsersPassword", mock.Anything, mock.Anything).
					Return(errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.UpdateUsersPassword(context.Background(), arg)
				require.Error(t, err)
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

func TestUpdateUserIDcard(t *testing.T) {
	arg := db.UpdateUserIDcardParams{
		ID:     dummyUser[1].ID,
		IDCard: "ktp.jpg",
	}
	testCase := []struct {
		name          string
		arg           db.UpdateUserIDcardParams
		buildStubs    func(mockStore *dbmocks.Store)
		checkresponse func(t *testing.T, svc Service)
	}{
		{
			name: "OK",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.Anything).
					Return(db.User{}, nil)
				mockStore.On("UpdateUserIDcard", mock.Anything, mock.Anything).
					Return(nil)
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.UpdateUserIDcard(context.Background(), arg)
				require.NoError(t, err)
			},
		},
		{
			name: "Not Found",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.Anything).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.UpdateUserIDcard(context.Background(), arg)
				require.Error(t, err)
				require.EqualError(t, err, fmt.Sprintf("user with id %d not found", arg.ID))
			},
		},
		{
			name: "Unexpected Error",
			arg:  arg,
			buildStubs: func(mockStore *dbmocks.Store) {
				mockStore.On("GetUserById", mock.Anything, mock.Anything).
					Return(db.User{}, nil)
				mockStore.On("UpdateUserIDcard", mock.Anything, mock.Anything).
					Return(errors.New(""))
			},
			checkresponse: func(t *testing.T, svc Service) {
				err := svc.UpdateUserIDcard(context.Background(), arg)
				require.Error(t, err)
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
