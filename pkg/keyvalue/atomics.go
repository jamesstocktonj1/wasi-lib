package keyvalue

import "github.com/jamesstocktonj1/wasi-lib/pkg/gen/wasi/keyvalue/atomics"

func (b *Bucket) Increment(key string, delta uint64) (uint64, error) {
	value, Err, isErr := atomics.Increment(b.bucket, key, delta).Result()
	if isErr {
		return 0, mapError(Err)
	}
	return value, nil
}
