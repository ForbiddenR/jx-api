package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipOffPeakChargeRequest struct {
	services.Base
	Data *EquipOffPeakChargeRequestDetail `json:"data"`
}

func (r *EquipOffPeakChargeRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipOffPeakChargeRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Data == nil {
		return errors.New("data is nil")
	}
	if aux.Protocol.Equal(services.OCPP16()) || aux.Protocol.Equal(services.OCPP201()) {
		return errors.New("off-peak charge is not supported by " + aux.Protocol.String())
	}
	return nil
}

type EquipOffPeakChargeRequestDetail struct {
	EVSE
	IdTokenType   IdTokenType `json:"idTokenType" validate:"required"`
	RemoteStartId int64       `json:"remoteStartId" validate:"required"`
	OffPeak       OffPeak     `json:"offPeak" validate:"required"`
}