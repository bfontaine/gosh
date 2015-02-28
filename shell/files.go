package shell

import (
	"os"
	"path/filepath"
)

func expandPath(path string) string {
	// replace env variables
	expandedPath := os.ExpandEnv(path)

	// get an absolute path, ignoring errors
	if absPath, err := filepath.Abs(expandedPath); err == nil {
		expandedPath = absPath
	}

	// cleanup
	expandedPath = filepath.Clean(expandedPath)

	return expandedPath
}
