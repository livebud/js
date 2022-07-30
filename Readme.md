# JS

[![Go Reference](https://pkg.go.dev/badge/github.com/livebud/js.svg)](https://pkg.go.dev/github.com/livebud/js)

The JS package provides a common interface for running Javascript in Go in a consistent environment.

# Examples

## V8

```go
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
```

## Goja

```go
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
```

> Note: This package is still a work in progress. Please see the [issues](https://github.com/livebud/js/issues) for what needs to be done.

## Currently supported Javascript VMs

- [V8](https://github.com/rogchap/v8go)
- [Goja](https://github.com/dop251/goja)

## Goals

- **Swappable JS VMs**: The available VMs each have pros and cons. You should be able to swap VMs out based on your needs.
- **Consistent Runtime**: For each of these VMs, there should be consistent and well-tested globals (e.g. `console`, `setTimeout`, `URL`) that match the web's behavior as much as possible.

## Non-Goals

- **Secure sandbox for user-submitted Javascript**: To provide better performance, the environment is re-used across evaluations. This means that you can set globals to be read in subsequent evaluations. This type of environment is not suitable for user-submitted code.
- **Support non-standard runtime APIs**: There's no plans to add APIs that are specific to certain runtime environments such as Cloudflare Workers, Deno, etc. There's no science to this, but the following heuristic: It should be a [Web API](https://developer.mozilla.org/en-US/docs/Web/API) and be available in Node.js.

## License

MIT
