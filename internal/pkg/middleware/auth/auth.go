package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/GodYao1995/Goooooo/internal/admin/types"
	"github.com/GodYao1995/Goooooo/internal/domain"
	"github.com/GodYao1995/Goooooo/internal/pkg/errno"
	"github.com/GodYao1995/Goooooo/internal/pkg/session"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(store *session.RedisStore) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := types.CommonResponse{Code: 0}
		var user domain.UserSession
		session, err := store.New(ctx.Request, "SESSIONID")
		if errors.Is(err, errno.ErrorRedisEmpty) || errors.Is(err, http.ErrNoCookie) {
			resp.Message = errno.ErrorUserNotLogin.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}
		err = json.Unmarshal(session.Values["user"].([]byte), &user)
		if err != nil {
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
