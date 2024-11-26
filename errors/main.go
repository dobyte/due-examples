package main

import (
	"fmt"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/errors"
)

func main() {
	err := errors.NewError(codes.NewCode(404, "not found"))

	code := codes.Convert(err)

	fmt.Println(err)

	fmt.Println(code.Code())
	fmt.Println(code.Message())
}
