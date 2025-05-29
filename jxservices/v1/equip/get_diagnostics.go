package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipGetDiagnosticsRequest struct {
	services.Base
	Data *EquipGetDiagnosticsRequestDetail `json:"data"`
}

type EquipGetDiagnosticsRequestDetail struct {
	Location      string `json:"location" validate:"required"`
	RequestID     *int64 `json:"requestId,omitempty"`
	Retries       *int   `json:"retries,omitempty"`
	RetryInterval *int   `json:"retryInterval,omitempty"`
	StartTime     *int64 `json:"startTime,omitempty"`
	StopTime      *int64 `json:"stopTime,omitempty"`
}

func (r *EquipGetDiagnosticsRequest) UnmarshalJSON(b []byte) error {
	type Alias EquipGetDiagnosticsRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	if aux.Data == nil {
		return errors.New("data is nil")
	}
	if !aux.Protocol.Equal(services.OCPP16()) {
		if aux.Data.RequestID == nil {
			return errors.New(aux.Protocol.String() + ":request id is nil")
		}
	}
	return nil
}