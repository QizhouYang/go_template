package constant

import (
	"time"

	"github.com/kataras/iris/v12/sessions"
)

const (
	SessionUserKey         = "user"
	CookieNameForSessionID = "ksession"

	AuthMethodSession = "session"
	AuthMethodJWT     = "jwt"
)

var Sess = sessions.New(sessions.Config{Cookie: CookieNameForSessionID, Expires: 12 * time.Hour})
