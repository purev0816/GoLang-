// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
    owner, 
    balance,
    currency
) VALUES (
    $1, $2, $3
) RETURNING id, owner, balance, currency, created_at
`

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  string `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error) {
	row := q.queryRow(ctx, q.createAccountStmt, createAccount, arg.Owner, arg.Balance, arg.Currency)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccounts = `-- name: DeleteAccounts :one
DELETE FROM accounts WHERE id = $1 RETURNING id, owner, balance, currency, created_at
`

func (q *Queries) DeleteAccounts(ctx context.Context, id int64) (Accounts, error) {
	row := q.queryRow(ctx, q.deleteAccountsStmt, deleteAccounts, id)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, owner, balance, currency, created_at FROM accounts WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Accounts, error) {
	row := q.queryRow(ctx, q.getAccountStmt, getAccount, id)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, owner, balance, currency, created_at FROM accounts ORDER BY id LIMIT $1 OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Accounts, error) {
	rows, err := q.query(ctx, q.listAccountsStmt, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Accounts
	for rows.Next() {
		var i Accounts
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccounts = `-- name: UpdateAccounts :one
UPDATE accounts SET balance = $2 WHERE id = $1 RETURNING id, owner, balance, currency, created_at
`

type UpdateAccountsParams struct {
	ID      int64  `json:"id"`
	Balance string `json:"balance"`
}

func (q *Queries) UpdateAccounts(ctx context.Context, arg UpdateAccountsParams) (Accounts, error) {
	row := q.queryRow(ctx, q.updateAccountsStmt, updateAccounts, arg.ID, arg.Balance)
	var i Accounts
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}