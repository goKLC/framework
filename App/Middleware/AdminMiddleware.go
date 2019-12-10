package Middleware

import (
	"fmt"
	"github.com/goKLC/goKLC"
)

type adminMiddleware struct {
	goKLC.Middleware
}

var AdminMiddleware = adminMiddleware{}

func (a adminMiddleware) Handle(r *goKLC.Request) *goKLC.Response {

	fmt.Println("adminMiddleware@handle")

	return nil
}
