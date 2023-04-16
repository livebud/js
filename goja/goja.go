package goja

import (
	"context"
	"fmt"
	"sync"

	"github.com/dop251/goja"
	"github.com/livebud/js"
)

func New(c *js.Console) *VM {
	vm := goja.New()
	vm.GlobalObject().Set("console", console(c))
	return &VM{sync.Mutex{}, vm}
}

type VM struct {
	mu sync.Mutex
	vm *goja.Runtime
}

var _ js.VM = (*VM)(nil)

// Evaluate the expression. Evaluate can be called concurrently.
func (v *VM) Evaluate(ctx context.Context, path, expr string) (string, error) {
	v.mu.Lock()
	defer v.mu.Unlock()
	result, err := v.vm.RunScript(path, expr)
	if err != nil {
		return "", fmt.Errorf("goja: error evaluating %s. %w", path, err)
	}
	return result.String(), nil
}
