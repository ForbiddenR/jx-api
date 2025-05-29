package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipGetIntellectChargeRequest struct {
	services.Base
	Data *EquipGetIntellectChargeRequestDetail `json:"data"`
}

type EquipGetIntellectChargeRequestDetail struct {
	EVSE EVSE `json:"evse"`
}

func (s *EquipGetIntellectChargeRequest) Unmarshal(data []byte) error {
	type Alias EquipGetIntellectChargeRequest
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