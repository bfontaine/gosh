package shell

import (
	"fmt"
	"github.com/bfontaine/go-readline"
	"os"
	"os/exec"
	"strings"
)

type Repl struct {
	Prompt string

	lastError error
}

var builtins = []string{"cd", "exit", "quit"}

func (r *Repl) complete(input, line string, start, end int) (cmp []string) {
	for _, builtin := range builtins {
		if strings.HasPrefix(builtin, input) {
			cmp = append(cmp, builtin)
		}
	}
	return
}

func (r *Repl) fail(err error) {
	r.lastError = err
	fmt.Printf("%v\n", err)
}

func (r *Repl) execute(line string) (exit, history bool) {
	words := strings.Fields(strings.TrimSpace(line))
	wordsCount := len(words)

	if wordsCount == 0 {
		return
	}

	history = true

	// try builtin commands
	switch words[0] {
	case "cd":
		directory := strings.Join(words[1:], " ")
		if err := os.Chdir(directory); err != nil {
			r.fail(err)
			history = false
			return
		}
	case "quit":
		fallthrough
	case "exit":
		exit = true
		return
	}

	cmd := exec.Command(words[0], words[1:]...)

	out, err := cmd.CombinedOutput()

	if err != nil {
		r.fail(err)
		history = false
		return
	}

	fmt.Print(string(out))

	return
}

func (r *Repl) Loop() (err error) {

	// tab-completion
	readline.SetCompletionFunction(r.complete)
	readline.ParseAndBind("TAB: menu-complete")

	for {
		line := readline.Readline(&r.Prompt)
		// ^D
		if line == nil {
			break
		}

		if exit, history := r.execute(*line); exit {
			break
		} else if history {
			readline.AddHistory(*line)
		}
	}

	return err
}
