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

	fmt.Println("authMiddleware@handle")

	return nil
}

func (a authMiddleware) Terminate(response *goKLC.Response) {

	fmt.Println("authMiddleware@terminate")
}
