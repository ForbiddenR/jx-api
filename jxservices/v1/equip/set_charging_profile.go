package equip

import (
	"encoding/json"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipSetChargingProfileRequest struct {
	services.Base
	Data *EquipSetChargingProfileRequestDetail `json:"data"`
}

type EquipSetChargingProfileRequestDetail struct {
	EvseId          string           `json:"evseId"`
	ConnectorId     string           `json:"connectorId"`
	ChargingProfile *ChargingProfile `json:"chargingProfile,omitempty"`
}

type ChargingProfile struct {
	Id               uint64              `json:"id,omitempty"`
	ValidFrom        string              `json:"validFrom,omitempty"`
	ValidTo          string              `json:"validTo,omitempty"`
	TransactionID    string              `json:"transactionID,omitempty"`
	StartTime        string              `json:"startTime,omitempty"`
	ChargingSchedule []*ChargingSchedule `json:"chargingSchedule,omitempty"`
}

type ChargingSchedule struct {
	ChargingRateUnit       int32                         `json:"chargingRateUnit,omitempty"`
	ChargingSchedulePeriod []*ChargingSchedulePeriodType `protobuf:"bytes,2,rep,name=chargingSchedulePeriod,proto3" json:"chargingSchedulePeriod,omitempty"`
}

type ChargingSchedulePeriodType struct {
	From  int64 `json:"from,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	To    int64 `json:"to,omitempty"`
}

func (r *EquipSetChargingProfileRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipSetChargingProfileRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	return nil
}