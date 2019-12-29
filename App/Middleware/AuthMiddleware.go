package Middleware

import (
	"fmt"
	"github.com/goKLC/goKLC"
)

type authMiddleware struct {
	goKLC.Middleware
}

var AuthMiddleware = authMiddleware{}

func (a authMiddleware) Handle(request *goKLC.Request) *goKLC.Response {

	auth := goKLC.GetApp().Auth()

	if !auth.Check(request) {

		return goKLC.NewResponse().Error("unauthorization action")
	}

	return nil
}

func (a authMiddleware) Terminate(response *goKLC.Response) {

	fmt.Println("authMiddleware@terminate")
}
