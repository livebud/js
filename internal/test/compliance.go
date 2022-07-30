package test

import (
	"context"
	"testing"

	"github.com/livebud/js"
	"github.com/matryer/is"
)

type test struct {
	Name   string
	Expect string
	Func   func(t testing.TB, vm js.VM)
}

// Compliance tests
var tests = []*test{
	newTest("TestEval.js", "10", "2*5"),
	newTest("TestJSON.js", `{"a":3}`, `JSON.stringify(JSON.parse("{\"a\":3}"))`),
	newTest("TestDate.js", `2022-08-30T16:12:00.000Z`, `(new Date('2022-08-30T16:12:00.000Z')).toISOString()`),
}

// Compliance runs all the tests as subtests
func Compliance(t *testing.T, vm js.VM) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			test.Func(t, vm)
		})
	}
}

func newTest(name, expect, script string) *test {
	return &test{
		Name:   name,
		Expect: expect,
		Func: func(t testing.TB, vm js.VM) {
			t.Helper()
			is := is.New(t)
			ctx := context.Background()
			result, err := vm.Evaluate(ctx, name, script)
			if err != nil {
				is.Equal(err.Error(), expect)
				return
			}
			is.Equal(result, expect)
		},
	}
}
