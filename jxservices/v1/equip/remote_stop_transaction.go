package equip

import (
	"encoding/json"
	"errors"

	services "github.com/ForbiddenR/jxapi/jxservices"
)

type EquipRemoteStopTransactionRequest struct {
	services.Base
	Data *EquipRemoteStopTransactionRequestDetail `json:"data"`
}

type EquipRemoteStopTransactionRequestDetail struct {
	TransactionId string  `json:"transactionId" validate:"required"`
	EvseId        *string `json:"evseSerial"`
	ConnectorId   *string `json:"connectorSerial"`
}

func (r *EquipRemoteStopTransactionRequest) UnmarshalJSON(data []byte) error {
	type Alias EquipRemoteStopTransactionRequest
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
	if aux.Protocol.Equal(services.OCPP201()) {
		if aux.Data.EvseId == nil {
			return errors.New(aux.Protocol.Name + ":evse serial is nil")
		}
		if aux.Data.ConnectorId == nil {
			return errors.New(aux.Protocol.Name + ":connector serial is nil")
		}
	} else if !aux.Protocol.Equal(services.OCPP16()) {
		if aux.Data.ConnectorId == nil {
			return errors.New(aux.Protocol.Name + ":connector serial is nil")
		}
	}

	return nil
}