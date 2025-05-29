package equip

import (
	"context"
	"encoding/json"
	"time"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"

	"github.com/makasim/amqpextra/publisher"
	amqp "github.com/rabbitmq/amqp091-go"
)

// This feature is transferred by rabbitmq.

const meterQueue = services.QueuePrefix + "metervalues"

var _ services.Request = &equipMeterValuesRequest{}

type equipMeterValuesRequest struct {
	services.Base
	Data *equipMeterValuesRequestDetail `json:"data"`
}

type equipMeterValuesRequestDetail struct {
	EvseId        *string     `json:"evseSerial,omitempty"`
	Actime        int64       `json:"actime"`
	TransactionId *string     `json:"transactionId,omitempty"`
	ConnectorId   *string     `json:"connectorSerial,omitempty"`
	Timestamp     *int64      `json:"timestamp,omitempty"`
	MeterValue    *MeterValue `json:"meterValue"`
}

func (equipMeterValuesRequest) GetName() services.Request2ServicesNameType {
	return services.MeterValues
}

func (e *equipMeterValuesRequest) TraceId() string {
	return e.MsgID
}

func (equipMeterValuesRequest) IsCallback() bool {
	return false
}

func NewEquipMeterValuesOCPP16Request(sn, pod, msgID string, connectorId string) *equipMeterValuesRequest {
	meterValue := &equipMeterValuesRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    services.OCPP16(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipMeterValuesRequestDetail{
			ConnectorId: &connectorId,
			Actime:      time.Now().Unix(),
			MeterValue:  &MeterValue{},
		},
	}
	return meterValue
}

func NewEquipMeterValuesRequest(sn, pod, msgID string, p *services.Protocol) *equipMeterValuesRequest {
	meterValue := &equipMeterValuesRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Data: &equipMeterValuesRequestDetail{
			MeterValue: &MeterValue{},
		},
	}
	return meterValue
}

func MeterValuesRequest(req *equipMeterValuesRequest, p *publisher.Publisher) error {
	ctx := context.Background()

	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	message := publisher.Message{
		Context: ctx,
		Key:     meterQueue,
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
