package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type EquipUpdateFirmwareRequest struct {
	services.Base
	Data *EquipUpdateFirmwareRequestDetail `json:"data"`
}

type EquipUpdateFirmwareRequestDetail struct {
	Location        string   `json:"location"`
	RequestID       *int64   `json:"requestId,omitempty"`
	Retries         *int     `json:"retries,omitempty"`
	RetryInterval   *int     `json:"retryInterval,omitempty"`
	InstallDateTime *int64   `json:"installDateTime,omitempty"`
	Params          []string `json:"params,omitempty"`
	VendorId        *string  `json:"vendorId,omitempty"`
	// add for Yunkuaichong
	ChargerType *uint8  `json:"chargerType,omitempty"`
	Power       *int16  `json:"power,omitempty"`
	Domain      *string `json:"domain,omitempty"`
	Port        *uint16 `json:"port,omitempty"`
	Username    *string `json:"username,omitempty"`
	Password    *string `json:"password,omitempty"`
	Path        *string `json:"path,omitempty"`
	Strategy    *uint8  `json:"strategy,omitempty"`
	Timeout     *uint8  `json:"timeout,omitempty"`
}

func (r *EquipUpdateFirmwareRequest) UnmarshalJSON(b []byte) error {
	type Alias EquipUpdateFirmwareRequest
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
	if aux.Protocol.Equal(services.YunKuaiChong()) {
		if aux.Data.ChargerType == nil {
			return errors.New(aux.Protocol.Name() + ":charger type is nil")
		}
		if aux.Data.Power == nil {
			return errors.New(aux.Protocol.Name() + ":power is nil")
		}
		if aux.Data.Domain == nil {
			return errors.New(aux.Protocol.Name() + ":domain is nil")
		}
		if aux.Data.Port == nil {
			return errors.New(aux.Protocol.Name() + ":port is nil")
		}
		if aux.Data.Username == nil {
			return errors.New(aux.Protocol.Name() + ":username is nil")
		}
		if aux.Data.Password == nil {
			return errors.New(aux.Protocol.Name() + ":password is nil")
		}
		if aux.Data.Path == nil {
			return errors.New(aux.Protocol.Name() + ":path is nil")
		}
		if aux.Data.Strategy == nil {
			return errors.New(aux.Protocol.Name() + ":strategy is nil")
		}
		if aux.Data.Timeout == nil {
			return errors.New(aux.Protocol.Name() + ":timeout is nil")
		}
	} else if !aux.Protocol.Equal(services.OCPP16()) {
		if aux.Data.RequestID == nil {
			return errors.New(aux.Protocol.String() + ":request id is nil")
		}
	}

	return nil
}
