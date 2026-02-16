package routes

import (
	"crypto/subtle"
	"mystly/internal/core"

	"github.com/fernet/fernet-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserLoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UsersHandler struct {
	Server *core.Server
	Path   string
}

func NewUserHandler(path string, server *core.Server) IHandler {
	return &UsersHandler{Path: path, Server: server}
}

func (h *UsersHandler) RegisterRoutes() {
	group := h.Server.Router.Group(h.Path)
	// auth := core.NewAuthMiddleware(h.Server)

	// Unauthed...
	group.GET("users/checksetup", h.shouldSetup())
	group.POST("users/login", h.userLogin())
	group.POST("users/initial", h.addInitialUser())

	// Authed...
	// group.Use(auth.Handler())
}

func (f *UserLoginForm) Validate(ctx *gin.Context) bool {
	if len(f.Password) < 8 {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Password must be 8 or more characters."})
		return false
	}

	if len(f.Password) > 256 {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Password must not be more than 256 characters long."})
		return false
	}

	if len(f.Username) < 3 {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Username must be 3 or more characters long."})
		return false
	}

	if len(f.Username) > 56 {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Username must not be more than 56 characters long."})
		return false
	}

	if core.ContainsSpace(f.Password) {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Password must not contain spaces."})
		return false
	}

	if core.ContainsSpace(f.Username) {
		ctx.AbortWithStatusJSON(400, gin.H{"error": "Username must not contain spaces."})
		return false
	}

	return true
}

func (h *UsersHandler) shouldSetup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := h.Server.Database.FetchUsers()

		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		if len(users) == 0 {
			ctx.Data(204, "text/plain", []byte{})
			return
		}

		ctx.Data(409, "text/plain", []byte{})
	}
}

func (h *UsersHandler) userLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		form := UserLoginForm{}
		err := ctx.Bind(&form)

		if err != nil {
			ctx.AbortWithStatus(400)
			return
		}

		user, err := h.Server.Database.FetchUserByName(form.Username)
		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		msg := fernet.VerifyAndDecrypt([]byte(user.Key), 0, []*fernet.Key{h.Server.Config.FKey})
		if msg == nil {
			ctx.AbortWithStatus(403)
			return
		}

		if subtle.ConstantTimeCompare(msg, []byte(form.Password)) != 1 {
			ctx.AbortWithStatus(403)
			return
		}

		session := sessions.Default(ctx)
		session.Clear()

		tokenString, err := h.Server.NewWebToken(user.UserId.String())

		if err != nil {
			h.Server.Log.Error(err)
			ctx.AbortWithError(500, err)
			return
		}

		session.Set("_sess_key", tokenString)
		session.Save()

		ctx.Data(204, "text/plain", []byte{})
	}
}

func (h *UsersHandler) addInitialUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := h.Server.Database.FetchUsers()

		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		if len(users) != 0 {
			ctx.AbortWithStatus(403)
			return
		}

		form := UserLoginForm{}
		err = ctx.Bind(&form)

		if err != nil {
			ctx.AbortWithStatus(400)
			return
		}

		if !form.Validate(ctx) {
			return
		}

		tok, err := fernet.EncryptAndSign([]byte(form.Password), h.Server.Config.FKey)
		if err != nil {
			ctx.AbortWithStatus(500)
			return
		}

		userId := uuid.New()
		err = h.Server.Database.AddUser(userId.String(), form.Username, string(tok))

		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		session := sessions.Default(ctx)
		session.Clear()

		tokenString, err := h.Server.NewWebToken(userId.String())

		if err != nil {
			h.Server.Log.Error(err)
			ctx.AbortWithError(500, err)
			return
		}

		session.Set("_sess_key", tokenString)
		session.Save()

		ctx.Data(204, "text/plain", []byte{})
	}
}
