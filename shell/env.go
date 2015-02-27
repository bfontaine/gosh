package shell

import (
	"io/ioutil"
	"os"
	"strings"
)

var MaxComplete = 10

func isExecutable(f os.FileInfo) bool {
	return !f.IsDir() && f.Mode()&0111 != 0
}

func completeCommand(cmdPrefix string) (cmp []string, err error) {
	// a large portion of this code is copied from os/exec/lp_unix.go

	complete := 0

	if cmdPrefix == "" {
		return
	}

	if strings.Contains(cmdPrefix, "/") {
		// we don't support this right now
		return
	}

	pathenv := os.Getenv("PATH")
	if pathenv == "" {
		return
	}

	for _, dir := range strings.Split(pathenv, ":") {
		if dir == "" {
			dir = "."
		}

		files, ok := ioutil.ReadDir(dir)

		if ok != nil {
			err = ok
			return
		}

		for _, f := range files {
			name := f.Name()

			if strings.HasPrefix(name, cmdPrefix) && isExecutable(f) {
				cmp = append(cmp, name)
				if complete++; complete >= MaxComplete {
					return
				}
			}
		}

	}

	return
}
