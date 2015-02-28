package shell

import (
	"github.com/franela/goblin"
	o "github.com/onsi/gomega"
	"os"
	"testing"
)

func TestRepl(t *testing.T) {
	g := goblin.Goblin(t)

	o.RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("NewRepl", func() {
		g.It("Should not return nil", func() {
			o.Expect(NewRepl("")).NotTo(o.BeNil())
		})
	})

	g.Describe("Repl", func() {
		origPath := os.Getenv("PATH")
		var r *Repl

		g.BeforeEach(func() {
			os.Setenv("PATH", "")
			r = NewRepl("> ")
		})
		g.AfterEach(func() { os.Setenv("PATH", origPath) })

		g.Describe("#complete", func() {
			g.It("Should not complete empty lines", func() {
				o.Expect(r.complete("", "", 0, 0)).To(o.Equal([]string{}))
			})
			g.It("Should not complete commented lines", func() {
				o.Expect(r.complete("#", "#", 0, 1)).To(o.Equal([]string{}))
			})

			g.It("Should complete builtin commands", func() {
				o.Expect(r.complete("ali", "", 0, 0)).To(o.Equal([]string{
					"alias",
				}))

				o.Expect(r.complete("e", "", 0, 0)).To(o.Equal([]string{
					"echo", "exit",
				}))
			})
		})

		g.Describe("#fail", func() {})
		g.Describe("#trace", func() {})

		g.Describe("#execute", func() {
			g.It("Should not execute empty lines", func() {
				exit, hist := r.execute("")
				o.Expect(exit).To(o.BeFalse())
				o.Expect(hist).To(o.BeFalse())
			})

			g.It("Should not execute commented lines", func() {
				exit, hist := r.execute("#exit")
				o.Expect(exit).To(o.BeFalse())
				o.Expect(hist).To(o.BeFalse())
			})
		})

		g.Describe("#Loop", func() {})
	})
}
