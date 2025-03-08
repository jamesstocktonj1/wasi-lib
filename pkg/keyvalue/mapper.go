package keyvalue

import (
	"errors"

	"github.com/jamesstocktonj1/wasi-lib/pkg/gen/wasi/keyvalue/store"
)

var (
	ErrNoSuchStore  = errors.New("no such store")
	ErrAccessDenied = errors.New("access denied")
)

func mapError(err store.Error) error {
	switch err.String() {
	case "no-such-store":
		return ErrNoSuchStore
	case "access-denied":
		return ErrAccessDenied
	default:
		return errors.New(*err.Other())
	}
}
