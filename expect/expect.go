package expect

import (
	"reflect"
	"testing"
)

type Assertions interface {
	ToEqual(actual interface{}) 
	Not() Assertions
	Expect(expected interface{}) Assertions
	ToHaveLength(length int)
	ToHaveProp(prop string, value interface{})
}

type Expecter struct {
	Inverted bool
	ExpectedValue interface{}
	T *testing.T
}

// func equal(v1, v2 interface{}) bool {
// 	return fmt.Sprintf("%v", v1) == fmt.Sprintf("%v", v2)
// }
// func notEqual(v1, v2 interface{}) bool {
// 	return fmt.Sprintf("%v", v1) != fmt.Sprintf("%v", v2)
// }
func equal(expected, actual interface{}) bool {
	if(expected == nil || actual == nil){
		return expected == actual
	}

	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		return false
	}
	expectedValue := reflect.ValueOf(expected)
	if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
		// Attempt comparison after type conversion
		return reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
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

func (e *Expecter) ToEqual(actual interface{}){
	if notEqual(e.ExpectedValue, actual) && !e.Inverted {
		e.T.Errorf("Expected %v, received %v", e.ExpectedValue, actual)
	} else if equal(e.ExpectedValue, actual) && e.Inverted {
		e.T.Errorf("Expected not equal, but both values were: %v", actual)
	}
}

func (e *Expecter) ToHaveLength(actualLength int){
	ok, expectedLength := getLen(e.ExpectedValue)

	if !ok {
		e.T.Errorf("Type %T does not support length", e.ExpectedValue)
		return
	}

	if expectedLength != actualLength {
		e.T.Errorf("Expected length %v, received %v", expectedLength, actualLength )
	}
}

func (e *Expecter) ToHaveProp(prop string, actualValue interface{}) {
	ok, expectedValue := getField(e.ExpectedValue, prop)

	if !ok {
		e.T.Errorf("Type %T does not have prop %v", e.ExpectedValue, prop)
		return
	}

	if expectedValue != actualValue {
		e.T.Errorf("Expected value %v, received %v", expectedValue, actualValue )
	}
}

func (e *Expecter) Not() Assertions {
	e.Inverted = true

	return e
}

func (e *Expecter) Expect(expected interface{}) Assertions {
	e.ExpectedValue = expected
	e.Inverted = false

	return e 
}

func WithExpect(fn func(e *Expecter)) func (t *testing.T) {
	return func (t *testing.T){
		fn(&Expecter{T: t})
	}
}



