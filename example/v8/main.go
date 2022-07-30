package main

import (
	"context"
	"fmt"
	"os"

	"github.com/livebud/js"
	v8 "github.com/livebud/js/v8"
)

func main() {
	vm, _ := v8.Load(&js.Console{
		Log:   os.Stdout,
		Error: os.Stderr,
	})
	defer vm.Close()
	ctx := context.Background()
	vm.Evaluate(ctx, "math.js", `const multiply = (a, b) => a * b`)
	value, _ := vm.Evaluate(ctx, "run.js", "multiply(3, 2)")
	fmt.Println(value)
}
