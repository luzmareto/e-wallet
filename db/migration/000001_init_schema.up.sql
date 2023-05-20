CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" VARCHAR(255) UNIQUE NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "phone_number" VARCHAR(20) NOT NULL,
  "registration_date" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "wallets" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" INT,
  "balance" DECIMAL(10,2) NOT NULL,
  "currency" VARCHAR(10) NOT NULL
);

CREATE TABLE "transactions" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" INT,
  "wallet_id" INT,
  "amount" DECIMAL(10,2) NOT NULL,
  "transaction_date" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()',
  "description" VARCHAR(255)
);

CREATE TABLE "transfers" (
  "id" BIGSERIAL PRIMARY KEY,
  "from_wallet_id" INT,
  "to_wallet_id" INT,
  "amount" DECIMAL(10,2) NOT NULL,
  "transfer_date" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()',
  "description" VARCHAR(255)
);

CREATE TABLE "topups" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" INT,
  "wallet_id" INT,
  "amount" DECIMAL(10,2) NOT NULL,
  "topup_date" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()',
  "description" VARCHAR(255)
);

CREATE TABLE "withdrawals" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" INT,
  "wallet_id" INT,
  "amount" DECIMAL(10,2) NOT NULL,
  "withdrawal_date" TIMESTAMPTZ NOT NULL DEFAULT 'NOW()',
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

ALTER TABLE "wallets" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "topups" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "topups" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "withdrawals" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "withdrawals" ADD FOREIGN KEY ("wallet_id") REFERENCES "wallets" ("id");

ALTER TABLE "transaction_merchants" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id");

ALTER TABLE "transaction_merchants" ADD FOREIGN KEY ("merchant_id") REFERENCES "merchants" ("id");