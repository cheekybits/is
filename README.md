is
==

A mini testing helper for Go.

  * Simple interface (`is.OK` and `is.Equal`)
  * Plugs into existing Go toolchain (uses `testing.T`)
  * Obvious for newcomers and noobs

### OK

Make sure an object is not `nil`:

```
func TestSomething(t *testing.T) {
  is := is.New(t)
  obj := SomeFunc()
  is.OK(obj)
}
```

Make sure no errors occurred:

```
func TestSomething(t *testing.T) {
  is := is.New(t)
  obj, err := SomeFunc()
  is.OK(err)
}
```

Make sure a `bool` is not `false`:

```
func TestSomething(t *testing.T) {
  is := is.New(t)
  b := SomeFunc()
  is.OK(b)
}
```

Check many things in one go:

```
func TestSomething(t *testing.T) {
  is := is.New(t)
  b := SomeBool()
  obj, err := SomeFunc()
  is.OK(b, obj, err) // just list things
}
```

### Equal

Are two values equal?

```
func TestSomething(t *testing.T) {
  is := is.New(t)
  obj := SomeFunc()
  is.Equal(obj, "Expected value")
}
```