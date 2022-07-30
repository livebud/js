package test

import (
	"io"

	"github.com/livebud/js"
)

func Console(stdout, stderr io.Writer) *js.Console {
	return &js.Console{
		Log:   stdout,
		Error: stderr,
	}
}
