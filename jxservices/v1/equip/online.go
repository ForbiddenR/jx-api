package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipOnlineRequest{}

type equipOnlineRequest struct {
	services.Base
	Data *equipOnlineRequestDetail `json:"data"`
}

type equipOnlineRequestDetail struct {
	RemoteAddress *string `json:"remoteAddress"`
}

type OnlineConfig struct {
	services.ReusedConfig
}

func NewEquipOnlineRequestWithConfig(config OnlineConfig) *equipOnlineRequest {
	return NewEquipOnlineRequest(config.Sn, config.Id, config.Pod, config.MsgId, config.Protocol)
}

func NewEquipOnlineRequest(sn, id, pod, msgID string, protocol *services.Protocol) *equipOnlineRequest {
	return &equipOnlineRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    protocol,
			AccessPod:   pod,
			MsgID:       msgID,
		},

		Data: &equipOnlineRequestDetail{},
	}
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

var _ services.Response = &equipOnlineResponse{}

type equipOnlineResponse struct {
	api.Response
	Data *equipOnlineResponseDetail `json:"data"`
}

type equipOnlineResponseDetail struct {
	EquipmentID string `json:"equipmentId" validate:"-"`
	EquipmentSN string `json:"equipmentSN" validate:"-"`
}

func (resp *equipOnlineResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipOnlineResponse) GetMsg() string {
	return resp.Msg
}

func OnlineRequest(ctx context.Context, req *equipOnlineRequest) error {
	// header := services.GetSimpleHeaderValue(services.Online)
	url := services.GetSimpleURL(req)
	return services.RequestWithoutResponse(ctx, req, url, &equipOnlineResponse{})
}