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

func examples(expect Expect) {
	expect(uint32(123)).ToEqual(int64(123))
	expect(123).ToBe(123)
	expect("hello").ToBe("hello")
	expect(nil).ToEqual(nil)
	expect(1).Not().ToEqual(2)

	Y := Dummy{}.Y
	expect(Y).ToBeNil()

	list := [2]int{2, 3}
	expect(list).ToHaveLength(2)

	dummy1 := Dummy{X: "123"}
	expect(dummy1).ToHaveProp("X", "123")

	dummy2 := Dummy{X: "123"}
	expect(dummy1).ToEqual(dummy2)

	dummy3 := Dummy{X: "hello"}
	expect(dummy1).Not().ToEqual(dummy3)
}
