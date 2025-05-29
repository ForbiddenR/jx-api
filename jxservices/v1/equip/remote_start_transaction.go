package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipRemoteStartTransactionRequest struct {
	services.Base
	Data *EquipRemoteStartTransactionRequestDetail `json:"data"`
}

type EquipRemoteStartTransactionRequestDetail struct {
	EvseId        *string     `json:"evseSerial,omitempty"`
	ConnectorId   string      `json:"connectorSerial" validate:"required"`
	IdTokenType   IdTokenType `json:"idTokenType" validate:"required"`
	RemoteStartId *uint64     `json:"remoteStartId,omitempty"`
	Intellect     *Intellect  `json:"intellect,omitempty"`
}

func (r *EquipRemoteStartTransactionRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipRemoteStartTransactionRequest
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
		if aux.Data.EvseId == nil {
			return errors.New(aux.Protocol.Name + ":evse serial is nil")
		}
		if aux.Data.RemoteStartId == nil {
			return errors.New(aux.Protocol.Name + ":remote start id is nil")
		}
		if aux.Data.IdTokenType.Type == nil {
			return errors.New(aux.Protocol.Name + ":id token type is nil")
		}
	}
	return nil
}