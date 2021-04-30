## go-expect

A super simple assertion library that is heavily inspired by [jest](https://github.com/facebook/jest).

NOTE: This library is incomplete and should not be used in production

## Motivation

I use it for my hobby projects that do not have such high demands on being able to test everything.

Writing an assertion library is a great way to start learning a language, which I'm currently doing with go. I really enjoy writing tests in [jest](https://github.com/facebook/jest) and thought it would be nice to have something similar for my go projects.

## Installation

`go get github.com/hampuslavin/go-expect`

## Example

```go
package your_package

import (
	"testing"
    "github.com/hampuslavin/go-expect/expect"
)

func TestSuite(t *testing.T){
	t.Run("examples", expect.WithExpect(exampleCase))
}

type Dummy struct {
	X string
	Y *struct{ Z string }
}

func examples(expect Expect) {
	expect(uint32(123)).ToEqual(int64(123))
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

```
