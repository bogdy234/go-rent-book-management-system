package config

import (
	"os"

	"github.com/gin-contrib/sessions/cookie"
)

var SessionStore cookie.Store

func InitSessionStore() {
	bytesSecret := []byte(os.Getenv("SESSION_SECRET"))
	SessionStore = cookie.NewStore(bytesSecret)
}
