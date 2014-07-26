package main

import (
	"github.com/Masterminds/cookoo"
	"github.com/Masterminds/cookoo/cli"
	"github.com/Masterminds/cookoo/fmt"

	"flag"
)

const (
	Summary = "A Hello World program"
	Description = `This program writes Hello World to standard output.

With the -a flag, the second word can be replaced by an artitary string.
`
)

func main() {
	reg, router, cxt := cookoo.Cookoo()

	flags := flag.NewFlagSet("global", flag.PanicOnError)
	flags.Bool("h", false, "Show help text")
	flags.String("a", "World", "A string to place after 'Hello'")

	reg.Route("hello", "A Hello World route").
		Does(fmt.Printf, "_").
		Using("format").WithDefault("Hello %s!\n").
		Using("0").From("cxt:a")

	cli.New(reg, router, cxt).Help(Summary, Description, flags).Run("hello")
}
