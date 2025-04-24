package jxutils

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	type testStruct struct {
		Duration Duration `json:"duration"`
	}
	tests := []struct {
		name     string
		input    string
		expected Duration
	}{
		{"TestStringType", "\"1m30s\"", Duration(90 * time.Second)},
		{"TestMillisecondsType", "\"300ms\"", Duration(300 * time.Millisecond)},
		{"TestHoursAndMinutesType", "\"2h30m\"", Duration(2*time.Hour + 30*time.Minute)},
		{"TestIntegerType", "10", Duration(10 * time.Second)},
		{"TestFloatType", "10.5", Duration(10 * time.Second)},
		{"TestNullType", "null", 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var data testStruct
			if err := json.Unmarshal([]byte(fmt.Sprintf(`{"duration":%s}`, test.input)), &data); err != nil {
				t.Errorf("Error parsing %s: %v", test.input, err)
			}
			if data.Duration != test.expected {
				t.Errorf("For input %s, expected %v but got %v", test.input, test.expected, data.Duration)
			}
		})
	}
}
