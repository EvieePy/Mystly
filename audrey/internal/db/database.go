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

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Database struct {
	Log  *zap.SugaredLogger
	Pool *pgxpool.Pool
}

func NewDatabase(log *zap.SugaredLogger) *Database {
	database := &Database{}
	database.Log = log

	return database
}

func (d *Database) Connect(dsn string) {
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		d.Log.Fatalf("Unable to connect to database with DSN: %s. %s", dsn, err)
	}

	d.Pool = pool
	d.Log.Infoln("Successfully connected to Database pool.")
}

func (d *Database) Close() {
	defer d.Pool.Close()
}
