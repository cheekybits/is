package is

import (
	"errors"
	"testing"
)

type customErr struct{}

func (_ *customErr) Error() string {
	return "Oops"
}

type mockT struct {
	failed bool
}

func (m *mockT) FailNow() {
	m.failed = true
}
func (m *mockT) Failed() bool {
	return m.failed
}

func TestIs(t *testing.T) {

	for _, test := range []struct {
		N    string
		F    func(is I)
		Fail string
	}{
		// is.OK
		{
			N: "OK(false)",
			F: func(is I) {
				is.OK(false)
			},
			Fail: "unexpected false",
		}, {
			N: "OK(true)",
			F: func(is I) {
				is.OK(true)
			},
		}, {
			N: "OK(nil)",
			F: func(is I) {
				is.OK(nil)
			},
			Fail: "unexpected nil",
		}, {
			N: "OK(1,2,3)",
			F: func(is I) {
				is.OK(1, 2, 3)
			},
		}, {
			N: "OK(0)",
			F: func(is I) {
				is.OK(0)
			},
			Fail: "unexpected zero",
		}, {
			N: "OK(1)",
			F: func(is I) {
				is.OK(1)
			},
		}, {
			N: "OK(\"\")",
			F: func(is I) {
				is.OK("")
			},
			Fail: "unexpected \"\"",
		}, {
			N: "OK(errors.New(\"an error\"))",
			F: func(is I) {
				is.OK(errors.New("an error"))
			},
			Fail: "unexpected error: an error",
		}, {
			N: "OK(&customErr{})",
			F: func(is I) {
				is.OK(&customErr{})
			},
			Fail: "unexpected error: Oops",
		}, {
			N: "OK(error(nil))",
			F: func(is I) {
				var err error
				is.OK("Yep", err)
			},
		}, {
			N: "OK(customErr(nil))",
			F: func(is I) {
				var err *customErr
				is.OK("Yep", err)
			},
		}, {
			N: "OK(func) panic",
			F: func(is I) {
				is.OK(func() {
					panic("panic message")
				})
			},
			Fail: "unexpected panic: panic message",
		}, {
			N: "OK(func) no panic",
			F: func(is I) {
				is.OK(func() {})
			},
		},
		// is.Panic
		{
			N: "PanicWith(\"panic message\", func(){ panic() })",
			F: func(is I) {
				is.PanicWith("panic message", func() {
					panic("panic message")
				})
			},
		},
		{
			N: "PanicWith(\"panic message\", func(){ /* no panic */ })",
			F: func(is I) {
				is.PanicWith("panic message", func() {
				})
			},
			Fail: "expected panic: \"panic message\"",
		},
		{
			N: "Panic(func(){ panic() })",
			F: func(is I) {
				is.Panic(func() {
					panic("panic message")
				})
			},
		},
		{
			N: "Panic(func(){ /* no panic */ })",
			F: func(is I) {
				is.Panic(func() {
				})
			},
			Fail: "expected panic",
		},
		// is.Equal
		{
			N: "Equal(1,1)",
			F: func(is I) {
				is.Equal(1, 1)
			},
		}, {
			N: "Equal(1,2)",
			F: func(is I) {
				is.Equal(1, 2)
			},
			Fail: "1 != 2",
		}, {
			N: "Equal(1,nil)",
			F: func(is I) {
				is.Equal(1, nil)
			},
			Fail: "1 != <nil>",
		}, {
			N: "Equal(nil,1)",
			F: func(is I) {
				is.Equal(nil, 1)
			},
			Fail: "<nil> != 1",
		}, {
			N: "Equal(false,false)",
			F: func(is I) {
				is.Equal(false, false)
			},
		}, {
			N: "Equal(map1,map2)",
			F: func(is I) {
				is.Equal(
					map[string]interface{}{"package": "is"},
					map[string]interface{}{"package": "is"},
				)
			},
		}} {

		tt := new(mockT)
		is := New(tt)
		var rec interface{}

		func() {
			defer func() {
				rec = recover()
			}()
			test.F(is)
		}()

		if len(test.Fail) > 0 {
			if !tt.Failed() {
				t.Errorf("%s should fail", test.N)
			}
			if test.Fail != is.(*i).last {
				t.Errorf("expected fail \"%s\" but was \"%s\".", test.Fail, is.(*i).last)
			}
		}

	}

}
