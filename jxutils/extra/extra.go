package extra

import (
	"encoding/json"

	"github.com/tidwall/sjson"
)

type Extra struct {
	Key   string
	Value any
}

type Option func([]byte) ([]byte, error)

func WithExtra(k, v string) Option {
	return func(b []byte) ([]byte, error) {
		return sjson.SetBytes(b, k, v)

	}
}

func Marshal(v any, opts ...Option) ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return bytes, err
	}
	for _, opt := range opts {
		bytes, err = opt(bytes)
		if err != nil {
			return bytes, err
		}
	}
	return bytes, nil
}
