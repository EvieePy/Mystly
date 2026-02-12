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
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IMiddleware interface {
	Handler() gin.HandlerFunc
}

// Auth Middleware
type AuthMiddleware struct {
	Server *Server
}

func NewAuthMiddleware(server *Server) IMiddleware {
	return &AuthMiddleware{Server: server}
}

func (m *AuthMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isBearer := false
		session := sessions.Default(ctx)
		tok := session.Get("_sess_key")

		if tok == nil {
			bearer := ctx.GetHeader("Authorization")

			var ok bool
			_, tok, ok = strings.Cut(bearer, "Bearer ")

			if !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			isBearer = true
		}

		tokenString, ok := tok.(string)
		if !ok || tokenString == "" {
			session.Delete("_sess_key")
			session.Save()

			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userId, err := m.Server.ParseWebToken(tokenString)
		if err != nil {
			session.Delete("_sess_key")
			session.Save()

			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check Auth token validity; not session...
		if isBearer {

		} else {
			
		}

		// TODO: Fetch user from Database...
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
