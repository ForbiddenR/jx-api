package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

var _ services.Request = &equipStopTransactionRequest{}

type equipStopTransactionRequest struct {
	services.Base
	Data *equipStopTransactionRequestDetail `json:"data"`
}

type equipStopTransactionRequestDetail struct {
	IdTokenType     *IdTokenType  `json:"idTokenType,omitempty"`
	MeterStop       *int          `json:"meterStop"`
	EvseSerial      *string       `json:"evseSerial,omitempty"`
	ConnectorSerial *string       `json:"connectorSerial,omitempty"`
	ReservationId   *int64        `json:"reservationId,omitempty"`
	TransactionId   string        `json:"transactionId"`
	RemoteStartId   *int64        `json:"remoteStartId,omitempty"`
	Offline         bool          `json:"offline"`
	Timestamp       int64         `json:"timestamp"`
	MeterValue      *MeterValue   `json:"meterValue,omitempty"`
	Tariff          *Tariff       `json:"tariff,omitempty"`
	ChargingState   uint8         `json:"chargingState"`
	Vin             *string       `json:"vin,omitempty"`
	StopReason      int           `json:"stopReason"`
	Temperatures    *Temperatures `json:"temperatures,omitempty"`
}

func (equipStopTransactionRequest) GetName() services.Request2ServicesNameType {
	return services.StopTransaction
}

func (e *equipStopTransactionRequest) TraceId() string {
	return e.MsgID
}

func (equipStopTransactionRequest) IsCallback() bool {
	return false
}

func NewEquipStopTransactionRequest(sn, pod, msgID string, p *services.Protocol, reason int, transactionId string, isOffline bool, timestamp int64) *equipStopTransactionRequest {
	req := &equipStopTransactionRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipStopTransactionRequestDetail{
			IdTokenType:   &IdTokenType{},
			StopReason:    reason,
			TransactionId: transactionId,
			Timestamp:     timestamp,
			Offline:       isOffline,
			MeterValue:    &MeterValue{},
		},
	}
	req.Data.Tariff = &Tariff{
		Id: -1,
	}

	return req
}

var _ services.Response = &equipStopTransactionResponse{}

type equipStopTransactionResponse struct {
	api.Response
	Data *equipStopTransactionResponseDetail `json:"data"`
}

func (resp *equipStopTransactionResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipStopTransactionResponse) GetMsg() string {
	return resp.Msg
}

type equipStopTransactionResponseDetail struct {
	IdTokenInfo
}

func StopTransactionRequest(ctx context.Context, req *equipStopTransactionRequest) (*equipStopTransactionResponse, error) {
	resp := new(equipStopTransactionResponse)
	err := services.TransportWithResp(ctx, req, resp)
	return resp, err
}
