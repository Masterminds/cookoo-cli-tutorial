# 5. Subcommand Flags

In the last chapter we saw how Cookoo's CLI runner can support
subcommands. We also saw how it parses command line flags. In this
section, we'll see how to add subcommand-specific flags.

## Hello Again

The only big difference in this version of `main.go` is new version of
the "hello" route:

```go
	reg.Route("hello", "A Hello World route").
		Does(cli.ParseArgs, "extras").
			Using("flagset").WithDefault(helloFlags).
			Using("args").From("cxt:runner.Args").
      Using("subcommand").WithDefault(true).
		Does(fmt.Printf, "_").
			Using("format").WithDefault("%s %s!\n").
			Using("0").From("cxt:s").
			Using("1").From("cxt:a")
```

For readability, we've indented to show the structure more clearly.

The "hello" route now does two things in order:

1. Parse the subcommand flags
2. Run the Printf command

## Parsing Command Line Arguments

The first step in our two-step route is to parse the arguments for our
subcommand:

```
		Does(cli.ParseArgs, "extras").
			Using("flagset").WithDefault(helloFlags).
			Using("args").From("cxt:runner.Args").
      Using("subcommand").WithDefault(true).
```

This will use `helloFlags` to extract the subcommand-specific flags.
Wiht a glance at the code, we can see that we have defined two flagsets:

```go
	flags := flag.NewFlagSet("global", flag.PanicOnError)
	flags.Bool("h", false, "Show help text")
	flags.String("a", "World", "A string to place after 'Hello'")

	helloFlags := flag.NewFlagSet("hello", flag.PanicOnError)
	helloFlags.String("s", "Hello", "Alternate salutation")

```

The first set is the now-familar global flags. The second set,
`helloFlags` is our command-specific flags for the `hello` command.

We pass this second FlagSet into our `cli.ParseArgs` call above.

**Important:** When working with subcommands and `cli.ParseArgs`, make
sure to set `subcommand` is set to `true`. This tells the argument
parser where to start parsing. Check out the next chapter for detailed
information on this.

And now when we run `fmt.Printf` we pass it two parameters:

```
		Does(fmt.Printf, "_").
			Using("format").WithDefault("%s %s!\n").
			Using("0").From("cxt:s"). // The -s subcommand flag
			Using("1").From("cxt:a")  // The -a global flag
```

## In Action

Put all this together, and our commandline works like this:

```
$ go run main.go hello
Hello World!

$ go run main.go hello -s Hi
Hi World!

$ go run main.go -a You hello
Hello You!

$ go run main.go -a You hello -s Hi
Hi You!
```

While this isn't exactly the most user-friendly experience, it should be
clear how the argument parsing works in Cookoo. (For a real app, we'd
probably move the `-a` flag to the subcommand arguments.)

In the next chapter, we'll show a little trick for debugging argument
parsing. And there we'll also go into detail about how argument parsing
works.
