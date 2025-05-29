package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type ResetType int

const (
	ResetTypeHard ResetType = iota + 1
	ResetTypeSoft
	ResetTypeImmediate
	ResetTypeOnIdle
)

type EquipResetRequest struct {
	services.Base
	Data *EquipResetRequestDetail `json:"data"`
}

type EquipResetRequestDetail struct {
	ResetType ResetType `json:"type" validate:"required"`
}

func (r *EquipResetRequest) UnmarshalJSON(b []byte) error {
	type Alias EquipResetRequest
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
		if aux.Data.ResetType == ResetTypeOnIdle {
			return errors.New(aux.Protocol.String() + ":resetType: idle is invalid")
		}
		if aux.Data.ResetType == ResetTypeImmediate {
			return errors.New(aux.Protocol.String() + ":resetType: immediate is invalid")
		}
	} else {
		if aux.Data.ResetType == ResetTypeHard {
			return errors.New(aux.Protocol.String() + ":resetType: hard is invalid")
		}
		if aux.Data.ResetType == ResetTypeSoft {
			return errors.New(aux.Protocol.String() + ":resetType: soft is invalid")
		}
	}
	return nil
}