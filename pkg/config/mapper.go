package config

import (
	"errors"

	"github.com/jamesstocktonj1/wasi-lib/pkg/gen/wasi/config/runtime"
)

var (
	ErrNotFound = errors.New("not found")
	ErrUpstream = errors.New("upstream error")
	ErrIO       = errors.New("io error")
)

func mapError(err runtime.ConfigError) error {
	switch err.String() {
	case "upstream":
		return ErrUpstream
	default:
		return ErrIO
	}
}
