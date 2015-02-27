package main

import (
	"flag"
	"github.com/bfontaine/gosh/shell"
)

func main() {
	flag.Parse()

	repl := shell.Repl{Prompt: "$ "}

	repl.Loop()
}
