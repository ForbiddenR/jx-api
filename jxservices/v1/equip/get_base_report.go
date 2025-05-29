package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type GetBaseReportType int

const (
	GetBaseReportTypeConfiguration GetBaseReportType = iota
	GetBaseReportTypeFull
	GetBaseReportTypeSummary
)

type EquipGetBaseReportRequest struct {
	services.Base
	Data *EquipGetBaseReportRequestDetail `json:"data"`
}

type EquipGetBaseReportRequestDetail struct {
	Keys      []string           `json:"keys,omitempty"`
	RequestId *uint64            `json:"requestId,omitempty"`
	Type      *GetBaseReportType `json:"type,omitempty"`
}

func (r *EquipGetBaseReportRequest) UnmarshalJSON(b []byte) error {
	type Alias EquipGetBaseReportRequest
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

	if aux.Protocol.Equal(services.OCPP16()) {
		if len(aux.Data.Keys) == 0 {
			return errors.New("keys is nil")
		}
	} else {
		if aux.Data.RequestId == nil {
			return errors.New("requestId is nil")
		}
		if aux.Data.Type == nil {
			return errors.New("type is nil")
		}
	}
	return nil
}