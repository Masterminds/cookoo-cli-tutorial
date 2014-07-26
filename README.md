# 1. The Simplest CLI

Let's dive right in! In this chapter, let's create the simplest Cookoo CLI we
can.

## What Is Cookoo?

Cookoo is a development framework. Developers tend to refer to it as a
*middleware platform* because it still leaves you (the developer) with
the flexibility to structure applications the way you want, but it
handles the top-level structure for you.

Cookoo is not just designed to be an easy starting point, but a
structure that encourages developers to structure applications for
future development. In other words, it's a forward-thinking framework.

Cookoo values:

- Ease of use
- Separation of concerns
- Dependency injection
- Component re-use
- Extensibility and flexibility at the topmost levels

When it comes to Cookoo CLIs, these values manifiest in the way that
Cookoo separates a CLI into obvious chunks. We'll get into that in the
coming chapters, but let's start with the simplest program we can make.

## Setup

If you are a [Glide](https://github.com/Masterminds/glide) user, you can
just run `glide install` in this directory. Then `glide in` to set your
`$GOPATH`.

If you're not, you may want to `go get github.com/Masterminds/cookoo`.
Make sure that your `$GOPATH` is correctly configured.

## 10 Lines to Happiness!

Without further ado, here's our first app. You can view it in all its
glory (???) in `main.go`.

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

That's it! The above is now a stimulatingly simple and functional
commandline app.

What does it do? I'm glad you asked.

It does.... nothing.

Well, that's not completely true. You can run it like this:

```
$ go run main.go -h
```

And it will dutifully print some very boilerplate help text.

You can also run it like this:

```
$ go run main.go
```

And it will dutifully print some very boilerplate help text.

Let's take a quick look at why.

## How This Actually Works

Really, there's only one important line in the program above:

```go
cli.New(cookoo.Cookoo()).Run("help")
```

Unintuitively, let's start with `cookoo.Cookoo()`. That function is the
high progenitor of all Cookoo apps. Or, stated less grandiously,
`cookoo.Cookoo()` creates Cookoo apps.

It returns three things:

- A registry (`*cookoo.Registry`) for you to declare routes.
- A router (`*cookoo.Router`) for you to execute routes.
- A context (`cookoo.Context`) for the app to pass data and dependencies
  around.

Coincidentally (or, well... not coincidentally at all, actually), a new
Cookoo CLI requires all three of those pieces.

So when we create a new Cookoo CLI with `cli.New()`, we pass it a
registry, a router, and a context.

`cli.New` creates a new CLI application which can then be run whenever
we're ready. That's what the `Run()` function does. It takes one
argument, which is the name of the route to run. We'll look at this in
detail in the next chapter. For now, we can see that it runs `help`,
which we can safely (and correctly) assume generates help text.

That's *almost* all there is to say about this program. But there is one
more thing: When we add `-h` (`go run main.go -h`), it also prints help text.
Pass in any other flag and it will cause an error.  Why? Because The `-h` is a
built-in default for Cookoo CLIs. Later, we'll see how to work with
command line flags.

So there we are. We have the simplest Cookoo app we can build. In the
next chapter, we'll turn this into *ye olde Hello Worlde* app.
