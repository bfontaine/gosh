package shell

import (
	"github.com/franela/goblin"
	"os"
	"os/user"
	"runtime"
	"strings"
	"testing"
)

func TestFiles(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("expandPath", func() {

		tmp := os.TempDir()

		usr, _ := user.Current()
		username := usr.Username
		home := usr.HomeDir

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

		g.It("Should leave empty paths as is", func() {
			g.Assert(expandPath("")).Equal("")
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

		g.It("Should replace ~ with the home directory", func() {
			g.Assert(expandPath("~")).Equal(home)
			g.Assert(expandPath("~/")).Equal(home)
			g.Assert(expandPath("~/foo")).Equal(home + slash + "foo")
		})

		g.It("Should not replace ~ if it's not at the beginning", func() {
			g.Assert(expandPath("a/~")).Equal(tmp + "a" + slash + "~")
			g.Assert(expandPath("./~")).Equal(tmp + "~")
			g.Assert(expandPath("a~")).Equal(tmp + "a~")
		})

		g.It("Should replace ~user with the user's home directory", func() {
			g.Assert(expandPath("~" + username)).Equal(home)
			g.Assert(expandPath("~" + username + "/")).Equal(home)
			g.Assert(expandPath("~" + username + "/a")).Equal(home + slash + "a")
		})

		g.It("Should not replace ~user with the user doesn't exist", func() {
			g.Assert(expandPath("~idontexist2")).Equal(tmp + "~idontexist2")
			g.Assert(expandPath("~idontexist2/")).Equal(tmp + "~idontexist2")
			g.Assert(expandPath("~idontexist2/a")).Equal(tmp + "~idontexist2/a")
		})

		g.It("Should not replace ~user if not at the beginning", func() {
			g.Assert(expandPath("x~" + username)).Equal(tmp + "x~" + username)
			g.Assert(expandPath("/~" + username)).Equal("/~" + username)
			g.Assert(expandPath("./~" + username)).Equal(tmp + "~" + username)
		})
	})
}
