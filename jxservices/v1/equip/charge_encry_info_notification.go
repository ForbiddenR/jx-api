package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipChargeEncryInfoNotificationRequest{}

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
	return NewEquipChargeEncryInfoNotificationRequest(config.Sn, config.Id, config.Pod, config.MsgId, config.Protocol,
		config.ConnectorId, config.TransactionId, config.EncryptedData, config.MeterNum, config.ElecmeterVersion, config.EncryType)
}

func NewEquipChargeEncryInfoNotificationRequest(sn, id, pod, msgID string, p *services.Protocol,
	connectorId, transactionId, encryptedData, meterNum string, elecmeterVersion int64, encryType uint64) *equipChargeEncryInfoNotificationRequest {
	req := &equipChargeEncryInfoNotificationRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.ChargeEncryInfoNotification.FirstUpper(),
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

var _ services.Response = &equipChargeEncryInfoNotificationResponse{}

type equipChargeEncryInfoNotificationResponse struct {
	api.Response
	Data *equipChargeEncryInfoNotificationResponseDetail `json:"data"`
}

func (resp *equipChargeEncryInfoNotificationResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipChargeEncryInfoNotificationResponse) GetMsg() string {
	return resp.Msg
}

type equipChargeEncryInfoNotificationResponseDetail struct {
}

func ChargeEncryInfoNotificationReqeust(ctx context.Context, req services.Request) error {
	header := services.GetSimpleHeaderValue(services.ChargeEncryInfoNotification)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipChargeEncryInfoNotificationResponse{})
}
