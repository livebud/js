package v8_test

import (
	"context"
	"os"
	"testing"

	"github.com/livebud/js/internal/test"
	v8 "github.com/livebud/js/v8"
	"github.com/matryer/is"
)

func TestEvaluateTwice(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	vm, err := v8.Load(test.Console(os.Stdout, os.Stderr))
	is.NoErr(err)
	defer vm.Close()
	value, err := vm.Evaluate(ctx, "math.js", `const multiply = (a, b) => a * b`)
	is.NoErr(err)
	is.Equal("undefined", value)
	value, err = vm.Evaluate(ctx, "run.js", "multiply(3, 2)")
	is.NoErr(err)
	is.Equal(value, "6")
}

// TODO: move to compliance once goja supports this
func TestFetch(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	vm, err := v8.Load(test.Console(os.Stdout, os.Stderr))
	is.NoErr(err)
	defer vm.Close()
	value, err := vm.Evaluate(ctx, "fetch.js", `fetch("http://google.com").then(res => res.status)`)
	is.NoErr(err)
	is.Equal("200", value)
}

// TODO: move to compliance once goja supports this
func TestURL(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	vm, err := v8.Load(test.Console(os.Stdout, os.Stderr))
	is.NoErr(err)
	defer vm.Close()
	value, err := vm.Evaluate(ctx, "url.js", `(new URL("http://google.com/hi")).host`)
	is.NoErr(err)
	is.Equal("google.com", value)
}

func TestCompliance(t *testing.T) {
	is := is.New(t)
	vm, err := v8.Load(test.Console(os.Stdout, os.Stderr))
	is.NoErr(err)
	test.Compliance(t, vm)
}
