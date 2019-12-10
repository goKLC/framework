package Controller

import (
	"fmt"
	"github.com/goKLC/goKLC"
)

type indexController goKLC.Controller

var IndexController = indexController{}

func (ic indexController) Index(r *goKLC.Request) *goKLC.Response {
	session := goKLC.NewSession(r)
	session.Set("foo", "bar")

	return goKLC.NewResponse().Ok(goKLC.NewView("index", goKLC.Context{}).Render())
}

func (ic indexController) User(r *goKLC.Request) *goKLC.Response {
	session := goKLC.NewSession(r)
	value := session.Get("foo", "baz")

	fmt.Println(value)

	return goKLC.NewResponse().Ok(fmt.Sprintf("%s : %s : %s", "indexcontroller@user", goKLC.GetRoute("admin.user"), r.GetParameter("user")))
}
func (ic indexController) Profile(r *goKLC.Request) *goKLC.Response {

	return goKLC.NewResponse().Ok("indexcontroller@profile")
}

func (ic indexController) Get(r *goKLC.Request) *goKLC.Response {
	session := goKLC.NewSession(r)
	value := session.Get("baz", "foo")

	fmt.Println(value)

	return goKLC.NewResponse().Ok("indexController@get")
}

func (ic indexController) Post(r *goKLC.Request) *goKLC.Response {

	return goKLC.NewResponse().Error("indexController@post")
}
