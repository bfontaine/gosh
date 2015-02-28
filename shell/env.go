package shell

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// some parts of this code are copied from os/exec/lp_unix.go

var MaxComplete = 10

func isExecutable(f os.FileInfo) bool {
	return !f.IsDir() && f.Mode()&0111 != 0
}

func completeCommand(cmdPrefix string) (cmp []string, err error) {

	cmp = make([]string, 0, MaxComplete)

	complete := 0

	if cmdPrefix == "" || complete >= MaxComplete {
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

	for _, dir := range filepath.SplitList(pathenv) {
		if dir == "" {
			dir = "."
		}

		files, failure := ioutil.ReadDir(dir)

		if failure != nil {
			err = failure
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
