package config

import (
	"sync"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type (
	// Config is the interface of config.
	// It is used to get the value of the key. If the value is not found, it will return the default value.
	Config interface { //nolint:interfacebloat
		// GetString returns the value associated with the key as a string.
		GetString(key string) string
		// GetBool returns the value associated with the key as a boolean.
		GetBool(key string) bool
		// GetInt returns the value associated with the key as an integer.
		GetInt(key string) int
		// GetInt32 returns the value associated with the key as an integer.
		GetInt32(key string) int32
		// GetInt64 returns the value associated with the key as an integer.
		GetInt64(key string) int64
		// GetUint returns the value associated with the key as an unsigned integer.
		GetUint(key string) uint
		// GetUint32 returns the value associated with the key as an unsigned integer.
		GetUint32(key string) uint32
		// GetUint64 returns the value associated with the key as an unsigned integer.
		GetUint64(key string) uint64
		// GetFloat64 returns the value associated with the key as a float64.
		GetFloat64(key string) float64
		// GetTime returns the value associated with the key as time.
		GetTime(key string) time.Time
		// GetDuration returns the value associated with the key as a duration.
		GetDuration(key string) time.Duration
		// GetIntSlice returns the value associated with the key as a slice of int values.
		GetIntSlice(key string) []int

		// GetStringSlice returns the value associated with the key as a slice of strings.
		GetStringSlice(key string) []string
		// GetStringMap returns the value associated with the key as a map of interfaces.
		GetStringMap(key string) map[string]any
	}

	confImpl struct{}
)

var (
	_cache sync.Map //nolint:gochecknoglobals
)

// NewConfig returns a new Config.
func NewConfig() Config {
	return &confImpl{}
}

func valuesCache[T any](key string, getter func(string) T, caster func(any) T) T {
	if value, ok := _cache.Load(key); ok {
		return caster(value)
	}

	value := getter(key)

	_cache.Store(key, value)

	return caster(value)
}

func (*confImpl) GetString(key string) string {
	return valuesCache[string](key, viper.GetString, cast.ToString)
}

func (*confImpl) GetBool(key string) bool {
	return valuesCache(key, viper.GetBool, cast.ToBool)
}

func (*confImpl) GetInt(key string) int {
	return valuesCache(key, viper.GetInt, cast.ToInt)
}

func (*confImpl) GetInt32(key string) int32 {
	return valuesCache(key, viper.GetInt32, cast.ToInt32)
}

func (*confImpl) GetInt64(key string) int64 {
	return valuesCache(key, viper.GetInt64, cast.ToInt64)
}

func (*confImpl) GetUint(key string) uint {
	return valuesCache(key, viper.GetUint, cast.ToUint)
}

func (*confImpl) GetUint32(key string) uint32 {
	return valuesCache(key, viper.GetUint32, cast.ToUint32)
}

func (*confImpl) GetUint64(key string) uint64 {
	return valuesCache(key, viper.GetUint64, cast.ToUint64)
}

func (*confImpl) GetFloat64(key string) float64 {
	return valuesCache(key, viper.GetFloat64, cast.ToFloat64)
}

func (*confImpl) GetTime(key string) time.Time {
	return valuesCache(key, viper.GetTime, cast.ToTime)
}

func (*confImpl) GetDuration(key string) time.Duration {
	return valuesCache(key, viper.GetDuration, cast.ToDuration)
}

func (*confImpl) GetIntSlice(key string) []int {
	return valuesCache(key, viper.GetIntSlice, cast.ToIntSlice)
}

func (*confImpl) GetStringSlice(key string) []string {
	return valuesCache(key, viper.GetStringSlice, cast.ToStringSlice)
}

func (*confImpl) GetStringMap(key string) map[string]any {
	return valuesCache(key, viper.GetStringMap, cast.ToStringMap)
}
