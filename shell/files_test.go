package shell

import (
	"github.com/franela/goblin"
	"testing"
)

func Test(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("expandPath", func() {
		g.It("Should leave absolute paths as is", func() {
			g.Assert(expandPath("/foo/bar")).Equal("/foo/bar")
		})
	})
}
