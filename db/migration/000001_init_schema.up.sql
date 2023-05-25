CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,
    "role" VARCHAR(10) NOT NULL DEFAULT 'user',
    "username" VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "phone_number" VARCHAR(20) NOT NULL,
    "id_card" VARCHAR(255) NOT NULL DEFAULT '',
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
    "transaction_type" VARCHAR(10) NOT NULL,
    "description" VARCHAR(255) NOT NULL
);

CREATE TABLE "transfers" (
    "id" BIGSERIAL PRIMARY KEY,
    "from_wallet_id" INT NOT NULL,
    "to_wallet_id" INT NOT NULL,
    "amount" NUMERIC(10, 2) NOT NULL,
    "transfer_date" timestamptz NOT NULL DEFAULT 'NOW()',
    "description" VARCHAR(255) NOT NULL DEFAULT ''
);

CREATE TABLE "topups" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" INT NOT NULL,
    "wallet_id" INT NOT NULL,
    "amount" NUMERIC(10, 2) NOT NULL,
    "topup_date" timestamptz NOT NULL DEFAULT 'NOW()',
    "description" VARCHAR(255) NOT NULL DEFAULT ''
);

CREATE TABLE "withdrawals" (
    "id" BIGSERIAL PRIMARY KEY,
    "user_id" INT NOT NULL,
    "wallet_id" INT NOT NULL,
    "amount" NUMERIC(10, 2) NOT NULL,
    "withdrawal_date" timestamptz NOT NULL DEFAULT 'NOW()',
    "description" VARCHAR(255) NOT NULL DEFAULT ''
);

CREATE TABLE "merchants" (
    "id" BIGSERIAL PRIMARY KEY,
    "merchant_name" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL DEFAULT '',
    "website" VARCHAR(255) NOT NULL DEFAULT '',
    "address" VARCHAR(255) NOT NULL DEFAULT '',
    "balance" NUMERIC(10, 2) NOT NULL DEFAULT 0
);

CREATE TABLE "transaction_merchants" (
    "transaction_id" INT,
    "merchant_id" INT
);

CREATE TABLE "sessions" (
    "id" uuid PRIMARY KEY,
    "username" varchar NOT NULL,
    "refresh_token" varchar NOT NULL,
    "user_agent" varchar NOT NULL,
    "client_ip" varchar NOT NULL,
    "is_blocked" boolean NOT NULL DEFAULT false,
    "expired_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

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