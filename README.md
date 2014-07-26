# 4. Subcommand Flags

In the last chapter we saw how Cookoo's CLI runner can support
subcommands. We also saw how it parses command line flags. In this
section, we'll see how to add subcommand-specific flags.

## Hello Again

The only big difference in this version of `main.go` is this vastly
expanded version of the "hello" route:

```go
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
```

For readability, we've indented to show the structure more clearly.

The "hello" route now does three things in order:

1. Shift the arguments
2. Parse the subcommand flags
3. Run the Printf command

Step 1 might seem a little strange, so let's start with that.

## The Args

When we start up a command line client, we get the arguments from
`os.Args`. Let's say that we run the program like this:

```
$ go run main.go -a Matt hello -s Hi
```

This will generate an `os.Args` array that looks like this:

```go
[]string{ "-a", "Matt" "hello" "-s", "Hi" }
```

When the program first runs, we scan through the arguments and extract
the global flags. Everything else gets put into the context as
"runner.Args".

So effectively, we're doing something like this:


```go
cxt.Put("runner.Args", []string{ "hello" , "-s", "Hi" }
```

The "-a" and "Matt" have been interpreted already.

When the "hello" command runs, we now want it to parse out its flags.
But if it sees "hello" as the first value, it will assume that there are
no flags.

So we have to do this first:

```go
		Does(cli.ShiftArgs, "cmd").
			Using("args").WithDefault("runner.Args").
			Using("n").WithDefault(1).
```

This says "Shift runner.Args over n=1 places and store the shifted value
in 'cmd'".

Effectively, that will set "cxt:cmd" to "hello" and "cxt:runner.Args" to
`[]string{ "-s", "Hi" }`.

And *NOW* we can parse the subcommand arguments:

```
		Does(cli.ParseArgs, "extras").
			Using("flagset").WithDefault(helloFlags).
			Using("args").From("cxt:runner.Args").
```

This will use `helloFlags` to extract the subcommand-specific flags.

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
clear how the argument parsing works in Cookoo.

In the next chapter, we'll show a little trick for debugging argument
parsing.
