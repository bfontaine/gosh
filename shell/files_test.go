package shell

import (
	"github.com/franela/goblin"
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestFiles(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("expandPath", func() {

		tmp := os.TempDir()
		slash := string(os.PathSeparator)

		// On OS X, /var is a symlink to /private/var and expandPath doesn't
		// follow symlinks.
		if runtime.GOOS == "darwin" {
			tmp = "/private" + tmp
		}

		if !strings.HasSuffix(tmp, slash) {
			tmp += slash
		}

		g.BeforeEach(func() {
			os.Setenv("foo", "bar")
			os.Setenv("parent", "..")
			os.Chdir(tmp)
		})

		g.AfterEach(func() {
			// os.Unsetenv is not defined in Go <1.4
			os.Setenv("foo", "")
			os.Setenv("parent", "")
		})

		g.It("Should leave absolute paths as is", func() {
			g.Assert(expandPath("/foo/bar")).Equal("/foo/bar")
		})

		g.It("Should expand $env variables", func() {
			g.Assert(expandPath("/foo/$foo")).Equal("/foo/bar")
		})

		g.It("Should expand ${env} variables", func() {
			g.Assert(expandPath("/foo/x${foo}x")).Equal("/foo/xbarx")
		})

		g.It("Should return absolute paths", func() {
			g.Assert(expandPath("./bar")).Equal(tmp + "bar")
		})

		g.It("Should interpret $env vars before absolut-ing paths", func() {
			g.Assert(expandPath("./foo/$parent/bar")).Equal(tmp + "bar")
		})

		g.It("Should cleanup paths", func() {
			g.Assert(expandPath("/./././foo/.././qux")).Equal("/qux")
		})
	})
}
