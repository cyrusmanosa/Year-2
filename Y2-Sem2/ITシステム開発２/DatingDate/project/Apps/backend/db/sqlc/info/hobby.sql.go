// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: hobby.sql

package db

import (
	"context"

)

const createHobby = `-- name: CreateHobby :one
INSERT INTO hobby (
    user_id,
    era,
    city,
    gender,
    speaklanguage,
    find_type,
    experience
) VALUES (
    $1,$2,$3,$4,$5,$6,$7
) RETURNING user_id, era, city, gender, speaklanguage, find_type, experience, info_changed_at
`

type CreateHobbyParams struct {
	UserID        int32       `json:"user_id"`
	Era           int32 `json:"era"`
	City          []string    `json:"city"`
	Gender        string `json:"gender"`
	Speaklanguage []string    `json:"speaklanguage"`
	FindType      string `json:"find_type"`
	Experience    int32 `json:"experience"`
}

func (q *Queries) CreateHobby(ctx context.Context, arg CreateHobbyParams) (Hobby, error) {
	row := q.db.QueryRow(ctx, createHobby,
		arg.UserID,
		arg.Era,
		arg.City,
		arg.Gender,
		arg.Speaklanguage,
		arg.FindType,
		arg.Experience,
	)
	var i Hobby
	err := row.Scan(
		&i.UserID,
		&i.Era,
		&i.City,
		&i.Gender,
		&i.Speaklanguage,
		&i.FindType,
		&i.Experience,
		&i.InfoChangedAt,
	)
	return i, err
}

const deleteHobby = `-- name: DeleteHobby :exec
DELETE FROM hobby 
WHERE user_id = $1
`

func (q *Queries) DeleteHobby(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deleteHobby, userID)
	return err
}

const getHobby = `-- name: GetHobby :one
SELECT user_id, era, city, gender, speaklanguage, find_type, experience, info_changed_at FROM hobby
WHERE user_id = $1
`

func (q *Queries) GetHobby(ctx context.Context, userID int32) (Hobby, error) {
	row := q.db.QueryRow(ctx, getHobby, userID)
	var i Hobby
	err := row.Scan(
		&i.UserID,
		&i.Era,
		&i.City,
		&i.Gender,
		&i.Speaklanguage,
		&i.FindType,
		&i.Experience,
		&i.InfoChangedAt,
	)
	return i, err
}

const listHobby = `-- name: ListHobby :many
SELECT user_id, era, city, gender, speaklanguage, find_type, experience, info_changed_at FROM hobby
ORDER BY user_id
`

func (q *Queries) ListHobby(ctx context.Context) ([]Hobby, error) {
	rows, err := q.db.Query(ctx, listHobby)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Hobby{}
	for rows.Next() {
		var i Hobby
		if err := rows.Scan(
			&i.UserID,
			&i.Era,
			&i.City,
			&i.Gender,
			&i.Speaklanguage,
			&i.FindType,
			&i.Experience,
			&i.InfoChangedAt,
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

const updateHobby = `-- name: UpdateHobby :one
UPDATE hobby
SET era = $2,
    city = $3,
    gender = $4,
    speaklanguage = $5,
    find_type = $6,
    experience = $7
WHERE user_id = $1
RETURNING user_id, era, city, gender, speaklanguage, find_type, experience, info_changed_at
`

type UpdateHobbyParams struct {
	UserID        int32       `json:"user_id"`
	Era           int32 `json:"era"`
	City          []string    `json:"city"`
	Gender        string `json:"gender"`
	Speaklanguage []string    `json:"speaklanguage"`
	FindType      string `json:"find_type"`
	Experience    int32 `json:"experience"`
}

func (q *Queries) UpdateHobby(ctx context.Context, arg UpdateHobbyParams) (Hobby, error) {
	row := q.db.QueryRow(ctx, updateHobby,
		arg.UserID,
		arg.Era,
		arg.City,
		arg.Gender,
		arg.Speaklanguage,
		arg.FindType,
		arg.Experience,
	)
	var i Hobby
	err := row.Scan(
		&i.UserID,
		&i.Era,
		&i.City,
		&i.Gender,
		&i.Speaklanguage,
		&i.FindType,
		&i.Experience,
		&i.InfoChangedAt,
	)
	return i, err
}
