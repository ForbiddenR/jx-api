package equip

import (
	"context"
	"encoding/json"
	"time"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
	"github.com/ForbiddenR/toolkit/transport"
	"github.com/makasim/amqpextra/publisher"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tidwall/sjson"
)

const updateTransactionQueue = services.QueuePrefix + "transaction"

var _ services.Request = &equipUpdateTransactionRequest{}

type equipUpdateTransactionRequest struct {
	services.Base
	Data *equipUpdateTransactionReqeustDetail `json:"data"`
}

func (equipUpdateTransactionRequest) GetName() services.Request2ServicesNameType {
	return services.UpdateTransaction
}

func (e *equipUpdateTransactionRequest) TraceId() string {
	return e.MsgID
}

func (equipUpdateTransactionRequest) IsCallback() bool {
	return false
}

type equipUpdateTransactionReqeustDetail struct {
	Actime        int64         `json:"actime"`
	TransactionId string        `json:"transactionId"`
	EvseId        *string       `json:"evseSerial,omitempty"`
	ConnectorId   string        `json:"connectorSerial"`
	Offline       bool          `json:"offline"`
	Timestamp     int64         `json:"timestamp"`
	MeterValue    *MeterValue   `json:"meterValue,omitempty"`
	Tariff        *Tariff       `json:"tariff,omitempty"`
	ChargingState uint8         `json:"chargingState"`
	RemainingTime *int          `json:"remainingTime,omitempty"`
	VIN           *string       `json:"vin,omitempty"`
	Temperatures  *Temperatures `json:"temperatures,omitempty"`
}

func NewUpdateTransactionRequest(sn, pod, msgID string, p *services.Protocol, transactionId, connectorId string, offline bool, timestamp int64, chargeState uint8) *equipUpdateTransactionRequest {
	updateTransaction := &equipUpdateTransactionRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipUpdateTransactionReqeustDetail{
			Actime:        time.Now().Unix(),
			TransactionId: transactionId,
			ConnectorId:   connectorId,
			Offline:       offline,
			Timestamp:     timestamp,
			MeterValue:    &MeterValue{},
			Tariff: &Tariff{
				Id: -1,
			},
			ChargingState: chargeState,
		},
	}

	return updateTransaction
}

func UpdateTransactionReqeust(ctx context.Context, req *equipUpdateTransactionRequest, p *publisher.Publisher) error {
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
	bytes, err = sjson.SetBytes(bytes, services.IdHandle.Value(), transport.EquipmentIdFromCtx(ctx))
	if err != nil {
		return err
	}
	api.Log.Info("send request to services. queue: " + updateTransactionQueue + " data: " + string(bytes))
	message := publisher.Message{
		Context: ctx,
		Key:     updateTransactionQueue,
		Publishing: amqp.Publishing{
			Headers:     amqp.Table{"TraceId": req.MsgID},
			ContentType: "application/json",
			Body:        bytes,
		},
	}
	err = p.Publish(message)
	if err != nil {
		return err
	}
	return nil
}
