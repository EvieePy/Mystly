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
package core

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Port          int
	SessionSecret string `envconfig:"SESSION_SECRET"`
	JWTSecret     string `envconfig:"JWT_SECRET"`
	PostgresDSN   string `envconfig:"POSTGRES_DSN"`
}

func NewConfig(logger *zap.SugaredLogger) *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Unable to load .env config.")
	}

	cfg := &Config{}
	err = envconfig.Process("", cfg)

	if err != nil {
		logger.Fatal("Unable to load .env config.")
	}

	return cfg
}
