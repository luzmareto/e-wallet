// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Merchant struct {
	ID           int64          `json:"id"`
	MerchantName string         `json:"merchant_name"`
	Description  sql.NullString `json:"description"`
	Website      sql.NullString `json:"website"`
	Address      sql.NullString `json:"address"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiredAt    time.Time `json:"expired_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Topup struct {
	ID          int64     `json:"id"`
	UserID      int32     `json:"user_id"`
	WalletID    int32     `json:"wallet_id"`
	Amount      float64   `json:"amount"`
	TopupDate   time.Time `json:"topup_date"`
	Description string    `json:"description"`
}

type Transaction struct {
	ID              int64          `json:"id"`
	UserID          int32          `json:"user_id"`
	WalletID        int32          `json:"wallet_id"`
	Amount          float64        `json:"amount"`
	TransactionDate time.Time      `json:"transaction_date"`
	Description     sql.NullString `json:"description"`
}

type TransactionMerchant struct {
	TransactionID sql.NullInt32 `json:"transaction_id"`
	MerchantID    sql.NullInt32 `json:"merchant_id"`
}

type Transfer struct {
	ID           int64          `json:"id"`
	FromWalletID int32          `json:"from_wallet_id"`
	ToWalletID   int32          `json:"to_wallet_id"`
	Amount       float64        `json:"amount"`
	TransferDate time.Time      `json:"transfer_date"`
	Description  sql.NullString `json:"description"`
}

type User struct {
	ID               int64     `json:"id"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phone_number"`
	RegistrationDate time.Time `json:"registration_date"`
}

type Wallet struct {
	ID       int64   `json:"id"`
	UserID   int32   `json:"user_id"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

type Withdrawal struct {
	ID             int64     `json:"id"`
	UserID         int32     `json:"user_id"`
	WalletID       int32     `json:"wallet_id"`
	Amount         float64   `json:"amount"`
	WithdrawalDate time.Time `json:"withdrawal_date"`
	Description    string    `json:"description"`
}
