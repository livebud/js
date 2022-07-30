package js

import (
	"context"
	"io"
)

type VM interface {
	Evaluate(ctx context.Context, path, code string) (string, error)
}

type Console struct {
	Log   io.Writer
	Error io.Writer
}
