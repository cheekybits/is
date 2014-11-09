// Package is is a mini testing helper.
//
//     func TestSomething(t *testing.T) {
//       is := is.New(t)
//       obj, err := MethodBeingTested()
//       is.OK(obj, err) // list of objects, all must be OK
//       is.Equal(obj, "Hello world")
//     }
//
// is.OK
//
// is.OK asserts that the specified object is OK, which means
// different things for different types:
//
//     bool  - OK means not false
//     int   - OK means not zero
//     error - OK means nil
//     string - OK means not ""
//     everything else - OK means not nil
//
// is.Equal
//
// is.Equal asserts that two objects are effectively equal.
//
// is.Panic and is.PanicWith
//
// is.Panic and is.PanicWith asserts that the func() will panic.
// PanicWith specifies the panic text that is expected:
//
//     func TestInvalidArgs(t *testing.T) {
//       is := is.New(t)
//       is.Panic(func(){
//         SomeMethod(1)
//       })
//       is.PanicWith("invalid args, both cannot be nil", func(){
//         OtherMethod(nil, nil)
//       })
//     }
package is
