package equip

import (
	"context"

	services "github.com/ForbiddenR/jxapi/v2/jxservices"
)

var _ services.Request = &equipOnlineRequest{}

type equipOnlineRequest struct {
	services.Base
	Data *equipOnlineRequestDetail `json:"data"`
}

type equipOnlineRequestDetail struct {
	RemoteAddress *string `json:"remoteAddress"`
}

func (equipOnlineRequest) GetName() services.Request2ServicesNameType {
	return services.Online
}

func (e *equipOnlineRequest) TraceId() string {
	return e.MsgID
}

func (equipOnlineRequest) IsCallback() bool {
	return false
}

func NewEquipOnlineRequest(sn string, protocol *services.Protocol, pod, msgID string) *equipOnlineRequest {
	return &equipOnlineRequest{
		Base: services.Base{
			EquipmentSn: sn,
			Protocol:    protocol,
			AccessPod:   pod,
			MsgID:       msgID,
		},

		Data: &equipOnlineRequestDetail{},
	}
}

type OnlineConfig struct {
	services.ReusedConfig
}

func NewEquipOnlineRequestWithConfig(config OnlineConfig) *equipOnlineRequest {
	return NewEquipOnlineRequest(config.Sn, config.Protocol, config.Pod, config.MsgID)
}

func OnlineRequest(ctx context.Context, req *equipOnlineRequest) error {
	return services.Transport(ctx, req)
}
