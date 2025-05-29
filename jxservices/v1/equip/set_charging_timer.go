package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipSetChargingTimerRequest struct {
	services.Base
	Data *EquipSetChargingTimerRequestDetail `json:"data"`
}

func (r *EquipSetChargingTimerRequest) UnmarshalJSON(b []byte) error {
	type Alias EquipSetChargingTimerRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}

// This struct below is a customized feature of DataTransfer to start charging at a specific time.
type EquipSetChargingTimerRequestDetail struct {
	IdToken        string `json:"idToken" validate:"required"`
	Action         string `json:"action" validate:"-"`
	ConnectorId    string `json:"connectorSerial" validate:"required"`
	VendorCode     string `json:"vendorCode" validate:"required"`
	Timer          string `json:"timer" validate:"-"`
	Enable         bool   `json:"enable" validate:"-"`
	TimerId        int64  `json:"timerId" validate:"required"`
	Version        int64  `json:"version" validate:"required"`
	ChargingTime   *int32 `json:"chargingTime"`
	TimezoneOffset int32  `json:"timezoneOffset" validate:"required"`
}