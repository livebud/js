package v8_test

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/livebud/js/internal/test"
	v8 "github.com/livebud/js/v8"
	"github.com/matryer/is"
	"golang.org/x/sync/errgroup"
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

func TestConsoleLog(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	stdout := new(bytes.Buffer)
	vm, err := v8.Load(test.Console(stdout, os.Stderr))
	is.NoErr(err)
	defer vm.Close()
	value, err := vm.Evaluate(ctx, "console.js", `console.log("a", 3, { hi: "world" })`)
	is.NoErr(err)
	is.Equal("undefined", value)
	is.Equal("a 3 [object Object]\n", stdout.String())
}

func TestConsoleError(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	stderr := new(bytes.Buffer)
	vm, err := v8.Load(test.Console(os.Stdout, stderr))
	is.NoErr(err)
	defer vm.Close()
	value, err := vm.Evaluate(ctx, "console.js", `console.error("a", 3, { hi: "world" })`)
	is.NoErr(err)
	is.Equal("undefined", value)
	is.Equal("a 3 [object Object]\n", stderr.String())
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

func TestPromiseWithSetTimeout(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	vm, err := v8.Load(test.Console(os.Stdout, os.Stderr))
	is.NoErr(err)
	defer vm.Close()
	result, err := vm.Evaluate(ctx, "promise.js", `
		new Promise(function (resolve) {
			setTimeout(function () {
				resolve("hello")
			}, 1000)
		})
	`)
	is.NoErr(err)
	is.Equal(result, "hello")
}

func TestConcurrency(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()
	vm, err := v8.Load(test.Console(os.Stdout, os.Stderr))
	is.NoErr(err)
	defer vm.Close()
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
