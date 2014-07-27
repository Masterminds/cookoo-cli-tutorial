# 6. Debugging Args

This chapter takes a look at another way of doing what was done in the
last chapter. The goal of this chapter is to understand how parsing
commandline arguments works.

This chapter is *optional*. If you're not into command line parsing, you
can skip this for now and come back to it when it strikes your fancy.

## A Revised Hello Route

Take a look at the `main.go` in this chapter. It is a different method
for doing the same thing we did in the last chapter:

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

Here let's just look at a brief way to visualize the argument parsing
from the last chapter.

We added a new route to the main function:

```go
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
```

This route dumps the values of various versions of `os.Args`,
`runner.Args` and (in the last case) the `extras` left over after
parsing the subcommand.

Let's try running this:

```
go run main.go -a Matt debug -s "OH HAI" test
First os.Args: [-a Matt debug -s OH HAI test]
runner.Args: [debug -s OH HAI test]

Second os.Args: [-a Matt debug -s OH HAI test]
runner.Args: [-s OH HAI test]

Third os.Args: [-a Matt debug -s OH HAI test]
runner.Args: [-s OH HAI test]
Extras: [test]
```

The above shows how `os.Args` always contains the full argument list,
while `runner.Args` changes after arguments are parsed.

Finally, we can see how after we've parsed out `debug -s "OH HAI"` there
was still one leftover: `test`.

With some creativity, you can write Cookoo CLI apps that handle deep
subcommands, like `$ mycmd run more commands -a foo`.

But we're gonna leave that up to you.
