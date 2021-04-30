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
	expect(uint32(123)).toEqual(int64(123))
	expect(nil).toEqual(nil)
	expect(1).not().toEqual(2)

	Y := Dummy{}.Y
	expect(Y).toBeNil()

	list := [2]int{2, 3}
	expect(list).toHaveLength(2)

	dummy1 := Dummy{X: "123"}
	expect(dummy1).toHaveProp("X", "123")

	dummy2 := Dummy{X: "123"}
	expect(dummy1).toEqual(dummy2)

	dummy3 := Dummy{X: "hello"}
	expect(dummy1).not().toEqual(dummy3)
}
