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

package routes

import (
	"fmt"
	"mystly/internal/core"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(server *core.Server) {
	// Register Handlers...
	NewTestHandler("/test", server).RegisterRoutes()

	// Add frontend...
	server.Router.Use(static.Serve("/", static.LocalFile("../asha/build", true)))
	server.Router.NoRoute(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path

		if (path == "" || path == "/") {
			ctx.File("../asha/build/index.html")
		} else {
			ctx.File(fmt.Sprintf("../asha/build/%s.html", path))
		}
	})
}


type IHandler interface {
	RegisterRoutes()
}
