package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi/v2"
	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

type equipConfigurationProfileRequest struct {
	services.Base
	Data *equipConfigurationProfileRequestDetail `json:"data"`
}

type equipConfigurationProfileRequestDetail struct {
	TransactionId            string  `json:"transactionId"`
	ConnectorId              string  `json:"connectorSerial"`
	BatteryAllowedMaxVoltage float64 `json:"batteryAllowedMaxVoltage"`
	BMSDynamicBatteryPower   float64 `json:"bmsDynamicBatteryPower"`
	MaxAllowedVoltage        float64 `json:"maxAllowedVoltage"`
	MaxAllowedTemperature    int     `json:"maxAllowedTemperature"`
	Soc                      float64 `json:"soc"`
	CurrentVoltage           float64 `json:"currentVoltage"`
	MaxOutputVoltage         float64 `json:"maxOutputVoltage"`
	MinOutputVoltage         float64 `json:"minOutputVoltage"`
	MaxOutputCurrent         float64 `json:"maxOutputCurrent"`
	MinOutputCurrent         float64 `json:"minOutputCurrent"`
}

func (r *equipConfigurationProfileRequest) GetName() services.Request2ServicesNameType {
	return services.ConfigurationProfile
}

func (r *equipConfigurationProfileRequest) TraceId() string {
	return r.MsgID
}

func (r *equipConfigurationProfileRequest) IsCallback() bool {
	return false
}

func NewEquipConfigurationProfileRequest(sn, pod, msgId string, p *services.Protocol) *equipConfigurationProfileRequest {
	return &equipConfigurationProfileRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    p,
			AccessPod:   pod,
			MsgID:       msgId,
		},
	}
}

type equipConfigurationProfileResponse struct {
	api.Response
}

func (r *equipConfigurationProfileResponse) GetMsg() string {
	return r.Msg
}

func (r *equipConfigurationProfileResponse) GetStatus() int {
	return r.Status
}

func ConfigurationProfileRequest(ctx context.Context, req services.Request) error {
	return services.Transport(ctx, req)
}
