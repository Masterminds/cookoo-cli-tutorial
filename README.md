# 4. Subcommands

In the last chapter we saw the canonical pattern for writing simple
Cookoo CLIs. To wit, a standard Cookoo CLI does the following:

1. Creates a Cookoo instance
2. Configures a FlagSet of command line flags
3. Creates a Route
4. Runs the route with the pattern `cli.New().Help().Run()`.

But many modern command line clients use what are called subcommands.
Here are some examples:

```
$ git clone
$ git push
$ go build
$ go get
$ glide install
$ glide update
```

Above, we can see three examples of how a program may do more than one
thing. `git clone`, for example, clones a repository, while `git push`
pushes the local repository to a remote.

Cookoo provides support for this kind of command/subcommand delegation
model. In this chapter we will extend our earlier command to have two
subcommands.

## Two Routes

Take a look at this chapter's `main.go`. You'll notice that we now have
two routes:

```go
	reg.Route("hello", "A Hello World route").
		Does(fmt.Printf, "_").
		Using("format").WithDefault("Hello %s!\n").
		Using("0").From("cxt:a")

	reg.Route("goodbye", "A Goodbye World route").
		Does(fmt.Printf, "_").
		Using("format").WithDefault("Goodbye %s!\n").
		Using("0").From("cxt:a")
```

The 'hello' route is unchanged from the last chapter. And the new
'goodbye' route is almost identical to it. It just says Goodbye instead
of Hello.

Now that we have two routes, we want to be able to allow users to call
whichever of these two they please. Unsurprisingly, Cookoo does this for
us. All we need to do is make a minor adjustment to tell the Cookoo CLI
runner to use the *subcommand* mode instead of the normal static runner:

```go
	cli.New(reg, router, cxt).Help(Summary, Description, flags).RunSubcommand()
```

The only change to this line is the change from `Run("hello")` to
`RunSubcommand()`. The `RunSubcommand()` method tells the CLI runner to
figure out which route to run based on the commandline.

## Running Subcommands

Let's see what happens when we call the program different ways:

```
$ go run main.go       # Now prints help text
$ go run main.go -h    # Now also prints help text
$ go run main.go help  # Whoa! It's the help text again!
```

Uh... we got a little zealous about providing good help. But wait! We
can also call `hello` and `goodbye` now!

```
$ go run main.go hello
Hello World!
$ go run main.go goodbye
Goodbye World!
```

And we now have support for the `-a` flag on both `hello` and `goodbye`:

```
$ go run main.go -a Matt hello
Hello Matt!
$ go run main.go -a Matt goodbye
Goodbye Matt!
```

And wait, there's more!

```
go run main.go goodbye -a Matt
Goodbye World!
```

Huh? Why did that print `Goodbye World!` instead of `Goodbye Matt!`?

The reason is that Cookoo's CLI runner has a notion of *global* flags
and *subcommand* flags.

**Global flags** are shared across the entire app.

**Subcommand flags** are parsed only for a given subcommand.

And we can tell them apart based on where they are on the command line:

```
go run main.go [GLOBAL_FLAGS] subcommand [LOCAL_FLAGS]
```

In the next chapter we'll see how to specify subcommand flags. Here, we
just wanted to point out how the location of the flag determines its
behavior.
