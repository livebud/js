package test

import (
	"context"
	"testing"

	"github.com/livebud/js"
	"github.com/matryer/is"
	"golang.org/x/sync/errgroup"
)

type test struct {
	Name string
	Func func(t testing.TB, vm js.VM)
}

// Compliance tests
var tests = []*test{
	newTest("TestEval.js", "10", "2*5"),
	newTest("TestJSON.js", `{"a":3}`, `JSON.stringify(JSON.parse("{\"a\":3}"))`),
	newTest("TestDate.js", `2022-08-30T16:12:00.000Z`, `(new Date('2022-08-30T16:12:00.000Z')).toISOString()`),
	&test{Name: "TestConcurrency", Func: testConcurrency},
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
		Name: name,
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

func testConcurrency(t testing.TB, vm js.VM) {
	is := is.New(t)
	ctx := context.Background()
	value, err := vm.Evaluate(ctx, "math.js", `const multiply = (a, b) => a * b`)
	is.NoErr(err)
	is.Equal("undefined", value)
	eg := new(errgroup.Group)
	for i := 0; i < 100; i++ {
		eg.Go(func() error {
			value, err := vm.Evaluate(ctx, "run.js", "multiply(3, 2)")
			if err != nil {
				return err
			}
			if value != "6" {
				return err
			}
			return nil
		})
	}
	is.NoErr(eg.Wait())
}
