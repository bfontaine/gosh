package shell

import (
	"github.com/franela/goblin"
	o "github.com/onsi/gomega"
	"testing"
)

func TestAliases(t *testing.T) {
	g := goblin.Goblin(t)

	o.RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("NewAliases", func() {
		g.It("Should not return nil", func() {
			o.Expect(NewAliases()).NotTo(o.BeNil())
		})
	})

	g.Describe("Aliases", func() {
		var aliases *Aliases

		g.BeforeEach(func() { aliases = NewAliases() })

		g.Describe("#parse", func() {
			g.It("Should return ErrNeedValue if no value is given", func() {
				o.Expect(aliases.parse("foo")).To(o.Equal(ErrNeedValue))
				o.Expect(aliases.parse(" foo ")).To(o.Equal(ErrNeedValue))
				o.Expect(aliases.parse("foo=")).To(o.Equal(ErrNeedValue))
				o.Expect(aliases.parse("foo= ")).To(o.Equal(ErrNeedValue))
			})

			g.It("Should return ErrNeedLabel if no label is given", func() {
				o.Expect(aliases.parse("=bar")).To(o.Equal(ErrNeedLabel))
				o.Expect(aliases.parse(" =bar")).To(o.Equal(ErrNeedLabel))
			})

			g.It("Should not fail if both the label and the value are non-empty", func() {
				o.Expect(aliases.parse("a=b")).To(o.BeNil())
			})

			g.It("Should ignore spaces arround the '=' symbol", func() {
				o.Expect(aliases.parse("a = b")).To(o.BeNil())
			})

			g.It("Should not fail if the alias already exist", func() {
				o.Expect(aliases.parse("a=b")).To(o.BeNil())
				o.Expect(aliases.parse("a=c")).To(o.BeNil())
			})
		})

		g.Describe("#has", func() {
			g.It("Should return false if the alias doesn't exist", func() {
				o.Expect(aliases.has("foo")).To(o.BeFalse())
			})

			g.It("Should return false if the alias is empty", func() {
				o.Expect(aliases.has("")).To(o.BeFalse())
			})

			g.It("Should return true if the alias exists", func() {
				o.Expect(aliases.parse("a=b")).To(o.BeNil())
				o.Expect(aliases.has("a")).To(o.BeTrue())
			})
		})

		g.Describe("#get", func() {
			g.It("Should return nil if the alias doesn't exist", func() {
				o.Expect(aliases.get("foo")).To(o.BeNil())
			})

			g.It("Should return the alias as a list of words", func() {
				o.Expect(aliases.parse("a=x y z")).To(o.BeNil())
				o.Expect(aliases.get("a")).To(o.Equal([]string{"x", "y", "z"}))
			})
		})
	})

}
