package is

import (
	"fmt"
	"reflect"
	"sync"
)

// I represents the is interface.
type I interface {
	// OK asserts that the specified objects are all OK.
	OK(o ...interface{})
	// Equal asserts that the two values are
	// considered equal. Non strict.
	Equal(a, b interface{})
	// Panic asserts that the specified function
	// panics.
	Panic(fn func())
	// PanicWith asserts that the specified function
	// panics with the specific message.
	PanicWith(m string, fn func())
}

// T represents the an interface for reporting
// failures.
// testing.T satisfied this interface.
type T interface {
	FailNow()
}

// i represents an implementation of interface I.
type i struct {
	t    T
	last string
	l    sync.Mutex
}

func (i *i) Log(args ...interface{}) {
	i.l.Lock()
	i.last = fmt.Sprint(args...)
	fmt.Print(decorate(i.last))
	i.l.Unlock()
}
func (i *i) Logf(format string, args ...interface{}) {
	i.l.Lock()
	i.last = fmt.Sprint(fmt.Sprintf(format, args...))
	fmt.Print(decorate(i.last))
	i.l.Unlock()
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
		i.Logf("%v != %v", a, b)
		i.t.FailNow()
	}
}

// Panic asserts that the specified function
// panics.
func (i *i) Panic(fn func()) {
	var r interface{}
	func() {
		defer func() {
			r = recover()
		}()
		fn()
	}()
	if r == nil {
		i.Log("expected panic")
		i.t.FailNow()
	}
}

// PanicWith asserts that the specified function
// panics with the specific message.
func (i *i) PanicWith(m string, fn func()) {
	var r interface{}
	func() {
		defer func() {
			r = recover()
		}()
		fn()
	}()
	if r != m {
		i.Logf("expected panic: \"%s\"", m)
		i.t.FailNow()
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
		// shouldn't panic
		var r interface{}
		func() {
			defer func() {
				r = recover()
			}()
			co()
		}()
		if r != nil {
			i.Logf("unexpected panic: %v", r)
			i.t.FailNow()
		}
	case error:
		if co != nil {
			i.Log("unexpected error: " + co.Error())
			i.t.FailNow()
		}
	case string:
		if len(co) == 0 {
			i.Log("unexpected \"\"")
			i.t.FailNow()
		}
	case bool:
		// false
		if co == false {
			i.Log("unexpected false")
			i.t.FailNow()
			return
		}
	}
	if isNil(o) {
		i.Log("unexpected nil")
		i.t.FailNow()
		return
	}
	if o == 0 {
		i.Log("unexpected zero")
		i.t.FailNow()
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
