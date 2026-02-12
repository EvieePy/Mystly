package routes

import (
	"mystly/internal/core"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type TestHandler struct {
	Server    *core.Server
	Path string
}

func NewTestHandler(path string, server *core.Server) IHandler {
	return &TestHandler{Path: path, Server: server}
}

func (h *TestHandler) RegisterRoutes() {
	group := h.Server.Router.Group(h.Path)

	auth := core.NewAuthMiddleware(h.Server)

	group.GET("/auth", h.getAuthed())
	group.Use(auth.Handler())
	group.GET("/check", h.checkAuth())
}

func (h *TestHandler) checkAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		val, exists := ctx.Get("userId")

		if exists {
			ctx.String(200, "You are authed as %s", val)
		} else {
			ctx.String(200, `You are not authed!`)
		}
	}
}

func (h *TestHandler) getAuthed() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, exists := ctx.Get("userId")

		if !exists {
			session := sessions.Default(ctx)

			tokenString, err := h.Server.NewWebToken("mystyTest123!")
			if err != nil {
				h.Server.Log.Error(err)
				ctx.String(500, "Bad")
				return
			}

			session.Set("_sess_key", tokenString)
			session.Save()
		}

		ctx.String(200, "kk")
	}
}
