package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipSetIntellectChargeRequest struct {
	services.Base
	Data *EquipSetIntellectChargeRequestDetail `json:"data"`
}

type EquipSetIntellectChargeRequestDetail struct {
	VendorId    string      `json:"vendorId"`
	IntellectId int64       `json:"intellectId"`
	StartTime   string      `json:"startTime"`
	StopTime    string      `json:"stopTime"`
	StopEnergy  int64       `json:"stopEnergy"`
	StopSoc     int64       `json:"stopSoc"`
	TimingType  int64       `json:"timingType"`
	Evse        EVSE        `json:"evse"`
	IdTokenType IdTokenType `json:"idTokenType"`
}

func (s *EquipSetIntellectChargeRequest) Unmarshal(data []byte) error {
	type Alias EquipSetIntellectChargeRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Data == nil {
		return errors.New("data is nil")
	}
	return nil
}