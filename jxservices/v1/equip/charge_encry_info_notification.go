package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipChargeEncryInfoNotificationRequest struct {
	services.Base
	Data *equipChargeEncryInfoNotificationRequestDetail `json:"data"`
}

type equipChargeEncryInfoNotificationRequestDetail struct {
	EvseId           *string `json:"evseSerial"`
	ConnectorId      string  `json:"connectorSerial"`
	TransactionId    string  `json:"transactionId"`
	EncryptedData    string  `json:"encryptedData"`
	MeterNum         string  `json:"meterNum"`
	ElecmeterVersion int64   `json:"elecmeterVersion"`
	EncryType        uint64  `json:"encryType"`
}

func (equipChargeEncryInfoNotificationRequest) GetName() services.Request2ServicesNameType {
	return services.ChargeEncryInfoNotification
}

func (e *equipChargeEncryInfoNotificationRequest) TraceId() string {
	return e.MsgID
}

func (equipChargeEncryInfoNotificationRequest) IsCallback() bool {
	return false
}

type ChargeEncryInfoNotificationRequestConfig struct {
	services.ReusedConfig
	ConnectorId      string
	TransactionId    string
	EncryptedData    string
	MeterNum         string
	ElecmeterVersion int64
	EncryType        uint64
}

func NewEquipChargeEncryInfoNotificationRequestWithConfig(config *ChargeEncryInfoNotificationRequestConfig) *equipChargeEncryInfoNotificationRequest {
	return NewEquipChargeEncryInfoNotificationRequest(config.Sn, config.Protocol, config.Pod, config.MsgID,
		config.ConnectorId, config.TransactionId, config.EncryptedData, config.MeterNum, config.ElecmeterVersion, config.EncryType)
}

func NewEquipChargeEncryInfoNotificationRequest(sn string, p *services.Protocol, pod, msgID string,
	connectorId, transactionId, encryptedData, meterNum string, elecmeterVersion int64, encryType uint64) *equipChargeEncryInfoNotificationRequest {
	req := &equipChargeEncryInfoNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipChargeEncryInfoNotificationRequestDetail{
			ConnectorId:      connectorId,
			TransactionId:    transactionId,
			EncryptedData:    encryptedData,
			MeterNum:         meterNum,
			ElecmeterVersion: elecmeterVersion,
			EncryType:        encryType,
		},
	}
	return req
}

func ChargeEncryInfoNotificationReqeust(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
