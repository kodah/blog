package service

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewCookieStore(secret string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(secret))

	return sessions.Sessions("user", store)
}

type SessionService interface {
	Get(key interface{}) interface{}
	Set(key interface{}, val interface{})
	Delete(key interface{})
	Clear()
	AddFlash(value interface{}, vars ...string)
	Flashes(vars ...string) []interface{}
	Options(sessions.Options)
	Save() error
}

func NewSessionService(ctx *gin.Context, close bool) SessionService {
	session := sessions.Default(ctx)

	opt := sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   3600 * 24,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	if close {
		opt.MaxAge = 0
	}

	session.Options(opt)

	return session
}
