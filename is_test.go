package is_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/metabition/is"
)

type mockT struct {
	fail string
}

func (m *mockT) Fatal(a ...interface{}) {
	m.fail = fmt.Sprint(a...)
}
func (m *mockT) Fatalf(f string, a ...interface{}) {
	m.fail = fmt.Sprintf(f, a...)
}
func (m *mockT) Failed() bool {
	return len(m.fail) > 0
}

func TestIs(t *testing.T) {

	for _, test := range []struct {
		N    string
		F    func(is is.I)
		Fail string
	}{
		// OK
		{
			N: "OK(false)",
			F: func(is is.I) {
				is.OK(false)
			},
			Fail: "unexpected false",
		}, {
			N: "OK(true)",
			F: func(is is.I) {
				is.OK(true)
			},
			Fail: "",
		}, {
			N: "OK(nil)",
			F: func(is is.I) {
				is.OK(nil)
			},
			Fail: "unexpected nil",
		}, {
			N: "OK(1,2,3)",
			F: func(is is.I) {
				is.OK(1, 2, 3)
			},
		}, {
			N: "OK(0)",
			F: func(is is.I) {
				is.OK(0)
			},
			Fail: "unexpected zero",
		}, {
			N: "OK(1)",
			F: func(is is.I) {
				is.OK(1)
			},
		}, {
			N: "OK(\"\")",
			F: func(is is.I) {
				is.OK("")
			},
			Fail: "unexpected \"\"",
		}, {
			N: "OK(errors.New(\"an error\"))",
			F: func(is is.I) {
				is.OK(errors.New("an error"))
			},
			Fail: "unexpected error: an error",
		},
		// Equal
		{
			N: "Equal(1,1)",
			F: func(is is.I) {
				is.Equal(1, 1)
			},
		}, {
			N: "Equal(1,2)",
			F: func(is is.I) {
				is.Equal(1, 2)
			},
			Fail: "1 != 2",
		}, {
			N: "Equal(1,nil)",
			F: func(is is.I) {
				is.Equal(1, nil)
			},
			Fail: "1 != <nil>",
		}, {
			N: "Equal(nil,1)",
			F: func(is is.I) {
				is.Equal(nil, 1)
			},
			Fail: "<nil> != 1",
		}, {
			N: "Equal(false,false)",
			F: func(is is.I) {
				is.Equal(false, false)
			},
		}, {
			N: "Equal(map1,map2)",
			F: func(is is.I) {
				is.Equal(
					map[string]interface{}{"package": "is"},
					map[string]interface{}{"package": "is"},
				)
			},
		}} {

		tt := new(mockT)
		is := is.New(tt)
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
			if test.Fail != tt.fail {
				t.Errorf("expected fail \"%s\" but was \"%s\".", test.Fail, tt.fail)
			}
		}

	}

}
