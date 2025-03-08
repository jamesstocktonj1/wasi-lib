package config

import (
	"strconv"
	"time"

	"github.com/jamesstocktonj1/wasi-lib/pkg/gen/wasi/config/runtime"
)

func GetString(key string) (string, error) {
	value, Err, isErr := runtime.Get(key).Result()
	if isErr {
		return "", mapError(Err)
	} else if value.None() {
		return "", ErrNotFound
	}
	return *value.Some(), nil
}

func GetStringDefault(key, defaultValue string) string {
	if value, err := GetString(key); err == nil {
		return value
	}
	return defaultValue
}

func GetBool(key string) (bool, error) {
	value, err := GetString(key)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(value)
}

func GetBoolDefault(key string, defaultValue bool) bool {
	if value, err := GetBool(key); err == nil {
		return value
	}
	return defaultValue
}

func GetInt(key string) (int, error) {
	value, err := GetString(key)
	if err != nil {
		return 0, err
	}
	intValue, err := strconv.ParseInt(value, 10, 64)
	return int(intValue), err
}

func GetIntDefault(key string, defaultValue int) int {
	if value, err := GetInt(key); err == nil {
		return value
	}
	return defaultValue
}

func GetFloat(key string) (float64, error) {
	value, err := GetString(key)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(value, 64)
}

func GetFloatDefault(key string, defaultValue float64) float64 {
	if value, err := GetFloat(key); err == nil {
		return value
	}
	return defaultValue
}

func GetDuration(key string) (time.Duration, error) {
	value, err := GetString(key)
	if err != nil {
		return time.Second, err
	}
	return time.ParseDuration(value)
}

func GetDurationDefault(key string, defaultValue time.Duration) time.Duration {
	if value, err := GetDuration(key); err == nil {
		return value
	}
	return defaultValue
}
