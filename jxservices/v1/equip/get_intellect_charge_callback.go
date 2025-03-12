package equip

import (
	"context"

	api "github.com/ForbiddenR/jxapi"
	"github.com/ForbiddenR/jxapi/apierrors"
	services "github.com/ForbiddenR/jxapi/jxservices"
)

var _ services.Request = &equipGetIntellectChargeRequest{}

type equipGetIntellectChargeRequest struct {
	services.Base
	Callback services.CB                           `json:"callback"`
	Data     *equipGetIntellectChargeRequestDetail `json:"data"`
}

type equipGetIntellectChargeRequestDetail struct {
	services.CB
	EVSE           EVSE    `json:"evse"`
	IntellectType  uint8   `json:"type"`
	IntellectId    string  `json:"intellectId"`
	StartTime      string  `json:"startTime"`
	EndTime        *string `json:"endTime"`
	EndElectricity *int    `json:"endElectricity"`
	EndSoc         *int    `json:"endSOC"`
	Status         uint8   `json:"status"`
}

func (*equipGetIntellectChargeRequest) GetName() services.Request2ServicesNameType {
	return services.GetIntellectCharge
}

func (e *equipGetIntellectChargeRequest) TraceId() string {
	return e.MsgID
}

func (equipGetIntellectChargeRequest) IsCallback() bool {
	return true
}

func NewEquipGetIntellectChargeCallbackRequest(sn, id, pod, msgID string, p *services.Protocol, status int,
	cid, evseId string, intellectType uint8, intellectId string, startTime string, intellectStatus uint8) *equipGetIntellectChargeRequest {
	req := &equipGetIntellectChargeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.GetIntellectCharge.FirstUpper(),
			AccessPod:   pod,
			MsgID:       msgID,
		},
		Callback: services.NewCB(status),
		Data: &equipGetIntellectChargeRequestDetail{
			EVSE: EVSE{
				Id:          evseId,
				ConnectorId: cid,
			},
			IntellectType: intellectType,
			IntellectId:   intellectId,
			StartTime:     startTime,
			Status:        intellectStatus,
		},
	}
	return req
}

func NewEquipGetIntellectChargeCallbackRequestError(sn, id, pod, msgId string, p *services.Protocol, err *apierrors.CallbackError) *equipGetIntellectChargeRequest {
	req := &equipGetIntellectChargeRequest{
		Base: services.Base{
			EquipmentSn: sn,
			EquipmentId: id,
			Protocol:    p,
			Category:    services.GetConfiguration.GetCallbackCategory(),
			AccessPod:   pod,
			MsgID:       msgId,
		},
		Callback: services.NewCBError(err),
		Data:     &equipGetIntellectChargeRequestDetail{},
	}
	return req
}

var _ services.Response = &equipGetIntellectChargeResponse{}

type equipGetIntellectChargeResponse struct {
	api.Response
	Data *equipGetIntellectChargeResponseDetail `json:"data"`
}

func (resp *equipGetIntellectChargeResponse) GetStatus() int {
	return resp.Status
}

func (resp *equipGetIntellectChargeResponse) GetMsg() string {
	return resp.Msg
}

type equipGetIntellectChargeResponseDetail struct {
}

func GetIntellectChargeCallbackRequest(ctx context.Context, req services.Request) error {
	header := services.GetCallbackHeaderValue(services.GetIntellectCharge)

	url := services.GetCallbackURL(req)

	return services.RequestWithoutResponse(ctx, req, url, header, &equipGetIntellectChargeResponse{})
}
