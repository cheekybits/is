is
==

A mini testing helper for Go.

  * Simple interface (`is.OK` and `is.Equal`)
  * Plugs into existing Go toolchain (uses `testing.T`)
  * Obvious for newcomers and noobs

### Usage

  1. Write test functions as usual
  1. Add `is := is.New(t)` at top
  1. Call target code
  1. Make assertions using new `is` object

```
func TestSomething(t *testing.T) {
  is := is.New(t)

  // OK
  // --

  // ensure not nil
  obj := SomeFunc()
  is.OK(obj)

  // ensure no error
  obj, err := SomeFunc()
  is.OK(err)

  // ensure not false
  b := SomeBool()
  is.OK(b)

  // ensure not ""
  s := SomeString()
  is.OK(s)

  // ensure not zero
  is.OK(len(something))

  // ensure many things in one go
  is.OK(b, err, obj, "something")

  // Equal
  // -----

  // make sure two values are equal
  is.Equal(1, 2)
  is.Equal(err, ErrSomething)
  is.Equal(a, b)

}
```