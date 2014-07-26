# 6. Debugging Args

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
