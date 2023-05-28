package service

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	db "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/db/sqlc"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

type Service interface {
	db.Store
	WalletHistory(ctx context.Context, arg db.GetTransactionWalletByidAndUserIDParams, historyType string) (pathCSV string, err error)
}

type service struct {
	store   db.Store
	queries *db.Queries
	sqlDB   *sql.DB
}

func New(sqlDB *sql.DB) Service {
	return &service{
		store:   db.NewStore(sqlDB),
		queries: db.New(sqlDB),
		sqlDB:   sqlDB,
	}
}

func GenerateCSVWalletHistory(directory string, transactions []db.Transaction, transfer []db.Transfer) (trxFilename string, tfFilename string, err error) {
	trxFilename = fmt.Sprintf("%s/history_transactions_wallet_%d.csv", directory, transactions[0].WalletID)
	tfFilename = fmt.Sprintf("%s/history_transfer_wallet_%d.csv", directory, transactions[0].WalletID)

	// Create the directory if it doesn't exist
	err = os.MkdirAll(directory, os.ModePerm)

	if err != nil {
		return trxFilename, tfFilename, err
	}
	// Create a new file for transactions
	filetrx, err := os.Create(trxFilename)
	if err != nil {
		return trxFilename, tfFilename, err
	}
	defer filetrx.Close()

	writer := csv.NewWriter(filetrx)
	trxheader := []string{
		"Wallet ID",
		"Amount",
		"Transaction Type",
		"Transaction Date",
		"Description",
	}
	if err := writer.Write(trxheader); err != nil {
		return trxFilename, tfFilename, err
	}
	for _, trx := range transactions {
		row := []string{
			fmt.Sprintf("%d", trx.WalletID),
			fmt.Sprintf("%.2f", trx.Amount),
			trx.TransactionType,
			fmt.Sprintf("%v", trx.TransactionDate),
			trx.Description,
		}
		if err := writer.Write(row); err != nil {
			return trxFilename, tfFilename, err
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return trxFilename, tfFilename, err
	}

	// Create a new file for transactions
	filetf, err := os.Create(tfFilename)
	if err != nil {
		return trxFilename, tfFilename, err
	}
	defer filetrx.Close()

	writer = csv.NewWriter(filetf)
	tfheader := []string{
		"Wallet ID",
		"To Wallet ID",
		"Amount",
		"Transaction Date",
		"Description",
	}
	if err := writer.Write(tfheader); err != nil {
		return trxFilename, tfFilename, err
	}
	for _, tf := range transfer {
		row := []string{
			fmt.Sprintf("%d", tf.FromWalletID),
			fmt.Sprintf("%d", tf.ToWalletID),
			fmt.Sprintf("%.2f", tf.Amount),
			fmt.Sprintf("%v", tf.TransferDate),
			tf.Description,
		}
		if err := writer.Write(row); err != nil {
			return trxFilename, tfFilename, err
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return trxFilename, tfFilename, err
	}

	return trxFilename, tfFilename, err
}

// WalletHistory implements Service.
func (s *service) WalletHistory(ctx context.Context, arg db.GetTransactionWalletByidAndUserIDParams, historyType string) (pathCSV string, err error) {
	_, err = s.store.GetWalletById(ctx, int64(arg.WalletID))
	if err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet with id %d not found", arg.WalletID),
			Err: err,
		}
		err = cstErr
		return
	}

	_, err = s.store.GetUserById(ctx, int64(arg.UserID))
	if err != nil {
		cstErr := &utils.CustomError{
			Msg: fmt.Sprintf("wallet with id %d not found", arg.WalletID),
			Err: err,
		}
		err = cstErr
		return
	}

	result, err := s.store.WalletHistoryGenerateCSV(ctx, arg)
	if err != nil {
		cstErr := &utils.CustomError{
			Msg: "failed to get wallet history",
			Err: err,
		}
		err = cstErr
		return
	}

	// define tmp directory
	directory := utils.DIRECTORY_REPORTS

	trx, tf, err := GenerateCSVWalletHistory(directory, result.Transactions, result.Transfers)
	_, _ = trx, tf

	if err != nil {
		cstErr := &utils.CustomError{
			Msg: "failed to get generate csv file",
			Err: err,
		}
		err = cstErr
		return
	}

	if historyType == "transactions" {
		pathCSV = trx
	}
	if historyType == "transfers" {
		pathCSV = tf
	}
	return
}
