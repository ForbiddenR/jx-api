package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipStartTransactionRequest{}

type equipStartTransactionRequest struct {
	services.Base
	Data *equipStartTransactionRequestDetail `json:"data"`
}

type equipStartTransactionRequestDetail struct {
	IdTokenType         IdTokenType   `json:"idTokenType"`
	MeterStart          *int          `json:"meterStart,omitempty"`
	EvseSerial          *string       `json:"evseSerial,omitempty"`
	ConnectorSerial     string        `json:"connectorSerial"`
	ReservationId       *int64        `json:"reservationId,omitempty"`
	TransactionId       *string       `json:"transactionId,omitempty"`
	RemoteStartId       *int64        `json:"remoteStartId,omitempty"`
	Offline             bool          `json:"offline"`
	Timestamp           int64         `json:"timestamp"`
	MeterValue          *MeterValue   `json:"meterValue"`
	Tariff              *Tariff       `json:"tariff,omitempty"`
	ChargingState       uint8         `json:"chargingState"`
	Vin                 *string       `json:"vin,omitempty"`
	RemainingTime       *int          `json:"remainingTime,omitempty"`
	ChargingProfileId   *int64        `json:"chargingProfileId,omitempty"`
	ChargingProfileUnit *uint8        `json:"chargingProfileUnit,omitempty"`
	Temperatures        *Temperatures `json:"temperatures,omitempty"`
}

type StartTransactionRequestConfig struct {
	services.ReusedConfig
	IdToken       string
	ConnectorId   string
	Timestamp     int64
	Offline       bool
	ChargingState uint8
}

func NewEquipStartTransactionRequestWithConfig(config *StartTransactionRequestConfig) *equipStartTransactionRequest {
	req := &equipStartTransactionRequest{
		Base: services.Base{
			EquipmentSn: config.Sn,
			EquipmentId: config.Id,
			Protocol:    config.Protocol,
			Category:    services.StartTransaction.FirstUpper(),
			AccessPod:   config.Pod,
			MsgID:       config.MsgId,
		},
		Data: &equipStartTransactionRequestDetail{
			IdTokenType: IdTokenType{
				IdToken: config.IdToken,
			},
			ConnectorSerial: config.ConnectorId,
			Timestamp:       config.Timestamp,
			Offline:         config.Offline,
			ChargingState:   config.ChargingState,
			Tariff: &Tariff{
				Id: -1,
			},
		},
	}

	if !config.Protocol.Equal(services.OCPP16()) {
		req.Data.MeterValue = &MeterValue{}
	}
	return req
}

func (equipStartTransactionRequest) GetName() services.Request2ServicesNameType {
	return services.StartTransaction
}

func (e *equipStartTransactionRequest) TraceId() string {
	return e.MsgID
}

func (equipStartTransactionRequest) IsCallback() bool {
	return false
}

func NewEquipStartTransactionRequest(sn, id, pod, msgId string, p *services.Protocol,
	idToken string, connectorId string,
	timestamp int64) *equipStartTransactionRequest {
	req := &equipStartTransactionRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.StartTransaction.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipStartTransactionRequestDetail{
			IdTokenType: IdTokenType{
				IdToken: idToken,
			},
			ConnectorSerial: connectorId,
			Timestamp:       timestamp,
		},
	}
	req.Data.Tariff = &Tariff{
		Id: -1,
	}
	if !p.Equal(services.OCPP16()) {
		req.Data.MeterValue = &MeterValue{}
	}
	return req
}

var _ services.Response = &equipStartTransactionResponse{}

type equipStartTransactionResponse struct {
	api.Response
	Data *equipStartTransactionResponseDetail `json:"data"`
}

func (resp *equipStartTransactionResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipStartTransactionResponse) GetMsg() string {
	return resp.Msg
}

type equipStartTransactionResponseDetail struct {
	TransactionId string      `json:"transactionId"`
	IdTokenInfo   IdTokenInfo `json:"idTokenInfo"`
}

func StartTransactionRequest(ctx context.Context, req services.Request) (*equipStartTransactionResponse, error) {
	// // header := services.GetSimpleHeaderValue(services.StartTransaction)

	url := services.GetSimpleURL(req)

	return services.RequestWithResponse(ctx, req, url, &equipStartTransactionResponse{})
}
