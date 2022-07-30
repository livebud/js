package js

import "context"

type VM interface {
	Evaluate(ctx context.Context, path, code string) (string, error)
}

type Console interface {
	Log(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}
