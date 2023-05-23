CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" VARCHAR(255) UNIQUE NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "phone_number" VARCHAR(20) NOT NULL,
  "registration_date" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "wallets" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "balance" NUMERIC(10, 2) NOT NULL,
  "currency" VARCHAR(10) NOT NULL
);

CREATE TABLE "transactions" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "wallet_id" INT NOT NULL,
  "amount" NUMERIC(10, 2) NOT NULL,
  "transaction_date" timestamptz NOT NULL DEFAULT 'NOW()',
  "description" VARCHAR(255) NOT NULL
);

CREATE TABLE "transfers" (
  "id" BIGSERIAL PRIMARY KEY,
  "from_wallet_id" INT NOT NULL,
  "to_wallet_id" INT NOT NULL,
  "amount" NUMERIC(10, 2) NOT NULL,
  "transfer_date" timestamptz NOT NULL DEFAULT 'NOW()',
  "description" VARCHAR(255) NOT NULL
);

CREATE TABLE "topups" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "wallet_id" INT NOT NULL,
  "amount" NUMERIC(10, 2) NOT NULL,
  "topup_date" timestamptz NOT NULL DEFAULT 'NOW()',
  "description" VARCHAR(255)
);

CREATE TABLE "withdrawals" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "wallet_id" INT NOT NULL,
  "amount" NUMERIC(10, 2) NOT NULL,
  "withdrawal_date" timestamptz NOT NULL DEFAULT 'NOW()',
  "description" VARCHAR(255)
);

CREATE TABLE "merchants" (
  "id" BIGSERIAL PRIMARY KEY,
  "merchant_name" VARCHAR(255) NOT NULL,
  "description" VARCHAR(255),
  "website" VARCHAR(255),
  "address" VARCHAR(255)
);

CREATE TABLE "transaction_merchants" (
  "transaction_id" INT,
  "merchant_id" INT
);