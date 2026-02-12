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
	"mystly/internal/db"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	Router   *gin.Engine
	Database *db.Database
	Log      *zap.SugaredLogger
	Config   *Config
}

func NewServer() *Server {
	server := &Server{}

	// Logger setup...
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	// Config setup...
	cfg := NewConfig(sugar)

	// Database setup...
	database := db.NewDatabase(sugar)

	// Router setup...
	router := gin.Default()
	router.SetTrustedProxies(nil)

	// Assign to server...
	server.Router = router
	server.Database = database
	server.Log = sugar
	server.Config = cfg

	// Setup Session Middleware...
	store := cookie.NewStore([]byte(cfg.SessionSecret))
	router.Use(sessions.Sessions("_sess_key", store))

	return server
}

func (s *Server) Run(addr string, dsn string) {
	// Start Database...
	s.Database.Connect(dsn)
	defer s.Database.Close()

	s.Router.Run(addr)
}
