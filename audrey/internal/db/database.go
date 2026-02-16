// Copyright 2026 Evie. P.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Database struct {
	Log  *zap.SugaredLogger
	Pool *pgxpool.Pool
}

func NewDatabase(log *zap.SugaredLogger) *Database {
	database := &Database{Log: log}
	return database
}

func (d *Database) Connect(dsn string) {
	pool, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		d.Log.Fatalf("Unable to connect to database with DSN: %s. %s", dsn, err)
	}

	d.Pool = pool
	d.Log.Infoln("Successfully connected to Database pool.")
	d.setup()
}

func (d *Database) Close() {
	defer d.Pool.Close()
}

func (d *Database) setup() {
	conn, err := d.Pool.Acquire(context.Background())
	if err != nil {
		d.Log.Fatalf("Unable to setup scripts for Database. %s", err)
	}

	defer conn.Release()

	cwd, err := os.Getwd()
	if err != nil {
		d.Log.Fatalf("Unable to setup scripts for Database. Unable to locate CWD. %s", err)
	}

	sql, err := os.ReadFile(filepath.Clean(fmt.Sprintf("%s/schema.sql", cwd)))
	if err != nil {
		d.Log.Fatalf("Unable to setup scripts for Database. Cannot read schema.sql. %s", err)
	}

	_, err = conn.Exec(context.Background(), string(sql))
	if err != nil {
		d.Log.Fatalf("Unable to setup scripts for Database. Cannot read schema.sql. %s", err)
	}

	d.Log.Info("Successfully setup Database.")
}

func (d *Database) AddUser(userId string, username string, key string) error {
	conn, err := d.Pool.Acquire(context.Background())

	if err != nil {
		return err
	}

	query := "INSERT INTO users(user_id, username, key) VALUES($1, $2, $3)"
	_, err = conn.Exec(context.Background(), query, userId, username, key)

	return err
}

func (d *Database) FetchUserByName(username string) (*UserModel, error) {
	conn, err := d.Pool.Acquire(context.Background())

	if err != nil {
		return nil, err
	}

	query := "SELECT * FROM users WHERE username = $1"
	rows, err := conn.Query(context.Background(), query, username)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[UserModel])
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *Database) FetchUsers() ([]*UserModel, error) {
	conn, err := d.Pool.Acquire(context.Background())

	if err != nil {
		return nil, err
	}

	query := "SELECT * FROM users"
	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[UserModel])
	if err != nil {
		return nil, err
	}

	return users, nil
}
