package expect

import (
	"reflect"
	"testing"
)

type Assertions interface {
	ToEqual(actual interface{})
	Not() Assertions
	ToHaveLength(length int)
	ToHaveProp(prop string, value interface{})
	ToBeNil()
	ToBe(actual interface{})
}

type Expecter struct {
	Inverted    bool
	ActualValue interface{}
	T           *testing.T
}

// func equal(v1, v2 interface{}) bool {
// 	return fmt.Sprintf("%v", v1) == fmt.Sprintf("%v", v2)
// }
// func notEqual(v1, v2 interface{}) bool {
// 	return fmt.Sprintf("%v", v1) != fmt.Sprintf("%v", v2)
// }
func equal(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		return false
	}
	ActualValue := reflect.ValueOf(expected)
	if ActualValue.IsValid() && ActualValue.Type().ConvertibleTo(actualType) {
		// Attempt comparison after type conversion
		return reflect.DeepEqual(ActualValue.Convert(actualType).Interface(), actual)
	}

	return false
}
func notEqual(v1, v2 interface{}) bool {
	return !equal(v1, v2)
}

func getLen(x interface{}) (ok bool, length int) {
	v := reflect.ValueOf(x)
	defer func() {
		if e := recover(); e != nil {
			ok = false
		}
	}()
	return true, v.Len()
}

func getField(v interface{}, field string) (ok bool, value interface{}) {
	r := reflect.ValueOf(v)
	defer func() {
		if e := recover(); e != nil {
			ok = false
		}
	}()

	f := reflect.Indirect(r).FieldByName(field)
	return true, f.Interface()
}

func (e *Expecter) ToEqual(expected interface{}) {
	if !e.Inverted && notEqual(e.ActualValue, expected) {
		e.T.Errorf("Expected %v, received %v", e.ActualValue, expected)
	} else if e.Inverted && equal(e.ActualValue, expected) {
		e.T.Errorf("Expected not equal, but both values were: %v", expected)
	}
}

func (e *Expecter) ToBe(expected interface{}) {
	if !e.Inverted && e.ActualValue != expected {
		e.T.Errorf("Expected %v, received %v", expected, e.ActualValue)
	} else if e.Inverted && e.ActualValue == expected {
		e.T.Errorf("Expected not equal, but both values were: %v", expected)
	}
}

func (e *Expecter) ToHaveLength(expectedLength int) {
	ok, actualLength := getLen(e.ActualValue)

	if !ok {
		e.T.Errorf("Type %T does not support length", e.ActualValue)
		return
	}

	if actualLength != expectedLength {
		e.T.Errorf("Expected length %v, received %v", expectedLength, actualLength)
	}
}

func (e *Expecter) ToHaveProp(prop string, expectedValue interface{}) {
	ok, actualValue := getField(e.ActualValue, prop)

	if !ok {
		e.T.Errorf("Type %T does not have prop %v", e.ActualValue, prop)
		return
	}

	if actualValue != expectedValue {
		e.T.Errorf("Expected value %v, received %v", expectedValue, actualValue)
	}
}

func (e *Expecter) ToBeNil() {
	if !e.Inverted && !reflect.ValueOf(e.ActualValue).IsNil() {
		e.T.Errorf("Expected non-nil value %v to be nil", e.ActualValue)
	}
	// else if e.Inverted && reflect.ValueOf(e.ActualValue).IsNil() {
	// 	e.T.Errorf("Expected nil value to not be nil")
	// }
}

func (e *Expecter) Not() Assertions {
	e.Inverted = true

	return e
}

type Expect func(expected interface{}) Assertions

func getExpect(t *testing.T) Expect {

	return func(expected interface{}) Assertions {
		return &Expecter{ActualValue: expected, Inverted: false, T: t}
	}
}

func WithExpect(fn func(e Expect)) func(t *testing.T) {
	return func(t *testing.T) {
		fn(getExpect(t))
	}
}
