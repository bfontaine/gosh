package shell

import (
	"github.com/franela/goblin"
	o "github.com/onsi/gomega"
	"os"
	"testing"
	"time"
)

type FakeFileInfo struct {
	FakeName  string
	FakeMode  os.FileMode
	FakeIsDir bool
}

func (f FakeFileInfo) Name() string       { return f.FakeName }
func (f FakeFileInfo) Size() int64        { return 0 }
func (f FakeFileInfo) Mode() os.FileMode  { return f.FakeMode }
func (f FakeFileInfo) ModTime() time.Time { return time.Time{} }
func (f FakeFileInfo) IsDir() bool        { return f.FakeIsDir }
func (f FakeFileInfo) Sys() interface{}   { return nil }

func TestEnv(t *testing.T) {
	g := goblin.Goblin(t)

	o.RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("IsExecutable", func() {
		g.It("Should return false for non-executable files", func() {
			fi := FakeFileInfo{FakeName: "foo", FakeMode: 0644}
			o.Expect(isExecutable(fi)).To(o.BeFalse())
		})
		g.It("Should return false for directories", func() {
			fi := FakeFileInfo{FakeName: "foo", FakeMode: 0644, FakeIsDir: true}
			o.Expect(isExecutable(fi)).To(o.BeFalse())
		})
		g.It("Should return false for cd-able directories", func() {
			fi := FakeFileInfo{FakeName: "foo", FakeMode: 0777, FakeIsDir: true}
			o.Expect(isExecutable(fi)).To(o.BeFalse())
		})
		g.It("Should return true for files executable by their owner", func() {
			fi := FakeFileInfo{FakeName: "foo", FakeMode: 0100}
			o.Expect(isExecutable(fi)).To(o.BeTrue())
		})
		g.It("Should return true for files executable by the owner's group", func() {
			fi := FakeFileInfo{FakeName: "foo", FakeMode: 0010}
			o.Expect(isExecutable(fi)).To(o.BeTrue())
		})
		g.It("Should return true for files executable by others", func() {
			fi := FakeFileInfo{FakeName: "foo", FakeMode: 0001}
			o.Expect(isExecutable(fi)).To(o.BeTrue())
		})
	})

	g.Describe("completeCommand", func() {
		g.It("Should not complete empty commands", func() {
			comp, err := completeCommand("")

			o.Expect(err).To(o.BeNil())
			o.Expect(comp).To(o.Equal([]string{}))
		})
	})
}
