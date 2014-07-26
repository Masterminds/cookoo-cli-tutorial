package main

import (
	"github.com/Masterminds/cookoo"
	"github.com/Masterminds/cookoo/cli"
	"github.com/Masterminds/cookoo/fmt"
)

func main() {
	reg, router, cxt := cookoo.Cookoo()

	reg.Route("hello", "A Hello World route").
		Does(fmt.Printf, "_").
		Using("format").WithDefault("Hello World!\n")

	cli.New(reg, router, cxt).Run("hello")
}
