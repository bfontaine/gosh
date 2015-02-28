package shell

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var slash = string(os.PathSeparator)

func expandPath(path string) string {

	length := len(path)

	if length == 0 {
		return path
	}

	// replace env variables
	expandedPath := os.ExpandEnv(path)

	// replace ~ with $HOME
	if (length == 1 && path[0] == '~') || (length > 1 && path[:2] == "~/") {
		usr, _ := user.Current()

		expandedPath = strings.Replace(expandedPath, "~", usr.HomeDir, 1)
	} else if path[:1] == "~" {
		// replace ~user with their $HOME

		firstSlash := strings.Index(path, slash)
		if firstSlash < 0 {
			firstSlash = len(path)
		}

		if firstSlash > 1 {
			usr, err := user.Lookup(path[1:firstSlash])
			if err == nil {
				expandedPath = usr.HomeDir + path[firstSlash:]
			}
		}
	}

	// get an absolute path, ignoring errors
	if absPath, err := filepath.Abs(expandedPath); err == nil {
		expandedPath = absPath
	}

	// cleanup
	expandedPath = filepath.Clean(expandedPath)

	return expandedPath
}
