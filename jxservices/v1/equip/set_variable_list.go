package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipSetVariableListRequest struct {
	services.Base
	Data []EquipSetVariableListRequestDetail `json:"data"`
}

func (r *EquipSetVariableListRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipSetVariableListRequest
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
		if len(aux.Data) == 0 {
			return errors.New(aux.Protocol.String() + ":data is nil")
		}
	}
	return nil
}

type EquipSetVariableListRequestDetail struct {
	Component Component `json:"component" validate:"required"`
	Key       string    `json:"key" validate:"required"`
	Value     string    `json:"value" validate:"required"`
}