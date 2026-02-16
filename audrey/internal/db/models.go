package db

import "github.com/jackc/pgx/v5/pgtype"

type UserModel struct {
	UserId    pgtype.UUID      `db:"user_id"`
	Username  string           `db:"username"`
	RoleId    pgtype.UUID      `db:"role_id"`
	CreatedAt pgtype.Timestamp `db:"created_at"`
	Key       string           `db:"key"`
}
