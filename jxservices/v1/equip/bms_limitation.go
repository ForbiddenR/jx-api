package equip

import (
	"context"
	"encoding/json"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
	"github.com/makasim/amqpextra/publisher"
	amqp "github.com/rabbitmq/amqp091-go"
)

var _ services.Request = &equipBMSLimitRequest{}

type BMSLimitRequestInterface interface {
	services.Request
	Construct(BMSLimitRequestConfig)
}

const bmsLimitQueue = services.QueuePrefix + "bms"

type equipBMSLimitRequest struct {
	services.Base
	Data *equipBMSLimitRequestDetail `json:"data"`
}

type equipBMSLimitRequestDetail struct {
	EvseId           *string `json:"evseSerial,omitempty"`
	ConnectorId      string  `json:"connectorSerial"`
	MaxTemp          uint64  `json:"maxTemp"`
	MaxCurrent       float64 `json:"maxCurrent"`
	MaxVoltage       float64 `json:"maxVoltage"`
	MaxOutputVoltage float64 `json:"maxOutputVoltage"`
	MaxOutputCurrent float64 `json:"maxOutputCurrent"`
	Soc              float64 `json:"soc"`
	Capacity         float64 `json:"capacity"`
	Vin              string  `json:"vin"`
	Prepare          bool    `json:"prepare"`
}

func (equipBMSLimitRequest) GetName() services.Request2ServicesNameType {
	return services.BMSLimit
}

func (e *equipBMSLimitRequest) TraceId() string {
	return e.MsgID
}

func (equipBMSLimitRequest) IsCallback() bool {
	return false
}

type BMSLimitRequestConfig struct {
	services.ReusedConfig
	ConnectorId      string
	MaxTemp          uint64
	MaxCurrent       float64
	MaxVoltage       float64
	MaxOutputVoltag  float64
	MaxOutputCurrent float64
	Soc              float64
	Capacity         float64
	Vin              string
	Prepare          bool
}

func NewEquipBMSLimitRequestWithConfig(config *BMSLimitRequestConfig) *equipBMSLimitRequest {
	return &equipBMSLimitRequest{
		Base: services.Base{
			EquipmentSn: config.Sn,
			EquipmentId: config.Id,
			Protocol:    config.Protocol,
			Category:    services.BMSLimit.FirstUpper(),
			AccessPod:   config.Pod,
			MsgID:       config.MsgId,
		},
		Data: &equipBMSLimitRequestDetail{
			ConnectorId:      config.ConnectorId,
			MaxTemp:          config.MaxTemp,
			MaxCurrent:       config.MaxCurrent,
			MaxVoltage:       config.MaxVoltage,
			MaxOutputVoltage: config.MaxOutputVoltag,
			MaxOutputCurrent: config.MaxOutputCurrent,
			Soc:              config.Soc,
			Capacity:         config.Capacity,
			Vin:              config.Vin,
			Prepare:          config.Prepare,
		},
	}
}

func NewEquipBMSLimitRequest(sn, id, pod, msgId string, protocol *services.Protocol, connecorId string, maxTemp uint64,
	maxCurrent, maxVoltage, maxOutputVoltage, maxOutputCurrent, soc, capacity float64, vin string, prepare bool) *equipBMSLimitRequest {
	req := &equipBMSLimitRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    protocol,
			Category:    services.BMSLimit.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipBMSLimitRequestDetail{
			ConnectorId:      connecorId,
			MaxTemp:          maxTemp,
			MaxCurrent:       maxCurrent,
			MaxVoltage:       maxVoltage,
			MaxOutputVoltage: maxOutputVoltage,
			MaxOutputCurrent: maxOutputCurrent,
			Soc:              soc,
			Capacity:         capacity,
			Vin:              vin,
			Prepare:          prepare,
		},
	}
	return req
}

var _ services.Response = &equipBMSLimitResponse{}

type equipBMSLimitResponse struct {
	api.Response
}

func (resp *equipBMSLimitResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipBMSLimitResponse) GetMsg() string {
	return resp.Msg
}

func BMSLimitRequest(ctx context.Context, req services.Request, p *publisher.Publisher) error {
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	messsage := publisher.Message{
		Context: ctx,
		Key:     bmsLimitQueue,
		Publishing: amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		},
	}

	err = p.Publish(messsage)

	if err != nil {
		return err
	}
	return nil
}
