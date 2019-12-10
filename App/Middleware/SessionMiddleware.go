package Middleware

import (
	"github.com/goKLC/goKLC"
)

type sessionMiddleware struct {
	goKLC.Middleware
	cookie *goKLC.Cookie
}

var SessionMiddleware = &sessionMiddleware{}

func (sm *sessionMiddleware) Handle(request *goKLC.Request) *goKLC.Response {
	cookie := goKLC.NewCookie()
	cookieName := app.Config().Get("SessionName", "goKLCSession").(string)
	ok, _ := cookie.Get(request, cookieName)

	if ok {

		return nil
	}

	cookie.Duration = app.Config().Get("SessionDuration", 60*60*60).(int)
	cookie.Name = cookieName
	cookie.Value = app.GetSessionKey()

	sm.cookie = cookie

	return nil
}

func (sm *sessionMiddleware) Terminate(response *goKLC.Response) {

	if sm.cookie != nil {
		response.AddCookie(sm.cookie)
	}
}
