package Middleware

import (
	"github.com/goKLC/goKLC"
	"github.com/goKLC/goResponse"
)

type sessionMiddleware struct {
	goKLC.Middleware
	cookie goKLC.Cookie
}

var SessionMiddleware = &sessionMiddleware{}

func (sm *sessionMiddleware) Handle(request *goKLC.Request) goKLC.Response {
	cookie := goResponse.NewCookie()
	cookieName := app.Config().Get("SessionName", "goKLCSession").(string)
	ok := request.GetCookie(cookieName)

	if ok != nil {

		return nil
	}

	value := app.GetSessionKey()
	duration := app.Config().Get("SessionDuration", 60*60*60).(int)
	cookie.Create(cookieName, value, duration, "/")

	sm.cookie = cookie

	return nil
}

func (sm *sessionMiddleware) Terminate(response goKLC.Response) {

	if sm.cookie != nil {
		response.WithCookie(sm.cookie)
	}
}
