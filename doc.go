// Package is is a mini testing helper.
//
//     func TestSomething(t *testing.T) {
//       is := is.New(t)
//       obj, err := MethodBeingTested()
//       is.OK(obj)
//       is.OK(err)
//       is.Equal(obj, "Hello world")
//     }
//
// OK
//
// The OK method asserts that the specified object is OK, which means
// different things for different types:
//
//     bool  - OK means not false
//     int   - OK means not zero
//     error - OK means nil
//     string - OK means not ""
//     everything else - OK means not nil
//
// Equal
//
// The Equal method asserts that two objects are effectively equal.
package is
