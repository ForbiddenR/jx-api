package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipSendLocalListRequest struct {
	services.Base
	Data *EquipSendLocalListRequestDetail `json:"data"`
}

func (r *EquipSendLocalListRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipSendLocalListRequest
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
		if aux.Data.IdTokenType.Type == nil {
			return errors.New(aux.Protocol.String() + ":idTokenType.type is nil")
		}
	}
	return nil
}

type EquipSendLocalListRequestDetail struct {
	IdTokenType   IdTokenType  `json:"idTokenType" validate:"required"`
	IdTokenInfo   IdTokenInfo  `json:"idTokenInfo" validate:"required"`
	Version       int          `json:"version" validate:"required"`
	ParentIdToken *string      `json:"parentIdToken,omitempty"`
	GroupIdToken  *IdTokenType `json:"groupIdToken,omitempty"`
}