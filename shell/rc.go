package shell

import (
	"bufio"
	"os"
)

var rc = expandPath("$HOME/.goshrc")

// Read ~/.goshrc and execute each line as if it were given to the REPL
func (r *Repl) ExecuteRC() error {
	file, err := os.Open(rc)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if exit, _ := r.execute(scanner.Text()); exit {
			return nil
		}
	}

	return scanner.Err()
}
