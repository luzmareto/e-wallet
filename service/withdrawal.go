package service

import (
	"context"
	"fmt"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

// Createwithdrawals implements Service
func (s *service) CreateWithdrawals(ctx context.Context, arg db.CreateWithdrawalsParams) (db.Withdrawal, error) {
	if _, err := s.queries.GetUserById(ctx, int64(arg.UserID)); err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("user with id %d not found", arg.UserID),
			Err: err,
		}
		return db.Withdrawal{}, cstErr
	}
	return s.queries.CreateWithdrawals(ctx, arg)
}
