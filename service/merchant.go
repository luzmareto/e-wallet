package service

import (
	"context"
	"fmt"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

// CreateMerchants implements Service
func (s *service) CreateMerchants(ctx context.Context, arg db.CreateMerchantsParams) (db.Merchant, error) {
	merchant, err := s.queries.CreateMerchants(ctx, arg)
	if err != nil {
		return db.Merchant{}, err
	}
	return merchant, nil
}

// DeleteMerchants implements Service
func (s *service) DeleteMerchants(ctx context.Context, id int64) error {
	err := s.queries.DeleteMerchants(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

// GetMerchantsById implements Service
func (s *service) GetMerchantsById(ctx context.Context, id int64) (db.Merchant, error) {
	var merchant db.Merchant

	merchant, err := s.queries.GetMerchantsById(ctx, id)
	if err != nil {
		return db.Merchant{}, err
	}
	return merchant, nil
}

// GetMerchantsByMerchantsName implements Service
func (s *service) GetMerchantsByMerchantsName(ctx context.Context, merchantName string) (db.Merchant, error) {
	var merchant db.Merchant

	merchant, err := s.queries.GetMerchantsByMerchantsName(ctx, merchantName)
	if err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("merchant with merchantname %s not found", merchantName),
			Err: err,
		}
		return merchant, cstErr
	}
	return merchant, nil
}

// ListMerchants implements Service
func (s *service) ListMerchants(ctx context.Context, arg db.ListMerchantsParams) ([]db.Merchant, error) {
	merchants, err := s.queries.ListMerchants(ctx, arg)
	if err != nil {
		return nil, err
	}
	return merchants, nil

}

// UpdatMerchants implements Service
func (s *service) UpdatMerchants(ctx context.Context, arg db.UpdatMerchantsParams) (db.Merchant, error) {
	var merchant db.Merchant

	merchant, err := s.queries.UpdatMerchants(ctx, arg)
	if err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("user with id %d not found", arg.ID),
			Err: err,
		}
		return merchant, cstErr
	}
	return s.queries.UpdatMerchants(ctx, arg)
}
