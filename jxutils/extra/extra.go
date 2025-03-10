package extra

import (
	"encoding/json"

	"github.com/tidwall/sjson"
)

type Extra struct {
	Key   string
	Value any
}

func Marshal(v any, extra ...Extra) ([]byte, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return bytes, err
	}
	for _, v := range extra {
		bytes, err = sjson.SetBytes(bytes, v.Key, v.Value)
		if err != nil {
			return bytes, err
		}
	}
	return bytes, nil
}
