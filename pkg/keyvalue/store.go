package keyvalue

import (
	"github.com/jamesstocktonj1/wasi-lib/pkg/gen/wasi/keyvalue/store"
	"go.bytecodealliance.org/cm"
)

type Bucket struct {
	bucket store.Bucket
}

func Open(identifier string) (*Bucket, error) {
	bucket, Err, isErr := store.Open(identifier).Result()
	if isErr {
		return nil, mapError(Err)
	}
	return &Bucket{bucket: bucket}, nil
}

func (b *Bucket) Close() error {
	b.bucket.ResourceDrop()
	return nil
}

func (b *Bucket) Get(key string) (string, error) {
	value, Err, isErr := b.bucket.Get(key).Result()
	if isErr {
		return "", mapError(Err)
	} else if value.None() {
		return "", nil
	}
	return string(value.Some().Slice()), nil
}

func (b *Bucket) Set(key, value string) error {
	_, Err, isErr := b.bucket.Set(key, cm.ToList([]byte(value))).Result()
	if isErr {
		return mapError(Err)
	}
	return nil
}

func (b *Bucket) Exists(key string) (bool, error) {
	exists, Err, isErr := b.bucket.Exists(key).Result()
	if isErr {
		return false, mapError(Err)
	}
	return exists, nil
}

func (b *Bucket) Delete(key string) error {
	_, Err, isErr := b.bucket.Delete(key).Result()
	if isErr {
		return mapError(Err)
	}
	return nil
}

func (b *Bucket) ListKeys() ([]string, error) {
	keys, Err, isErr := b.bucket.ListKeys(cm.None[uint64]()).Result()
	if isErr {
		return nil, mapError(Err)
	}
	return keys.Keys.Slice(), nil
}
