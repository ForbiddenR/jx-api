package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipGetVariablesRequest struct {
	services.Base
	Data *EquipGetVariablesRequestDetail `json:"data"`
}

func (r *EquipGetVariablesRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipGetVariablesRequest
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
	if !aux.Protocol.Equal(services.OCPP16()) {
		if aux.Data.Component == nil {
			return errors.New(aux.Protocol.String() + ":component is nil")
		}
	}
	return nil
}

type EquipGetVariablesRequestDetail struct {
	Component *Component `json:"component,omitempty"`
	Key       string     `json:"key" validate:"required"`
}