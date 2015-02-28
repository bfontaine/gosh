package main

import (
	"flag"
	"fmt"
	"github.com/bfontaine/gosh/shell"
)

func main() {
	debug := flag.Bool("debug", false, "debug mode")
	trace := flag.Bool("trace", false, "trace mode")

	flag.Parse()

	repl := shell.NewRepl("$ ")
	repl.Debug = *debug
	repl.Trace = *trace

	if err := repl.ExecuteRC(); err != nil && repl.Debug {
		fmt.Printf("%v\n", err)
	}

	if err := repl.Loop(); err != nil && repl.Debug {
		fmt.Printf("%v\n", err)
	}
}
