package equip

import (
	"context"
	"encoding/json"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
	"github.com/makasim/amqpextra/publisher"
	amqp "github.com/rabbitmq/amqp091-go"
)

type BatteryType uint8

const (
	BatteryTypeLeadAcidBattery BatteryType = iota
	BatteryTypeNickelCadmiumBattery
	BatteryTypePhosphorusSodiumBattery
	BatteryTypePhosphorusIronBattery
	BatteryTypeGalliumArsenideBattery
	BatteryTypeTripleMaterialBattery
	BatteryTypePolymerLiIonBattery
	BatteryTypeTitaniumArsenideBattery
	BatteryTypeOtherBattery
)

const bmsInfoQueue = services.QueuePrefix + "bms_info"

var _ services.Request = &equipBMSInfoRequest{}

type equipBMSInfoRequest struct {
	services.Base
	Data *equipBMSInfoRequestDetail `json:"data"`
}

type equipBMSInfoRequestDetail struct {
	EvseId       *string `json:"evseSerial"`
	ConnectorId  string  `json:"connectorSerial"`
	LimitVoltage float64 `json:"limitVoltage"`
	Version      string  `json:"version"`
	Type         uint8   `json:"type"`
	TotalVoltage float64 `json:"totalVoltage"`
	TotalEnergy  float64 `json:"totalEnergy"`
	Vin          string  `json:"vin"`
}

func (equipBMSInfoRequest) GetName() services.Request2ServicesNameType {
	return services.BMSInfo
}

func (e *equipBMSInfoRequest) TraceId() string {
	return e.MsgID
}

func (equipBMSInfoRequest) IsCallback() bool {
	return false
}

type BMsInfoRequestConfig struct {
	services.ReusedConfig
	ConnectorId  string
	LimitVoltage float64
	Version      string
	Type         uint8
	TotalVoltage float64
	TotalEnergy  float64
	Vin          string
}

func NewEquipBMSInfoRequestWithConfig(config *BMsInfoRequestConfig) *equipBMSInfoRequest {
	return NewEquipBMSInfoRequest(config.Sn, config.Id, config.Pod, config.MsgId, config.Protocol, config.ConnectorId,
		config.LimitVoltage, config.Version, config.Type, config.TotalVoltage, config.TotalEnergy, config.Vin)
}

func NewEquipBMSInfoRequest(sn, id, pod, msgId string, protocol *services.Protocol, connectorId string, limitVoltage float64, version string, ty uint8, totalVoltage, totalEnergy float64, vin string) *equipBMSInfoRequest {
	req := &equipBMSInfoRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    protocol,
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Data: &equipBMSInfoRequestDetail{
			ConnectorId:  connectorId,
			LimitVoltage: limitVoltage,
			Version:      version,
			Type:         ty,
			TotalVoltage: totalVoltage,
			TotalEnergy:  totalEnergy,
			Vin:          vin,
		},
	}
	return req
}

var _ services.Response = &equipBMSInfoResponse{}

type equipBMSInfoResponse struct {
	api.Response
}

func (resp *equipBMSInfoResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipBMSInfoResponse) GetMsg() string {
	return resp.Msg
}

func BMSInfoRequest(ctx context.Context, req services.Request, p *publisher.Publisher) error {
	bytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	messsage := publisher.Message{
		Context: ctx,
		Key:     bmsInfoQueue,
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
