package is

import (
	"fmt"
	"reflect"
)

// I represents the is interface.
type I interface {
	// OK asserts that the specified objects are all OK.
	OK(o ...interface{})
	// Equal asserts that the two values are
	// considered equal. Non strict.
	Equal(a, b interface{})
}

// T represents the an interface for reporting
// failures.
// testing.T satisfied this interface.
type T interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

// i represents an implementation of interface I.
type i struct {
	t T
}

// OK asserts that the specified objects are all OK.
func (i *i) OK(o ...interface{}) {
	for _, obj := range o {
		i.isOK(obj)
	}
}

// Equal asserts that the two values are
// considered equal. Non strict.
func (i *i) Equal(a, b interface{}) {
	if !areEqual(a, b) {
		i.t.Fatalf("%v != %v", a, b)
	}
}

// New creates a new I capable of making
// assertions.
func New(t T) I {
	return &i{t: t}
}

// isNil gets whether the object is nil or not.
func isNil(object interface{}) bool {
	if object == nil {
		return true
	}
	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}
	return false
}

func (i *i) isOK(o interface{}) {
	switch co := o.(type) {
	case func():
		// does it panic or not?
		var r interface{}
		func() {
			defer func() {
				r = recover()
			}()
			co()
		}()
		if r != nil {
			i.t.Fatalf("unexpected panic: %v", r)
		}
	case error:
		if co != nil {
			i.t.Fatal("unexpected error: " + co.Error())
		}
	case string:
		if len(co) == 0 {
			i.t.Fatal("unexpected \"\"")
		}
	case bool:
		// false
		if co == false {
			i.t.Fatal("unexpected false")
			return
		}
	}
	if isNil(o) {
		i.t.Fatal("unexpected nil")
		return
	}
	if o == 0 {
		i.t.Fatal("unexpected zero")
	}
}

// areEqual gets whether a equals b or not.
func areEqual(a, b interface{}) bool {
	if isNil(a) || isNil(b) {
		return a == b
	}
	if reflect.DeepEqual(a, b) {
		return true
	}
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	if aValue == bValue {
		return true
	}
	// Attempt comparison after type conversion
	if bValue.Type().ConvertibleTo(aValue.Type()) && aValue == bValue.Convert(aValue.Type()) {
		return true
	}
	// Last ditch effort
	if fmt.Sprintf("%#v", a) == fmt.Sprintf("%#v", b) {
		return true
	}
	return false
}
