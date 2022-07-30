package goja_test

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/livebud/js/goja"
	"github.com/livebud/js/internal/test"

	"github.com/matryer/is"
)

func TestEvaluateTwice(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	vm := goja.New(test.Console(os.Stdout, os.Stderr))
	value, err := vm.Evaluate(ctx, "math.js", `const multiply = (a, b) => a * b`)
	is.NoErr(err)
	is.Equal("undefined", value)
	value, err = vm.Evaluate(ctx, "run.js", "multiply(3, 2)")
	is.NoErr(err)
	is.Equal("6", value)
}

func TestConsoleLog(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	stdout := new(bytes.Buffer)
	vm := goja.New(test.Console(stdout, os.Stderr))
	value, err := vm.Evaluate(ctx, "console.js", `console.log("a", 3, { hi: "world" })`)
	is.NoErr(err)
	is.Equal("undefined", value)
	is.Equal("a 3 [object Object]\n", stdout.String())
}

func TestConsoleError(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	stderr := new(bytes.Buffer)
	vm := goja.New(test.Console(os.Stdout, stderr))
	value, err := vm.Evaluate(ctx, "console.js", `console.error("a", 3, { hi: "world" })`)
	is.NoErr(err)
	is.Equal("undefined", value)
	is.Equal("a 3 [object Object]\n", stderr.String())
}

func TestCompliance(t *testing.T) {
	test.Compliance(t, goja.New(test.Console(os.Stdout, os.Stderr)))
}
