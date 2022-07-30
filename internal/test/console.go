package test

import (
	"fmt"
	"io"

	"github.com/livebud/js"
)

func Console(stdout, stderr io.Writer) js.Console {
	return &console{stdout, stderr}
}

type console struct {
	Stdout io.Writer
	Stderr io.Writer
}

func (c *console) Log(args ...interface{}) {
	fmt.Fprintln(c.Stdout, args...)
}

func (c *console) Error(args ...interface{}) {
	fmt.Fprintln(c.Stderr, args...)
}

func (c *console) Warn(args ...interface{}) {
	fmt.Fprintln(c.Stderr, args...)
}
