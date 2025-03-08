package config

import (
	"errors"

	"github.com/jamesstocktonj1/wasi-lib/pkg/gen/wasi/config/runtime"
)

var (
	ErrNotFound = errors.New("not found")
)

func mapError(err runtime.ConfigError) error {
	return nil
}
