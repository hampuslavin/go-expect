package expect

import (
	"testing"
)

func TestSuite(t *testing.T) {
	t.Run("examples", WithExpect(examples))
}

type Dummy struct {
	X string
	Y *struct{ Z string }
}

func examples(e *Expecter) {
	e.Expect(uint32(123)).ToEqual(int64(123))
	e.Expect(nil).ToEqual(nil)
	e.Expect(1).Not().ToEqual(2)

	Y := Dummy{}.Y
	e.Expect(Y).ToEqual(nil)

	list := [2]int{2, 3}
	e.Expect(list).ToHaveLength(2)

	dummy1 := Dummy{X: "123"}
	e.Expect(dummy1).ToHaveProp("X", "123")

	dummy2 := Dummy{X: "123"}
	e.Expect(dummy1).ToEqual(dummy2)

	dummy3 := Dummy{X: "hello"}
	e.Expect(dummy1).Not().ToEqual(dummy3)
}
