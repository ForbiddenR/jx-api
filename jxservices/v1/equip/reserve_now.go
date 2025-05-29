package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipReserveNowRequest struct {
	services.Base
	Data *EquipReserveNowRequestDetail `json:"data"`
}

func (r *EquipReserveNowRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipReserveNowRequest
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
			return errors.New(aux.Protocol.String() + ":id token type is nil")
		}
	} else {
		if aux.Data.ConnectorID == nil {
			return errors.New(aux.Protocol.String() + ":connector id is nil")
		}
	}
	return nil
}

type EquipReserveNowRequestDetail struct {
	ID            int64        `json:"reserveId" validate:"required"`
	ExpireDate    int64        `json:"expireDate" validate:"required"`
	ConnectorID   *string      `json:"connectorSerial,omitempty"`
	EvseID        *string      `json:"evseSerial,omitempty"`
	IdTokenType   IdTokenType  `json:"idTokenType"`
	ParentIdToken *string      `json:"parentIdToken,omitempty"`
	GroupIdToken  *IdTokenType `json:"groupIdToken,omitempty"`
}