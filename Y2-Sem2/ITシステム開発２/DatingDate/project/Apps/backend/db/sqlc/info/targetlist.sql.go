// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: targetlist.sql

package db

import (
	"context"

)

const allTargetUserList = `-- name: AllTargetUserList :many
SELECT user_id, target_1_id, "t1_Type", target_2_id, "t2_Type", target_3_id, "t3_Type", updated_at FROM targetlist
ORDER BY user_id
`

func (q *Queries) AllTargetUserList(ctx context.Context) ([]Targetlist, error) {
	rows, err := q.db.Query(ctx, allTargetUserList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Targetlist{}
	for rows.Next() {
		var i Targetlist
		if err := rows.Scan(
			&i.UserID,
			&i.Target1ID,
			&i.T1Type,
			&i.Target2ID,
			&i.T2Type,
			&i.Target3ID,
			&i.T3Type,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteTargetList = `-- name: DeleteTargetList :exec
DELETE FROM targetlist 
WHERE user_id = $1
`

func (q *Queries) DeleteTargetList(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deleteTargetList, userID)
	return err
}

const getTargetUserList = `-- name: GetTargetUserList :one
SELECT user_id, target_1_id, "t1_Type", target_2_id, "t2_Type", target_3_id, "t3_Type", updated_at FROM targetlist
WHERE user_id = $1
`

func (q *Queries) GetTargetUserList(ctx context.Context, userID int32) (Targetlist, error) {
	row := q.db.QueryRow(ctx, getTargetUserList, userID)
	var i Targetlist
	err := row.Scan(
		&i.UserID,
		&i.Target1ID,
		&i.T1Type,
		&i.Target2ID,
		&i.T2Type,
		&i.Target3ID,
		&i.T3Type,
		&i.UpdatedAt,
	)
	return i, err
}

const targetUserList = `-- name: TargetUserList :one
INSERT INTO targetlist (
    user_id,
    target_1_id,
    "t1_Type",
    target_2_id,
    "t2_Type",
    target_3_id,
    "t3_Type"
) VALUES (
    $1,$2,$3,$4,$5,$6,$7
) RETURNING user_id, target_1_id, "t1_Type", target_2_id, "t2_Type", target_3_id, "t3_Type", updated_at
`

type TargetUserListParams struct {
	UserID    int32       `json:"user_id"`
	Target1ID int32       `json:"target_1_id"`
	T1Type    string      `json:"t1_Type"`
	Target2ID int32 `json:"target_2_id"`
	T2Type    string `json:"t2_Type"`
	Target3ID int32 `json:"target_3_id"`
	T3Type    string `json:"t3_Type"`
}

func (q *Queries) TargetUserList(ctx context.Context, arg TargetUserListParams) (Targetlist, error) {
	row := q.db.QueryRow(ctx, targetUserList,
		arg.UserID,
		arg.Target1ID,
		arg.T1Type,
		arg.Target2ID,
		arg.T2Type,
		arg.Target3ID,
		arg.T3Type,
	)
	var i Targetlist
	err := row.Scan(
		&i.UserID,
		&i.Target1ID,
		&i.T1Type,
		&i.Target2ID,
		&i.T2Type,
		&i.Target3ID,
		&i.T3Type,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTargetList = `-- name: UpdateTargetList :one
UPDATE targetlist
SET target_1_id = $2,
    "t1_Type" = $3,
    target_2_id = $4,
    "t2_Type" = $5,
    target_3_id = $6,
    "t3_Type" = $7
WHERE user_id = $1 
RETURNING user_id, target_1_id, "t1_Type", target_2_id, "t2_Type", target_3_id, "t3_Type", updated_at
`

type UpdateTargetListParams struct {
	UserID    int32       `json:"user_id"`
	Target1ID int32       `json:"target_1_id"`
	T1Type    string      `json:"t1_Type"`
	Target2ID int32 `json:"target_2_id"`
	T2Type    string `json:"t2_Type"`
	Target3ID int32 `json:"target_3_id"`
	T3Type    string `json:"t3_Type"`
}

func (q *Queries) UpdateTargetList(ctx context.Context, arg UpdateTargetListParams) (Targetlist, error) {
	row := q.db.QueryRow(ctx, updateTargetList,
		arg.UserID,
		arg.Target1ID,
		arg.T1Type,
		arg.Target2ID,
		arg.T2Type,
		arg.Target3ID,
		arg.T3Type,
	)
	var i Targetlist
	err := row.Scan(
		&i.UserID,
		&i.Target1ID,
		&i.T1Type,
		&i.Target2ID,
		&i.T2Type,
		&i.Target3ID,
		&i.T3Type,
		&i.UpdatedAt,
	)
	return i, err
}
