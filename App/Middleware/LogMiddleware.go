package Middleware

import (
	"fmt"
	"github.com/goKLC/goKLC"
)

type logMiddleware struct {
	goKLC.Middleware
}

var LogMiddleware = logMiddleware{}

func (lm logMiddleware) Handle(request *goKLC.Request) goKLC.Response {

	fmt.Println("logMiddleware@handle")

	return nil
}

func (lm logMiddleware) Terminate(response goKLC.Response) {

	fmt.Println("logMiddleware@terminate")
}
