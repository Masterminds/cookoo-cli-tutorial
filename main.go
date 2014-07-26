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

Subcommands:

	hello: Say hello
	goodbye: Say goodbye
`
)

func main() {
	reg, router, cxt := cookoo.Cookoo()

	flags := flag.NewFlagSet("global", flag.PanicOnError)
	flags.Bool("h", false, "Show help text")
	flags.String("a", "World", "A string to place after 'Hello'")

	helloFlags := flag.NewFlagSet("hello", flag.PanicOnError)
	helloFlags.String("s", "Hello", "Alternate salutation")

	reg.Route("debug", "Show how to debug arguments").
		Does(fmt.Printf, "_").
			Using("format").WithDefault("First os.Args: %v\nrunner.Args: %v\n\n").
			Using("0").From("cxt:os.Args").
			Using("1").From("cxt:runner.Args").
		Does(cli.ShiftArgs, "cmd").
			Using("args").WithDefault("runner.Args").
			Using("n").WithDefault(1).
		Does(fmt.Printf, "_").
			Using("format").WithDefault("Second os.Args: %v\nrunner.Args: %v\n\n").
			Using("0").From("cxt:os.Args").
			Using("1").From("cxt:runner.Args").
		Does(cli.ParseArgs, "extras").
			Using("flagset").WithDefault(helloFlags).
			Using("args").From("cxt:runner.Args").
		Does(fmt.Printf, "_").
		Using("format").WithDefault("Third os.Args: %v\nrunner.Args: %v\nExtras: %v\n\n").
			Using("0").From("cxt:os.Args").
			Using("1").From("cxt:runner.Args").
			Using("2").From("cxt:extras")

	reg.Route("hello", "A Hello World route").
		Does(cli.ShiftArgs, "cmd").
			Using("args").WithDefault("runner.Args").
			Using("n").WithDefault(1).
		Does(cli.ParseArgs, "extras").
			Using("flagset").WithDefault(helloFlags).
			Using("args").From("cxt:runner.Args").
		Does(fmt.Printf, "_").
			Using("format").WithDefault("%s %s!\n").
			Using("0").From("cxt:s").
			Using("1").From("cxt:a")

	reg.Route("goodbye", "A Goodbye World route").
		Does(fmt.Printf, "_").
		Using("format").WithDefault("Goodbye %s!\n").
		Using("0").From("cxt:a")

	cli.New(reg, router, cxt).Help(Summary, Description, flags).RunSubcommand()
}
