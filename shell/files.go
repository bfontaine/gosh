package shell

import (
	"os"
)

func expandPath(path string) string {
	return os.ExpandEnv(path)
}
