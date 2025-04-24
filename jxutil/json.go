package jxutils

import (
	"encoding/json"
	"fmt"
	"time"
)

type Duration time.Duration

func (d *Duration) UnmarshalJSON(b []byte) error {
	var value any
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	switch v := value.(type) {
	case float64:
		*d = Duration(time.Duration(v) * time.Second)
		return nil
	case string:
		duration, err := time.ParseDuration(v)
		if err != nil {
			return err
		}
		*d = Duration(duration)
		return nil
	case nil:
		*d = 0
		return nil
	default:
		return fmt.Errorf("invalid duration: %v", value)
	}
}
