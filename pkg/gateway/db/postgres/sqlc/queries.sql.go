// Code generated by sqlc. DO NOT EDIT.
// source: queries.sql

package sqlc

import (
	"context"
)

const saveTransaction = `-- name: SaveTransaction :one
INSERT INTO transactions (account_id, operation_type, amount)
VALUES ($1, $2::int, $3::int)
RETURNING id
`

type SaveTransactionParams struct {
	AccountID     string `json:"account_id"`
	OperationType int32  `json:"operation_type"`
	Amount        int32  `json:"amount"`
}

func (q *Queries) SaveTransaction(ctx context.Context, arg SaveTransactionParams) (string, error) {
	row := q.db.QueryRow(ctx, saveTransaction, arg.AccountID, arg.OperationType, arg.Amount)
	var id string
	err := row.Scan(&id)
	return id, err
}
