# 2. Hello World

In this chapter, we're going to turn `main.go` into a simple `Hello
World` CLI.

In the last chapter, we ended with a program that looked like this:

```go
package main

import (
	"github.com/Masterminds/cookoo"
	"github.com/Masterminds/cookoo/cli"
)

func main() {
	cli.New(cookoo.Cookoo()).Run("help")
}
```

In this chapter, we're going to expand on the body of the `main`
function. At any time, you can reference the code in this repository, so
we're not going to keep repeating the full source. We're just going to
focus on the changed parts.

We are going to need one more package from Cookoo, though.
`github.com/Masterminds/cookoo/fmt` contains some string formatting
commands that we will be using.

From there, we're mainly going to split the one line of code into a
couple lines, and then add a `hello world` route.

Here's the body of `main`:

```go
func main() {
	reg, router, cxt := cookoo.Cookoo()

	reg.Route("hello", "A Hello World route").
		Does(fmt.Printf, "_").
		Using("format").WithDefault("Hello World!\n")

	cli.New(reg, router, cxt).Run("hello")
}
```

## Cookoo! Cookoo!

First, let's understand what `cookoo.Cookoo()` does. (The duplication of
Cookoo was intentional, and is a reference to a cookoo clock.)

```go
reg, router, cxt := cookoo.Cookoo()
```

The registry, `reg` is where we declare routes. In Cookoo, a route has a
name and description, and then a *chain of commands*. You can think of
this as a task (`hello`) and then a series of steps required to complete
that task.

The router is what runs tasks. In Cookoo CLI apps, we usually don't need
to directly interact with the router. The Cooko CLI runner does that for
us.

The context is where Cookoo stores data during the running of a route.
Later in this tutorial we will interact directly with it, but for now we
just need to basically grok what it's there for.

Architecturally speaking, the context is the base of Cookoo's dependency
injection system. Database connections, caches, and all kinds of
services can get stored in the context, as can ephemeral data.

So when we call `cookoo.Cookoo()`, we get all the pieces we need to
start a Cookoo application.

Next in this code, we build a Cookoo route.

## Hello, the Route!

At the center of all Cookoo apps is the concept of the route. Again, a
route is a *task* with a number of *steps*. We call this patter the
*chain of command (CoCo)* pattern.

In our simple app, we have a single route, and it does only one thing:
It prints "Hello World!" to the standard output.

```go
	reg.Route("hello", "A Hello World route").
		Does(fmt.Printf, "_").
		Using("format").WithDefault("Hello World!\n")
```

The `Route()` function takes two arguments. The first is the *route name*.
This plays a strong functional role. We use the route name to refer to
this route. (Peek ahead to the `Run()` function. See?!)

A route also has a description. The description is targeted to
*developers*, and should not generally be displayed to end users. Cookoo
can use these descriptions to automatically document a program.

A route *does* stuff. So each `Route()` is followed by any number of
`Does()` calls. In our app, the *hello* route `Does(fmt.Printf, "_")`.

A `Does()` call takes two arguments: the Cookoo command to run and a name
for the output of this command. "_" conventionally means "ignore the output
of this command. It's not important.")

The `fmt.Printf` command is a Cookoo built-in function that wraps the
`fmt.Printf` function in Go's core.

Commands take arguments, and we pass the arguments into a command using
`Using()`:

```go
		Does(fmt.Printf, "_").
		Using("format").WithDefault("Hello World!\n")
```

The `Using` function is designed to pass name/value pairs into the
function in the `Does` command. So we can read the above as...

*Does fmt.Printf using the param "format" with the default value "Hello
World". Essentially, this describes `fmt.Printf("Hello World!"\n")`
(though with lots of hidden extras). In the next chapter we'll see how
to pass some non-default info into a command.

## Running Our Route

Now we're ready to turn our Cookoo app into a command line tool:

```go
	cli.New(reg, router, cxt).Run("hello")
```

As we saw in the last chapter, `cli.New()` creates a new CLI Runner. A
runner requires all three parts of a Cookoo app -- the registry, the
router, and the context.

Finally, whenever `main()` is executed, we want our app to run (`Run()`)
our new "hello" route.

If we execute the above, this is what we should get:

```
$ go run main.go
Hello World!
```

(And we still get really lame help with `go run main.go -h`)

That's it! We've got our second CLI. In the next chapter, we'll expand
this into the "Hello You!" app.
