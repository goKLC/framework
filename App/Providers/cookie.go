package Providers

import (
	"github.com/goKLC/goKLC"
	"github.com/goKLC/goResponse"
)

func init() {
	var app = goKLC.GetApp()
	var cookie = goResponse.NewCookie()

	app.SetCookie(cookie)
}
