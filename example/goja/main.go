package main

import (
	"context"
	"fmt"
	"os"

	"github.com/livebud/js"
	"github.com/livebud/js/goja"
)

func main() {
	vm := goja.New(&js.Console{
		Log:   os.Stdout,
		Error: os.Stderr,
	})
	ctx := context.Background()
	vm.Evaluate(ctx, "math.js", `const multiply = (a, b) => a * b`)
	value, _ := vm.Evaluate(ctx, "run.js", "multiply(3, 2)")
	fmt.Println(value)
}
