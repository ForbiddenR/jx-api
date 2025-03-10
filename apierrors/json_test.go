package apierrors

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ISO8601 time format, assuming Zulu timestamp.
const ISO8601 = "2006-01-02T15:04:05Z"

// DateTimeFormat to be used for all OCPP messages.
//
// The default dateTime format is RFC3339.
// Change this if another format is desired.
var DateTimeFormat = time.RFC3339

// DateTime wraps a time.Time struct, allowing for improved dateTime JSON compatibility.
type DateTime struct {
	time.Time
}

func NewDateTimeFromUnix(t int64) *DateTime {
	return NewDateTime(time.Unix(t, 0))
}

// Creates a new DateTime struct, embedding a time.Time struct.
func NewDateTime(time time.Time) *DateTime {
	return &DateTime{Time: time}
}

func (dt *DateTime) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	if DateTimeFormat == "" {
		defaultTime := time.Time{}
		err := json.Unmarshal(input, &defaultTime)
		if err != nil {
			return WrapFormatValueJsonError(err)
		}
		dt.Time = defaultTime.Local()
	} else {
		newTime, err := time.Parse(DateTimeFormat, strInput)
		if err != nil {
			return WrapFormatValueJsonError(err)
		}
		dt.Time = newTime.Local()
	}
	return nil
}

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	if DateTimeFormat == "" {
		return json.Marshal(dt.Time)
	}
	timeStr := dt.FormatTimestamp()
	return json.Marshal(timeStr)
}

// Formats the UTC timestamp using the DateTimeFormat setting.
// This function is used during JSON marshaling as well.
func (dt *DateTime) FormatTimestamp() string {
	return dt.UTC().Format(DateTimeFormat)
}

func FormatTimestamp(t time.Time) string {
	return t.UTC().Format(DateTimeFormat)
}

type StatusNotificationJsonErrorCode string

const StatusNotificationJsonErrorCodeConnectorLockFailure StatusNotificationJsonErrorCode = "ConnectorLockFailure"
const StatusNotificationJsonErrorCodeEVCommunicationError StatusNotificationJsonErrorCode = "EVCommunicationError"
const StatusNotificationJsonErrorCodeGroundFailure StatusNotificationJsonErrorCode = "GroundFailure"
const StatusNotificationJsonErrorCodeHighTemperature StatusNotificationJsonErrorCode = "HighTemperature"
const StatusNotificationJsonErrorCodeInternalError StatusNotificationJsonErrorCode = "InternalError"
const StatusNotificationJsonErrorCodeLocalListConflict StatusNotificationJsonErrorCode = "LocalListConflict"
const StatusNotificationJsonErrorCodeNoError StatusNotificationJsonErrorCode = "NoError"
const StatusNotificationJsonErrorCodeOtherError StatusNotificationJsonErrorCode = "OtherError"
const StatusNotificationJsonErrorCodeOverCurrentFailure StatusNotificationJsonErrorCode = "OverCurrentFailure"
const StatusNotificationJsonErrorCodeOverVoltage StatusNotificationJsonErrorCode = "OverVoltage"
const StatusNotificationJsonErrorCodePowerMeterFailure StatusNotificationJsonErrorCode = "PowerMeterFailure"
const StatusNotificationJsonErrorCodePowerSwitchFailure StatusNotificationJsonErrorCode = "PowerSwitchFailure"
const StatusNotificationJsonErrorCodeReaderFailure StatusNotificationJsonErrorCode = "ReaderFailure"
const StatusNotificationJsonErrorCodeResetFailure StatusNotificationJsonErrorCode = "ResetFailure"
const StatusNotificationJsonErrorCodeUnderVoltage StatusNotificationJsonErrorCode = "UnderVoltage"
const StatusNotificationJsonErrorCodeWeakSignal StatusNotificationJsonErrorCode = "WeakSignal"

type StatusNotificationJsonStatus string

const StatusNotificationJsonStatusAvailable StatusNotificationJsonStatus = "Available"
const StatusNotificationJsonStatusCharging StatusNotificationJsonStatus = "Charging"
const StatusNotificationJsonStatusFaulted StatusNotificationJsonStatus = "Faulted"
const StatusNotificationJsonStatusFinishing StatusNotificationJsonStatus = "Finishing"
const StatusNotificationJsonStatusPreparing StatusNotificationJsonStatus = "Preparing"
const StatusNotificationJsonStatusReserved StatusNotificationJsonStatus = "Reserved"
const StatusNotificationJsonStatusSuspendedEV StatusNotificationJsonStatus = "SuspendedEV"
const StatusNotificationJsonStatusSuspendedEVSE StatusNotificationJsonStatus = "SuspendedEVSE"
const StatusNotificationJsonStatusUnavailable StatusNotificationJsonStatus = "Unavailable"

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusNotificationJson) UnmarshalJSON(b []byte) error {
	var raw map[string]any
	if err := json.Unmarshal(b, &raw); err != nil {
		return WrapSyntaxJsonError(err)
	}
	if v, ok := raw["connectorId"]; !ok || v == nil {
		return WrapRequiredJsonError(fmt.Errorf("field connectorId in StatusNotificationJson: required"))
	}
	if v, ok := raw["errorCode"]; !ok || v == nil {
		return WrapRequiredJsonError(fmt.Errorf("field errorCode in StatusNotificationJson: required"))
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return WrapRequiredJsonError(fmt.Errorf("field status in StatusNotificationJson: required"))
	}
	type Plain StatusNotificationJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		if _, ok := IsJsonError(err); ok {
			return err
		} else {
			return WrapSyntaxJsonError(err)
		}
	}
	*j = StatusNotificationJson(plain)
	return nil
}

type StatusNotificationJson struct {
	// ConnectorId corresponds to the JSON schema field "connectorId".
	ConnectorId uint8 `json:"connectorId" yaml:"connectorId"`

	// ErrorCode corresponds to the JSON schema field "errorCode".
	ErrorCode StatusNotificationJsonErrorCode `json:"errorCode" yaml:"errorCode"`

	// Info corresponds to the JSON schema field "info".
	Info *string `json:"info,omitempty" yaml:"info,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status StatusNotificationJsonStatus `json:"status" yaml:"status"`

	// Timestamp corresponds to the JSON schema field "timestamp".
	Timestamp *DateTime `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`

	// VendorErrorCode corresponds to the JSON schema field "vendorErrorCode".
	VendorErrorCode *string `json:"vendorErrorCode,omitempty" yaml:"vendorErrorCode,omitempty"`

	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId *string `json:"vendorId,omitempty" yaml:"vendorId,omitempty"`
}

var enumValues_StatusNotificationJsonStatus = []any{
	"Available",
	"Preparing",
	"Charging",
	"SuspendedEVSE",
	"SuspendedEV",
	"Finishing",
	"Reserved",
	"Unavailable",
	"Faulted",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusNotificationJsonStatus) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return WrapEnumValueJsonError(err)
	}
	var ok bool
	for _, expected := range enumValues_StatusNotificationJsonStatus {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return WrapEnumValueJsonError(fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_StatusNotificationJsonStatus, v))
	}
	*j = StatusNotificationJsonStatus(v)
	return nil
}

var enumValues_StatusNotificationJsonErrorCode = []any{
	"ConnectorLockFailure",
	"EVCommunicationError",
	"GroundFailure",
	"HighTemperature",
	"InternalError",
	"LocalListConflict",
	"NoError",
	"OtherError",
	"OverCurrentFailure",
	"PowerMeterFailure",
	"PowerSwitchFailure",
	"ReaderFailure",
	"ResetFailure",
	"UnderVoltage",
	"OverVoltage",
	"WeakSignal",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusNotificationJsonErrorCode) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return WrapSyntaxJsonError(err)
	}
	var ok bool
	for _, expected := range enumValues_StatusNotificationJsonErrorCode {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return WrapEnumValueJsonError(fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_StatusNotificationJsonErrorCode, v))
	}
	*j = StatusNotificationJsonErrorCode(v)
	return nil
}

func TestJsonError(t *testing.T) {

	testSyntaxJson := []byte(`{"connectorId": 1, }`)
	err := json.Unmarshal(testSyntaxJson, &StatusNotificationJson{})
	assert.IsType(t, &json.SyntaxError{}, err)

	testRequiredJson := []byte(`{"connectorId": 1}`)
	err = json.Unmarshal(testRequiredJson, &StatusNotificationJson{})
	assert.IsType(t, &JsonError{}, err)
	assert.Equal(t, RequiredJsonError, err.(*JsonError).t)

	testFormatValueJson := []byte(`{"connectorId":1, "errorCode": "NoError", "status": "Preparing", "timestamp": "2022123123"}`)
	err = json.Unmarshal(testFormatValueJson, &StatusNotificationJson{})
	assert.IsType(t, &JsonError{}, err)
	assert.Equal(t, FormatValueJsonError, err.(*JsonError).t)

	testEnumValueJson := []byte(`{"connectorId":1, "errorCode": "NoError", "status": "Preparin"}`)
	err = json.Unmarshal(testEnumValueJson, &StatusNotificationJson{})
	assert.IsType(t, &JsonError{}, err)
	assert.Equal(t, EnumValueJsonError, err.(*JsonError).t)

}
