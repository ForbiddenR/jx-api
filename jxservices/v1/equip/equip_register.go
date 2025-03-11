package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipRegisterRequest{}

type equipRegisterRequest struct {
	services.Base
	Data *equipRegisterRequestDetail `json:"data"`
}

type equipRegisterRequestDetail struct {
	RemoteAddress *string `json:"remoteAddress"`
}

func NewEquipRegisterRequest(sn, id, pod, msgID string, protocol *services.Protocol) *equipRegisterRequest {
	return &equipRegisterRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    protocol,
			Category:    services.Register.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},

		Data: &equipRegisterRequestDetail{},
	}
}

func (equipRegisterRequest) GetName() services.Request2ServicesNameType {
	return services.Register
}

func (e *equipRegisterRequest) TraceId() string {
	return e.MsgID
}

func (equipRegisterRequest) IsCallback() bool {
	return false
}

var _ services.Response = &equipRegisterResponse{}

type equipRegisterResponse struct {
	api.Response
	Data *equipRegisterResponseDetail `json:"data"`
}

type equipRegisterResponseDetail struct {
	EquipmentID string `json:"equipmentId" validate:"required"`
	EquipmentSN string `json:"equipmentSN" validate:"required"`
}

func (resp *equipRegisterResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipRegisterResponse) GetMsg() string {
	return resp.Msg
}

func RegisterRequest(ctx context.Context, req *equipRegisterRequest) error {
	header := services.GetSimpleHeaderValue(services.Register)

	url := services.GetSimpleURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipRegisterResponse{})
}
