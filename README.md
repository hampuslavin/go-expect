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
	t.Run("examples", WithExpect(exampleCase))
}

type Dummy struct {
	X string
}

func exampleCase(e *Expecter){
	e.Expect(1).ToEqual(1)
	e.Expect(uint32(123)).ToEqual(int64(123))
	e.Expect(nil).ToEqual(nil)
	e.Expect(1).Not().ToEqual(2)

	list := [2]int{2, 3}
	e.Expect(list).ToHaveLength(2)

	dummy1 := Dummy{X: "123"}
	e.Expect(dummy1).ToHaveProp("X", "123")

	dummy2 := Dummy{X: "123"}
	e.Expect(dummy1).ToEqual(dummy2)

	dummy3 := Dummy{X: "hello"}
	e.Expect(dummy1).Not().ToEqual(dummy3)
}
```
