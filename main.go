package main

import (
	"github.com/Masterminds/cookoo"
	"github.com/Masterminds/cookoo/cli"
)

func main() {
	cli.New(cookoo.Cookoo()).Run("help")
}
