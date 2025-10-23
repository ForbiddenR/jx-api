package equip

import (
	"context"
	"encoding/json"
	"time"

	services "github.com/ForbiddenR/jxapi/jxservices"

	"github.com/makasim/amqpextra/publisher"
	amqp "github.com/rabbitmq/amqp091-go"
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
	ChargingState *uint8        `json:"chargingState,omitempty"`
	RemainingTime *int          `json:"remainingTime,omitempty"`
	VIN           *string       `json:"vin,omitempty"`
	Temperatures  *Temperatures `json:"temperatures,omitempty"`
}

func NewUpdateTransactionRequest(sn, id, pod, msgId string, p *services.Protocol, transactionId, connectorId string, offline bool, timestamp int64, chargeState uint8) *equipUpdateTransactionRequest {
	updateTransaction := &equipUpdateTransactionRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
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
			// ChargingState: chargeState,
		},
	}
	if chargeState != maxUint8 {
		updateTransaction.Data.ChargingState = &chargeState
	}
	return updateTransaction
}

func UpdateTransactionReqeust(req *equipUpdateTransactionRequest, p *publisher.Publisher) error {
	ctx := context.Background()
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}
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
